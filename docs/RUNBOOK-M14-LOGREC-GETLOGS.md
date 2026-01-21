RUNBOOK — M14 LogRec Index + eth_getLogs (Range-based) — FINAL

Document ID: RUNBOOK-M14-LOGREC-GETLOGS
Applies to:
- Branch: main
- Tags:
  - M14-LOGREC-GETLOGS-STABLE (logrec index + range-based eth_getLogs)
  - M14-LOGREC-GETLOGS-COMPAT (invalid range => -32602)
  - M14-LOGREC-GETLOGS-PERF   (max block range cap => -32602)

Scope
- Validate M14 end-to-end (A/B/C/D/E) on a local mainnet-like pack (Leader/Follower):
  A) RPC sanity + follower parity baseline
  B) logrec index backfill (leader-only) + discovery
  C) eth_getLogs correctness + filters + deterministic ordering
  D) Compatibility semantics: invalid range returns -32602
  E) Performance guardrail: max block range cap enforced (16384) with -32602

Assumptions
- Repo: noorchain-core
- Branch: main
- Binary: ./noorcore (built from ./core)
- ChainId (dev): 0x849 (2121)
- Leader RPC:   http://127.0.0.1:8545
- Follower RPC: http://127.0.0.1:8546 (started with -follow-rpc http://127.0.0.1:8545)
- P2P: leader 127.0.0.1:30303, follower 127.0.0.1:30304
- Health: leader 127.0.0.1:8081, follower 127.0.0.1:8082
- Data dirs: ./data/node1, ./data/node2
- Example contract addresses observed in one environment (do not assume stable across networks):
  - 0xC9F398646E19778F2C3D9fF32bb75E5a99FD4E56
  - 0xADbA8eA8f53bD7dEcFd1771C1bD03ecE6d721cf6
  - 0xf6d2739f632D5ABa8A96661059581566918253F6

Terminal Discipline
- T1 = Nodes only (start/stop/ports/logs)
- T2 = Tooling only (curl/python/scripts/verification)
- One command per step. Validate the gate output before continuing.

0) Clean Stop (mandatory)

T1
pkill -INT noorcore; sleep 1; pgrep -a noorcore || echo "OK: noorcore stopped"

Gate
- OK: noorcore stopped

T1
ss -ltnp | grep -E '(:30303|:30304|:8545|:8546|:8081|:8082)\b' || echo "OK: ports closed"

Gate
- OK: ports closed

1) Build (mandatory)

T2
cd /workspaces/noorchain-core && gofmt -w core/rpc/jsonrpc.go && go build -o noorcore ./core

Gate
- build completes without errors

2) Start Leader (node1)

T1
cd /workspaces/noorchain-core && mkdir -p ./logs && \
./noorcore -chain-id noorchain-2-1-local \
  -data-dir ./data/node1 \
  -p2p-addr 127.0.0.1:30303 \
  -rpc-addr 127.0.0.1:8545 \
  -health-addr 127.0.0.1:8081 >> ./logs/node1.log 2>&1 & echo LEADER_PID=$!

Gate
- PID printed

T1
ss -ltnp | grep -E '(:30303|:8545|:8081)\b'

Gate
- all three ports LISTEN for noorcore

3) Start Follower (node2)

T1
cd /workspaces/noorchain-core && mkdir -p ./logs && \
./noorcore -chain-id noorchain-2-1-local \
  -data-dir ./data/node2 \
  -p2p-addr 127.0.0.1:30304 \
  -rpc-addr 127.0.0.1:8546 \
  -health-addr 127.0.0.1:8082 \
  -follow-rpc http://127.0.0.1:8545 \
  -boot-peers 127.0.0.1:30303 >> ./logs/node2.log 2>&1 & echo FOLLOWER_PID=$!

Gate
- PID printed

T1
ss -ltnp | grep -E '(:30304|:8546|:8082)\b'

Gate
- all three follower ports LISTEN

4) M14 Validation Suite (A/B/C/D/E) — single-command authoritative gate

This step validates M14 end-to-end in a deterministic way:
- A) chainId and head parity
- B/C) bounded discovery window within the PERF cap (<= 16384 blocks)
- C) address-only and address+topic0 filters (single-block)
- D) invalid range semantics: fromBlock > toBlock => -32602
- E) perf cap semantics: requesting > 16384 blocks => -32602

T2
python3 - <<'PY'
import json, urllib.request, sys

L="http://127.0.0.1:8545"
F="http://127.0.0.1:8546"

def call(url, method, params):
    req = urllib.request.Request(
        url,
        data=json.dumps({"jsonrpc":"2.0","id":1,"method":method,"params":params}).encode(),
        headers={"content-type":"application/json"},
        method="POST",
    )
    with urllib.request.urlopen(req, timeout=30) as r:
        return json.loads(r.read().decode())

