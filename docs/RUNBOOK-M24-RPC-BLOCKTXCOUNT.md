# RUNBOOK — M24 RPC: Block Transaction Count (byNumber / byHash)

## Scope

This runbook validates JSON-RPC compatibility for:

- `eth_getBlockTransactionCountByNumber`
- `eth_getBlockTransactionCountByHash`

It also validates **pending semantics** (current design): `pending` is treated as `latest` (no “future block”), and verifies **leader/follower parity** via `-follow-rpc`.

This runbook is written for a controlled “mainnet-like” environment.

## Assumptions (current behavior)

- Blocks may be empty; `transactions: []` is expected in `eth_getBlockByNumber`.
- `pending` block tag returns the same block as `latest` (same `number` and `hash`).
- Block tx count is inferred from persisted block metadata (`blkmeta/v1/*`), consistent with current mining path.

## Terminals and discipline

- **T1** = nodes only (no tooling commands).
- **T2** = tooling only (curl, git, hardhat/viem, etc.).
- One command per step. After each step, run the gate checks before continuing.

## Variables (defaults)

- Chain ID: `noorchain-2-1-local` (EVM chainId = `2121` / `0x849`)
- Leader:
  - P2P `127.0.0.1:30303`
  - RPC `127.0.0.1:8545`
  - Health `127.0.0.1:8081`
  - Data dir `./data/node1`
- Follower:
  - P2P `127.0.0.1:30304`
  - RPC `127.0.0.1:8546`
  - Health `127.0.0.1:8082`
  - Data dir `./data/node2`
  - Follow RPC `http://127.0.0.1:8545`
- Logs:
  - `/tmp/noorcore_node1.log`
  - `/tmp/noorcore_node2.log`

## Required environment gates (MANDATORY)

- `NOOR_PRIVATE_KEY` must be exported at runtime:
  - `0x` prefixed
  - length = `66`
  - never committed
- `NOOR_POSS_REGISTRY` must be exported at runtime:
  - `0x` prefixed
  - length = `42`

---

## Step 0 — Stop nodes and free ports (T2)

