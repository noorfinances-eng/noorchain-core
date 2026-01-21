RUNBOOK — M14 LogRec Index + eth_getLogs (Range-based)

Document ID: RUNBOOK-M14-LOGREC-GETLOGS
Scope: Local mainnet-like pack (Leader/Follower) validation for logrec index write-path, leader-only boot backfill, and eth_getLogs correctness/parity.
Status: Applies to tag M14-LOGREC-GETLOGS-STABLE.

Assumptions

- Repo: noorchain-core, branch main
- Binary: ./noorcore (built from ./core)
- ChainId (dev): 0x849 (2121)
- Leader RPC: 127.0.0.1:8545
- Follower RPC: 127.0.0.1:8546 (with -follow-rpc http://127.0.0.1:8545)
- P2P: leader 127.0.0.1:30303, follower 127.0.0.1:30304
- Health: leader 127.0.0.1:8081, follower 127.0.0.1:8082
- Data dirs: ./data/node1, ./data/node2
- Example contract addresses observed in this environment (do not assume these are stable across networks):
  - 0xC9F398646E19778F2C3D9fF32bb75E5a99FD4E56
  - 0xADbA8eA8f53bD7dEcFd1771C1bD03ecE6d721cf6
  - 0xf6d2739f632D5ABa8A96661059581566918253F6

Terminal Discipline

- T1 = Nodes only (start/stop/ports/logs)
- T2 = Tooling only (curl, scripts, verification)
- One command per step. Validate gate output before continuing.

0) Clean Stop (mandatory)

T1
pkill -INT noorcore; sleep 1; pgrep -a noorcore || echo "OK: noorcore stopped"

Gate: OK: noorcore stopped

T1
ss -ltnp | grep -E '(:30303|:30304|:8545|:8546|:8081|:8082)\b' || echo "OK: ports closed"

Gate: OK: ports closed

1) Build (mandatory)

T2
cd /workspaces/noorchain-core && gofmt -w core/rpc/jsonrpc.go && go build -o noorcore ./core

Gate: build completes without errors.

2) Start Leader (node1)

T1
cd /workspaces/noorchain-core && mkdir -p ./logs && \
./noorcore -chain-id noorchain-2-1-local \
  -data-dir ./data/node1 \
  -p2p-addr 127.0.0.1:30303 \
  -rpc-addr 127.0.0.1:8545 \
  -health-addr 127.0.0.1:8081 >> ./logs/node1.log 2>&1 & echo PID=$!

Gate: PID printed.

T1
ss -ltnp | grep -E '(:30303|:8545|:8081)\b'

Gate: all three ports are LISTENing for noorcore.

T2
curl -s -X POST http://127.0.0.1:8545 -H 'content-type: application/json' \
  --data '{"jsonrpc":"2.0","id":1,"method":"eth_chainId","params":[]}' ; echo
curl -s -X POST http://127.0.0.1:8545 -H 'content-type: application/json' \
  --data '{"jsonrpc":"2.0","id":2,"method":"eth_blockNumber","params":[]}' ; echo

Gate: eth_chainId returns 0x849 and eth_blockNumber returns a non-empty hex quantity.

3) Start Follower (node2)

T1
cd /workspaces/noorchain-core && mkdir -p ./logs && \
./noorcore -chain-id noorchain-2-1-local \
  -data-dir ./data/node2 \
  -p2p-addr 127.0.0.1:30304 \
  -rpc-addr 127.0.0.1:8546 \
  -health-addr 127.0.0.1:8082 \
  -follow-rpc http://127.0.0.1:8545 \
  -boot-peers 127.0.0.1:30303 >> ./logs/node2.log 2>&1 & echo PID=$!

Gate: PID printed.

T1
ss -ltnp | grep -E '(:30304|:8546|:8082)\b'

Gate: all three follower ports are LISTENing.

T2
curl -s -X POST http://127.0.0.1:8546 -H 'content-type: application/json' \
  --data '{"jsonrpc":"2.0","id":1,"method":"eth_chainId","params":[]}' ; echo
curl -s -X POST http://127.0.0.1:8546 -H 'content-type: application/json' \
  --data '{"jsonrpc":"2.0","id":2,"method":"eth_blockNumber","params":[]}' ; echo

Gate: chainId 0x849 and blockNumber is a non-empty hex quantity.

4) logrec + eth_getLogs Functional Gates (Leader)

Notes (implementation-level intent)

- logrec is written on the mining path (per block) and used as the primary scan surface for eth_getLogs.
- On boot, leader performs an idempotent backfill from persisted receipts into logrec.
- Follower is expected to return identical results to leader when -follow-rpc is configured and routing is correct.

4.0 Auto-discovery (recommended)

T2
python3 - <<'PY2'
import json, urllib.request

RPC="http://127.0.0.1:8545"

def rpc(method, params):
    req = urllib.request.Request(
        RPC,
        data=json.dumps({"jsonrpc":"2.0","id":1,"method":method,"params":params}).encode(),
        headers={"content-type":"application/json"},
        method="POST",
    )
    with urllib.request.urlopen(req, timeout=30) as r:
        out = json.loads(r.read().decode())
    if "error" in out:
        raise RuntimeError(out["error"])
    return out["result"]

head_hex = rpc("eth_blockNumber", [])
head = int(head_hex, 16)