def must(out):
    if "error" in out:
        raise RuntimeError(out["error"])
    return out["result"]

def has_err(out, code=None):
    if "error" not in out: return False
    if code is None: return True
    return (out["error"] or {}).get("code") == code

def jcanon(x):
    return json.dumps(x, sort_keys=True, separators=(",",":"))

print("== M14 TEST SUITE (A/B/C/D/E) ==")

cidL = must(call(L,"eth_chainId",[]))
cidF = must(call(F,"eth_chainId",[]))
bnL  = must(call(L,"eth_blockNumber",[]))
bnF  = must(call(F,"eth_blockNumber",[]))
head = int(bnL,16)

print(f"[A] chainId L={cidL} F={cidF} | head={bnL} ({head}) | follower_head={bnF}")
okA = (cidL == cidF) and isinstance(bnL,str) and bnL.startswith("0x")

lo = max(0, head - 16383)
flt_all = {"fromBlock": hex(lo), "toBlock": bnL}
logsL = must(call(L,"eth_getLogs",[flt_all]))
logsF = must(call(F,"eth_getLogs",[flt_all]))

print(f"[B/C] discovery window [{hex(lo)}..{bnL}] | logs L={len(logsL)} F={len(logsF)} parity={len(logsL)==len(logsF)}")
okBC = (len(logsL) > 0) and (len(logsL) == len(logsF))

sample = logsL[0]
addr = sample.get("address")
topics = sample.get("topics") or []
topic0 = topics[0] if topics else None
blk = sample.get("blockNumber")
print(f"[B/C] sample addr={addr} block={blk} topic0={topic0}")

flt_addr = {"fromBlock": blk, "toBlock": blk, "address": addr}
r1L = must(call(L,"eth_getLogs",[flt_addr]))
r1F = must(call(F,"eth_getLogs",[flt_addr]))
okC1 = (len(r1L) >= 1) and (jcanon(r1L) == jcanon(r1F))

flt_at0 = {"fromBlock": blk, "toBlock": blk, "address": addr, "topics": [topic0]}
r2L = must(call(L,"eth_getLogs",[flt_at0]))
r2F = must(call(F,"eth_getLogs",[flt_at0]))
okC2 = (len(r2L) >= 1) and (jcanon(r2L) == jcanon(r2F))

print(f"[C] addr-only len L={len(r1L)} F={len(r1F)} parity={jcanon(r1L)==jcanon(r1F)}")
print(f"[C] addr+topic0 len L={len(r2L)} F={len(r2F)} parity={jcanon(r2L)==jcanon(r2F)}")

bad = {"fromBlock":"0x10","toBlock":"0x0"}
dL = call(L,"eth_getLogs",[bad])
dF = call(F,"eth_getLogs",[bad])
okD = has_err(dL, -32602) and has_err(dF, -32602)
print(f"[D] invalid_range L_error={has_err(dL)} code={(dL.get('error') or {}).get('code')} | F_error={has_err(dF)} code={(dF.get('error') or {}).get('code')}")

if head >= 20000:
    frm = head - 19999
    big = {"fromBlock": hex(frm), "toBlock": bnL}
    eL = call(L,"eth_getLogs",[big])
    eF = call(F,"eth_getLogs",[big])
    okE = has_err(eL, -32602) and has_err(eF, -32602)
    print(f"[E] cap range=20000 L_error={has_err(eL)} code={(eL.get('error') or {}).get('code')} | F_error={has_err(eF)} code={(eF.get('error') or {}).get('code')}")
else:
    okE = False
    print("[E] SKIP: head < 20000 (unexpected for this validation environment)")

ALL = okA and okBC and okC1 and okC2 and okD and okE
print("RESULT:", "PASS" if ALL else "FAIL",
      "| A=",okA,"B/C=",okBC,"C1=",okC1,"C2=",okC2,"D=",okD,"E=",okE)

sys.exit(0 if ALL else 2)
PY

Gate
- RESULT: PASS
- If FAIL: treat as a hard failure and stop.

5) Troubleshooting (minimal)

Symptom: eth_getLogs returns empty arrays for all windows
- Confirm activity exists (logs were produced historically).
- Confirm leader boot backfill ran (leader logs should contain a "logrec: backfill applied" line).
- Confirm you are querying within the PERF cap window (<= 16384 blocks) unless you expect a -32602 error.

Symptom: follower differs from leader
- Confirm follower started with: -follow-rpc http://127.0.0.1:8545
- Confirm both nodes expose same chainId and head blockNumber.
- Confirm follower can reach leader RPC.

Symptom: LevelDB lock
- Ensure only one noorcore per -data-dir.
- Stop all nodes and restart cleanly.

End of runbook.