### Command
```bash
pkill -INT noorcore || true
pgrep -a noorcore || echo NO_PROCESSES
ss -lntp | grep -E ':(30303|30304|8545|8546|8081|8082)\b' || echo PORTS_FREE
Gate
NO_PROCESSES

PORTS_FREE

Step 1 — Build the binary (T2)
Command
bash
Copier le code
cd /workspaces/noorchain-core && go build -o noorcore ./core
Gate
Command exits 0

./noorcore -h prints usage

Step 2 — Start leader + follower (T1)
Command
bash
Copier le code
cd /workspaces/noorchain-core && \
./noorcore -chain-id noorchain-2-1-local -data-dir ./data/node1 -role leader \
  -p2p-addr 127.0.0.1:30303 -rpc-addr 127.0.0.1:8545 -health-addr 127.0.0.1:8081 \
  -boot-peers 127.0.0.1:30304 > /tmp/noorcore_node1.log 2>&1 & echo "LEADER_PID=$!" && \
./noorcore -chain-id noorchain-2-1-local -data-dir ./data/node2 -role follower \
  -p2p-addr 127.0.0.1:30304 -rpc-addr 127.0.0.1:8546 -health-addr 127.0.0.1:8082 \
  -follow-rpc http://127.0.0.1:8545 -boot-peers 127.0.0.1:30303 > /tmp/noorcore_node2.log 2>&1 & echo "FOLLOWER_PID=$!"
Gate (T2)
bash
Copier le code
ss -lntp | grep -E ':(30303|30304|8545|8546|8081|8082)\b'
curl -s http://127.0.0.1:8545 -H 'content-type: application/json' --data '{"jsonrpc":"2.0","id":1,"method":"eth_chainId","params":[]}' && echo
curl -s http://127.0.0.1:8546 -H 'content-type: application/json' --data '{"jsonrpc":"2.0","id":2,"method":"eth_chainId","params":[]}' && echo
Expected:

Ports listening on both nodes

Both eth_chainId return 0x849

Step 3 — Private Key Gate + PoSS Registry Gate (T2) (MANDATORY)
Command
bash
Copier le code
python3 - <<'PY'
import os
k = os.getenv("NOOR_PRIVATE_KEY","")
r = os.getenv("NOOR_POSS_REGISTRY","")
print("NOOR_PRIVATE_KEY.len =", len(k))
print("NOOR_PRIVATE_KEY.ok  =", k.startswith("0x") and len(k)==66)
print("NOOR_POSS_REGISTRY.len =", len(r))
print("NOOR_POSS_REGISTRY.ok  =", r.startswith("0x") and len(r)==42)
PY
Gate
NOOR_PRIVATE_KEY.ok = True

NOOR_POSS_REGISTRY.ok = True

If this gate fails, stop here. Tooling will not have a signer and PoSS scripts may target the wrong contract.

Step 4 — Produce a tx-bearing block (submitSnapshot) (T2)
This step ensures we have at least one known block that contains a transaction linked to PoSS logs.

Command
bash
Copier le code
cd /workspaces/noorchain-core && npx hardhat run scripts/submit-snapshot.mjs --network noorcore --config hardhat.config.mjs
Gate
Script prints SUBMIT_OK

It prints a tx hash 0x... (record it as TX_HASH)

Step 5 — Receipt gate (leader) (T2)
Command
Replace TX_HASH with the hash from Step 4.

bash
Copier le code
curl -s http://127.0.0.1:8545 -H 'content-type: application/json' --data "{\"jsonrpc\":\"2.0\",\"id\":80,\"method\":\"eth_getTransactionReceipt\",\"params\":[\"TX_HASH\"]}" && echo
Gate
status is 0x1

blockNumber is present (record it as TX_BLOCK_NUMBER)

blockHash is present (record it as TX_BLOCK_HASH)

logs contains at least one entry with address == NOOR_POSS_REGISTRY

Step 6 — Validate block tx count (leader) (T2)
Command
Replace TX_BLOCK_NUMBER and TX_BLOCK_HASH from Step 5.

bash
Copier le code
echo "== :8545 countByNumber latest ==" &&
curl -s http://127.0.0.1:8545 -H 'content-type: application/json' --data '{"jsonrpc":"2.0","id":1,"method":"eth_getBlockTransactionCountByNumber","params":["latest"]}' && echo && \
echo "== :8545 countByNumber pending ==" &&
curl -s http://127.0.0.1:8545 -H 'content-type: application/json' --data '{"jsonrpc":"2.0","id":2,"method":"eth_getBlockTransactionCountByNumber","params":["pending"]}' && echo && \
echo "== :8545 countByNumber TX_BLOCK_NUMBER ==" &&
curl -s http://127.0.0.1:8545 -H 'content-type: application/json' --data "{\"jsonrpc\":\"2.0\",\"id\":3,\"method\":\"eth_getBlockTransactionCountByNumber\",\"params\":[\"TX_BLOCK_NUMBER\"]}" && echo && \
echo "== :8545 countByHash TX_BLOCK_HASH ==" &&
curl -s http://127.0.0.1:8545 -H 'content-type: application/json' --data "{\"jsonrpc\":\"2.0\",\"id\":4,\"method\":\"eth_getBlockTransactionCountByHash\",\"params\":[\"TX_BLOCK_HASH\"]}" && echo
Gate
latest returns 0x0 (typical if head is empty)

pending returns the same as latest (current design)

TX_BLOCK_NUMBER returns 0x1

TX_BLOCK_HASH returns 0x1

(If your head currently contains a tx, latest may be 0x1. That is acceptable as long as results are consistent with the block’s tx-bearing status.)

Step 7 — Pending == Latest block identity (leader) (T2)
Command
bash
Copier le code
echo "== latest ==" &&
curl -s http://127.0.0.1:8545 -H 'content-type: application/json' --data '{"jsonrpc":"2.0","id":20,"method":"eth_getBlockByNumber","params":["latest",false]}' && echo && \
echo "== pending ==" &&
curl -s http://127.0.0.1:8545 -H 'content-type: application/json' --data '{"jsonrpc":"2.0","id":21,"method":"eth_getBlockByNumber","params":["pending",false]}' && echo
Gate
pending.hash == latest.hash

pending.number == latest.number

Step 8 — Follower parity (T2)
Repeat Steps 5–7 using RPC :8546.

Command (tx count parity)
bash
Copier le code
echo "== :8546 countByNumber TX_BLOCK_NUMBER ==" &&
curl -s http://127.0.0.1:8546 -H 'content-type: application/json' --data "{\"jsonrpc\":\"2.0\",\"id\":30,\"method\":\"eth_getBlockTransactionCountByNumber\",\"params\":[\"TX_BLOCK_NUMBER\"]}" && echo && \
echo "== :8546 countByHash TX_BLOCK_HASH ==" &&
curl -s http://127.0.0.1:8546 -H 'content-type: application/json' --data "{\"jsonrpc\":\"2.0\",\"id\":31,\"method\":\"eth_getBlockTransactionCountByHash\",\"params\":[\"TX_BLOCK_HASH\"]}" && echo
Gate
Results match :8545 exactly

Evidence pack (recommended) (T2)
Collect the most important outputs into a single file.

Command
bash
Copier le code
E=/tmp/M24_EVIDENCE_RPC_BLOCKTXCOUNT.txt
: > "$E"
echo "[A] chainId leader/follower" >> "$E"
curl -s http://127.0.0.1:8545 -H 'content-type: application/json' --data '{"jsonrpc":"2.0","id":1,"method":"eth_chainId","params":[]}' >> "$E"; echo >> "$E"
curl -s http://127.0.0.1:8546 -H 'content-type: application/json' --data '{"jsonrpc":"2.0","id":2,"method":"eth_chainId","params":[]}' >> "$E"; echo >> "$E"
echo "[B] receipt TX_HASH (leader)" >> "$E"
curl -s http://127.0.0.1:8545 -H 'content-type: application/json' --data "{\"jsonrpc\":\"2.0\",\"id\":80,\"method\":\"eth_getTransactionReceipt\",\"params\":[\"TX_HASH\"]}" >> "$E"; echo >> "$E"
echo "[C] blockTxCount byNumber/byHash (leader)" >> "$E"
curl -s http://127.0.0.1:8545 -H 'content-type: application/json' --data '{"jsonrpc":"2.0","id":1,"method":"eth_getBlockTransactionCountByNumber","params":["latest"]}' >> "$E"; echo >> "$E"
curl -s http://127.0.0.1:8545 -H 'content-type: application/json' --data '{"jsonrpc":"2.0","id":2,"method":"eth_getBlockTransactionCountByNumber","params":["pending"]}' >> "$E"; echo >> "$E"
curl -s http://127.0.0.1:8545 -H 'content-type: application/json' --data "{\"jsonrpc\":\"2.0\",\"id\":3,\"method\":\"eth_getBlockTransactionCountByNumber\",\"params\":[\"TX_BLOCK_NUMBER\"]}" >> "$E"; echo >> "$E"
curl -s http://127.0.0.1:8545 -H 'content-type: application/json' --data "{\"jsonrpc\":\"2.0\",\"id\":4,\"method\":\"eth_getBlockTransactionCountByHash\",\"params\":[\"TX_BLOCK_HASH\"]}" >> "$E"; echo >> "$E"
echo "[D] pending == latest block identity (leader)" >> "$E"
curl -s http://127.0.0.1:8545 -H 'content-type: application/json' --data '{"jsonrpc":"2.0","id":20,"method":"eth_getBlockByNumber","params":["latest",false]}' >> "$E"; echo >> "$E"
curl -s http://127.0.0.1:8545 -H 'content-type: application/json' --data '{"jsonrpc":"2.0","id":21,"method":"eth_getBlockByNumber","params":["pending",false]}' >> "$E"; echo >> "$E"
echo "EVIDENCE_FILE=$E"
Failure handling
If any gate fails:

Stop nodes and free ports (Step 0).

Inspect logs:

/tmp/noorcore_node1.log

/tmp/noorcore_node2.log

Re-run from Step 1.

Completion criteria
M24 is considered validated when all gates pass:

HTTP JSON-RPC parity leader/follower for:

eth_getBlockTransactionCountByNumber

eth_getBlockTransactionCountByHash

pending returns the same block identity as latest

A known tx-bearing block returns txCount = 0x1 via both byNumber and byHash

Evidence pack captured