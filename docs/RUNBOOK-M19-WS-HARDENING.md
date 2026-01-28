# RUNBOOK — M19 WS HARDENING (Limits / Timeouts / Keepalive / Quotas)

Status: VALIDATION RUNBOOK (ops-first)  
Scope: WebSocket JSON-RPC hardening on NOORCHAIN 2.1 (leader + follower with FollowRPC proxy)  
Audience: operators / core devs  
Environment: local mainnet-like pack (2 nodes), controlled network (NOT public internet)

---

## 0) Objectives (M19)

This runbook validates the following hardening controls:

1) **Message limits**
- WS read message size is capped (`SetReadLimit`) to prevent memory abuse.

2) **Timeouts / keepalive**
- Read deadline + pong handler (`SetReadDeadline` + `SetPongHandler`)
- Periodic ping loop (`PingMessage`) to keep connections healthy

3) **Concurrency-safe writes**
- WS writes are serialized via a mutex (avoids concurrent write races when pings + app writes happen)

4) **Subscription quota**
- Hard cap on number of subscriptions per WS connection (`wsMaxSubsPerConn`)
- When exceeded, server returns `-32000 "too many subscriptions"`

This runbook also validates **follower WS proxy** behavior (M18): follower streams match leader streams via `FollowRPC`.

---

## 1) Terminals & Discipline

- **T1 = nodes only** (leader/follower)
- **T2 = tooling only** (git/build/curl/wscat)
- One command per step. Do not mix terminals.

---

## 2) Preconditions

- Repo: `/workspaces/noorchain-core`
- Branch: `main`
- Nodes data dirs:
  - Leader: `./data/node1`
  - Follower: `./data/node2`
- Ports:
  - Leader: P2P `30303`, RPC/WS `8545`, health `8081`
  - Follower: P2P `30304`, RPC/WS `8546`, health `8082`
- Follower follow-rpc:
  - `-follow-rpc http://127.0.0.1:8545`

Tools:
- `curl`
- `npx -y wscat`

---

## 3) Stop / Build / Start (Clean)