# Progressive search windows to avoid unbounded scans by default.
windows = [256, 1024, 4096, 16384, 65536]
found = []
used_from = 0

for w in windows:
    frm = max(0, head - w + 1)
    used_from = frm
    logs = rpc("eth_getLogs", [{"fromBlock": hex(frm), "toBlock": head_hex}])
    if logs:
        found = logs
        break

# Last resort (only if still empty): scan from genesis.
if not found:
    logs = rpc("eth_getLogs", [{"fromBlock": "0x0", "toBlock": head_hex}])
    found = logs
    used_from = 0

print(f"head={head_hex} ({head}) logs_total={len(found)} searched_from={hex(used_from)}")
if found:
    blocks = [int(l["blockNumber"],16) for l in found if "blockNumber" in l]
    addrs = sorted(set(l.get("address") for l in found if l.get("address")))
    print(f"blocks.min={hex(min(blocks))} blocks.max={hex(max(blocks))}")
    print("addresses=" + ",".join(addrs))
    top = max(blocks)
    print(f"suggested_window: fromBlock={hex(max(0, top-256))} toBlock={hex(top)}")
PY2

Gate:
- logs_total > 0
- Use suggested_window for bounded perf checks (4.1/4.2/4.3) to avoid assuming a specific historical block.

4.1 Sanity: query a narrow window near head

T2
# Replace FROM/TO once you know a block that contains logs (see 4.2).
# Example: fromBlock == toBlock for a single block query.
curl -s -X POST http://127.0.0.1:8545 -H 'content-type: application/json' \
  --data '{"jsonrpc":"2.0","id":41,"method":"eth_getLogs","params":[{"fromBlock":"latest","toBlock":"latest","address":"0xADbA8eA8f53bD7dEcFd1771C1bD03ecE6d721cf6"}]}' ; echo

Gate: returns a JSON array (possibly empty), without JSON-RPC error.

4.2 Find a block range that contains PoSS logs

T2
# Use a recent bounded range. Adjust the numbers if needed.
# Example: scan last 256 blocks (head-255 .. head).
# If your node is young, reduce the range.
curl -s -X POST http://127.0.0.1:8545 -H 'content-type: application/json' \
  --data '{"jsonrpc":"2.0","id":42,"method":"eth_getLogs","params":[{"fromBlock":"0x0","toBlock":"latest","address":"0xADbA8eA8f53bD7dEcFd1771C1bD03ecE6d721cf6"}]}' ; echo

Gate: returns a JSON array. If empty, ensure the address is correct and that events were produced.

Operational recommendation:
- Do not rely on unbounded scans for routine checks. Prefer bounded slices (e.g., 128–2048 blocks).

4.3 Ordering + logIndex stability (single block)

T2
# Once you have a known block height H containing multiple logs, pin it:
# Replace 0xHHH with that block number.
curl -s -X POST http://127.0.0.1:8545 -H 'content-type: application/json' \
  --data '{"jsonrpc":"2.0","id":43,"method":"eth_getLogs","params":[{"fromBlock":"0x18c","toBlock":"0x18c","address":"0xADbA8eA8f53bD7dEcFd1771C1bD03ecE6d721cf6"}]}' ; echo

Gate:
- Returned logs are in canonical order.
- For the same query repeated twice, results are byte-for-byte identical.
- If multiple logs are present in the same block, log ordering is stable.

5) Follower Parity Gates (eth_getLogs)

Run the same queries against follower RPC (8546) and compare results.

T2
curl -s -X POST http://127.0.0.1:8546 -H 'content-type: application/json' \
  --data '{"jsonrpc":"2.0","id":51,"method":"eth_getLogs","params":[{"fromBlock":"0x18c","toBlock":"0x18c","address":"0xADbA8eA8f53bD7dEcFd1771C1bD03ecE6d721cf6"}]}' ; echo

Gate: follower response matches leader response byte-for-byte for the same filter.

6) Edge Cases (behavioral checks)

6.1 Empty results are not errors

- Query a range with no logs (e.g., a fresh block with no contract activity).
- Expect: empty array [] (not an error).

6.2 Invalid range / malformed params

- If the node returns a JSON-RPC error, treat it as a hard failure.
- If the node returns [], validate the params and retry with a known-good filter.

6.3 Repeated calls must be stable

- Repeat the same eth_getLogs query twice.
- Expect: identical output ordering and content.

7) Performance Notes (practical constraints)

- Prefer bounded ranges.
- For large historical coverage, chunk by fixed block windows and merge client-side.
- Keep follower traffic “read-only parity”; run heavy scans on leader if resource-constrained.

8) Troubleshooting

Symptom: eth_getLogs always returns []
- Confirm the target contract address (PoSSRegistry) is correct for this environment.
- Confirm blocks contain events (submitSnapshot or other event-producing transactions mined).
- Confirm leader is running the tag build that includes logrec index + eth_getLogs routing.
- Restart leader once to force boot backfill to run (idempotent) and retry.

Symptom: follower differs from leader
- Confirm follower is started with -follow-rpc http://127.0.0.1:8545
- Confirm leader is reachable from follower.
- Confirm both nodes are on the same chain-id and see the same head height.

Symptom: node will not start (LevelDB lock)
- Ensure a single noorcore process per -data-dir.
- Stop all nodes, verify ports closed, then restart.