### T1 — Stop nodes
```bash
pkill -INT noorcore || true
T1 — Gate: ports free
bash
Copier le code
ss -lntp | grep -E ':(30303|30304|8545|8546|8081|8082)\b' || echo PORTS_FREE
Gate: PORTS_FREE

T2 — Build (always rebuild ./noorcore)
bash
Copier le code
cd /workspaces/noorchain-core && go build -o noorcore ./core
Gate: no output

T1 — Start leader (node1)
bash
Copier le code
cd /workspaces/noorchain-core && ./noorcore -chain-id noorchain-2-1-local -data-dir ./data/node1 -role leader -p2p-addr 127.0.0.1:30303 -rpc-addr 127.0.0.1:8545 -health-addr 127.0.0.1:8081 > /tmp/noorcore_node1.log 2>&1 &
T1 — Gate: leader ports + chainId
bash
Copier le code
ss -lntp | grep -E ':(30303|8545|8081)\b' && curl -s http://127.0.0.1:8545 -H 'content-type: application/json' --data '{"jsonrpc":"2.0","id":1,"method":"eth_chainId","params":[]}'
Gate: listening on 30303/8545/8081 and chainId 0x849

T1 — Start follower (node2)
bash
Copier le code
cd /workspaces/noorchain-core && ./noorcore -chain-id noorchain-2-1-local -data-dir ./data/node2 -role follower -p2p-addr 127.0.0.1:30304 -rpc-addr 127.0.0.1:8546 -health-addr 127.0.0.1:8082 -follow-rpc http://127.0.0.1:8545 -boot-peers 127.0.0.1:30303 > /tmp/noorcore_node2.log 2>&1 &
T1 — Gate: follower ports + chainId + P2P ESTAB
bash
Copier le code
ss -lntp | grep -E ':(30304|8546|8082)\b' && curl -s http://127.0.0.1:8546 -H 'content-type: application/json' --data '{"jsonrpc":"2.0","id":2,"method":"eth_chainId","params":[]}' && ss -ntp | grep -E ':(30303|30304)\b' | grep ESTAB || echo NO_ESTAB_YET
Gate: chainId 0x849 and at least one ESTAB between 30303/30304

4) WS Gates (Leader)
T2 — Gate A: subscribe newHeads (leader)
bash
Copier le code
npx -y wscat -c ws://127.0.0.1:8545 -x '{"jsonrpc":"2.0","id":1,"method":"eth_subscribe","params":["newHeads"]}'
Gate:

response: {"id":1,"result":"0x..."}

then periodic eth_subscription notifications (new blocks)

T2 — Gate B: subscribe logs (leader)
bash
Copier le code
npx -y wscat -c ws://127.0.0.1:8545 -x '{"jsonrpc":"2.0","id":2,"method":"eth_subscribe","params":["logs",{}]}'
Gate: response result:"0x..."

(Logs notifications require an on-chain event; validate separately if needed.)

5) WS Gates (Follower Proxy)
T2 — Gate C: subscribe newHeads (follower)
bash
Copier le code
npx -y wscat -c ws://127.0.0.1:8546 -x '{"jsonrpc":"2.0","id":1,"method":"eth_subscribe","params":["newHeads"]}'
Gate:

response result:"0x..."

notifications arrive continuously (proxy streaming parity)

T2 — Gate D: subscribe logs (follower)
bash
Copier le code
npx -y wscat -c ws://127.0.0.1:8546 -x '{"jsonrpc":"2.0","id":2,"method":"eth_subscribe","params":["logs",{}]}'
Gate: response result:"0x..."

6) Hardening Gates (M19)
6.1) Batch correctness (no blocking)
Use a short timeout to ensure the client exits deterministically.

T2 — Gate E: batch 5 newHeads subs (leader)
bash
Copier le code
timeout 3s npx -y wscat -c ws://127.0.0.1:8545 -x '[{"jsonrpc":"2.0","id":1,"method":"eth_subscribe","params":["newHeads"]},{"jsonrpc":"2.0","id":2,"method":"eth_subscribe","params":["newHeads"]},{"jsonrpc":"2.0","id":3,"method":"eth_subscribe","params":["newHeads"]},{"jsonrpc":"2.0","id":4,"method":"eth_subscribe","params":["newHeads"]},{"jsonrpc":"2.0","id":5,"method":"eth_subscribe","params":["newHeads"]}]'
Gate: results for ids 1..5 then timeout exits

6.2) Subscription quota
Validate that the server enforces wsMaxSubsPerConn and returns a controlled error.

T2 — Gate F: 129 newHeads subs (leader) must trigger quota error
bash
Copier le code
bash -lc 'python - <<'"'"'PY'"'"' > /tmp/ws129.json
import json
reqs=[{"jsonrpc":"2.0","id":i,"method":"eth_subscribe","params":["newHeads"]} for i in range(1,130)]
print(json.dumps(reqs))
PY
timeout 4s npx -y wscat -c ws://127.0.0.1:8545 -x "$(cat /tmp/ws129.json)" 2>/dev/null | grep -m1 "too many subscriptions" || echo QUOTA_NOT_TRIGGERED
rm -f /tmp/ws129.json'
Gate: output contains too many subscriptions (error code -32000)

7) Evidence Pack (Minimum)
Record:

git rev-parse --short HEAD

git tag --points-at HEAD

Snippets (copy/paste):

Gate A output (subscribe id + at least 1 notification)

Gate C output (subscribe id + at least 1 notification)

Gate F output showing too many subscriptions

Optional:

/tmp/noorcore_node1.log and /tmp/noorcore_node2.log tail (last 30 lines)

8) Operational Notes
WS is for a controlled environment; CheckOrigin currently allows all origins.

Quota is per WS connection; closing the connection releases subscriptions (cleanup on disconnect).

If follower WS is enabled, FollowRPC is authoritative for proxy behavior.

EOF