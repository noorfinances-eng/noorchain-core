# RUNBOOK — M17: WebSocket eth_subscribe("logs") (logrec, leader-only)

Status: **VALIDATED**  
Scope: **NOORCHAIN 2.1 (main)** — WS mainnet-like, milestone M17  
Semantics: `eth_subscribe("logs")` streams logs derived from `logrec/v1/*` via `eth_getLogs`, **leader-only** (until M18).  
Evidence anchor: tag **M17-WS-LOGS-STABLE**.

---

## 0. Objectives

This runbook validates that:

1) WS transport is available on the JSON-RPC endpoint.  
2) `eth_subscribe` supports `"logs"` with optional filter object.  
3) Logs are streamed as `eth_subscription` notifications, derived from the same semantics as `eth_getLogs` (logrec index).  
4) Follower endpoint (when `-follow-rpc` is set) returns `-32000 leader-only` for `eth_subscribe` / `eth_unsubscribe` (pre-M18).

---

## 1. Operational Discipline

**Terminal discipline (mandatory):**
- **T1 = nodes only** (start/stop leader/follower).
- **T2 = tooling only** (curl, wscat, scripts).
- One command per step, then validate the gate.

**Private key policy (mandatory):**
- Never commit or persist private keys.
- `NOOR_PRIVATE_KEY` must be supplied **at runtime only**.
- Format gate: `0x` + 64 hex chars (len = 66).

---

## 2. Preconditions

- Repository root: `/workspaces/noorchain-core`
- Branch: `main`
- Node binary: `./noorcore`
- Leader RPC: `http://127.0.0.1:8545` (WS: `ws://127.0.0.1:8545`)
- Follower RPC: `http://127.0.0.1:8546` (WS: `ws://127.0.0.1:8546`)
- Health (optional): leader `127.0.0.1:8081`, follower `127.0.0.1:8082`
- ChainId expected (local pack): `0x849` (2121)

---

## 3. Step-by-Step Validation

### Step 1 — Stop any running nodes (T1)

Command:
```bash
pkill -INT noorcore || true; sleep 0.4; pgrep -a noorcore || echo NO_PROCESSES
Gate:

Output contains NO_PROCESSES (or no processes listed).

Step 2 — Build the node binary (T2)
Command:

bash
Copier le code
cd /workspaces/noorchain-core && go build -o noorcore ./core && ./noorcore -h >/dev/null && echo BUILD_OK
Gate:

Output contains BUILD_OK.

Step 3 — Start leader (T1)
Command:

bash
Copier le code
cd /workspaces/noorchain-core && ./noorcore \
  -chain-id noorchain-2-1-local \
  -data-dir ./data/node1 \
  -role leader \
  -p2p-addr 127.0.0.1:30303 \
  -rpc-addr 127.0.0.1:8545 \
  -health-addr 127.0.0.1:8081 \
  > /tmp/noorcore_node1.log 2>&1 & disown; sleep 0.6; pgrep -a noorcore
Gate:

pgrep -a noorcore shows a leader process with -rpc-addr 127.0.0.1:8545.

Step 4 — Start follower (pre-M18 leader-only gate) (T1)
Command:

bash
Copier le code
cd /workspaces/noorchain-core && ./noorcore \
  -chain-id noorchain-2-1-local \
  -data-dir ./data/node2 \
  -role follower \
  -p2p-addr 127.0.0.1:30304 \
  -rpc-addr 127.0.0.1:8546 \
  -health-addr 127.0.0.1:8082 \
  -follow-rpc http://127.0.0.1:8545 \
  -boot-peers 127.0.0.1:30303 \
  > /tmp/noorcore_node2.log 2>&1 & disown; sleep 0.8; pgrep -a noorcore | grep 8546 || true
Gate:

Follower process is present with -rpc-addr 127.0.0.1:8546 and -follow-rpc http://127.0.0.1:8545.

Step 5 — RPC sanity gates (T2)
Command:

bash
Copier le code
curl -s http://127.0.0.1:8545 -H 'content-type: application/json' --data '{"jsonrpc":"2.0","id":1,"method":"eth_chainId","params":[]}' && echo && \
curl -s http://127.0.0.1:8545 -H 'content-type: application/json' --data '{"jsonrpc":"2.0","id":2,"method":"eth_blockNumber","params":[]}' && echo
Gate:

eth_chainId returns a 0x... value (expected 0x849).

eth_blockNumber returns a 0x... value.

Step 6 — WS client gate (T2)
This runbook uses wscat via npx (no global install required).

Command:

bash
Copier le code
npx -y wscat --version
Gate:

Output prints a version (e.g., 6.x.y).

Step 7 — WS subscribe (leader) (T2)
Command (interactive session):

bash
Copier le code
npx -y wscat -c ws://127.0.0.1:8545
Gate:

Output contains Connected.

Then send:

json
Copier le code
{"jsonrpc":"2.0","id":1,"method":"eth_subscribe","params":["logs",{}]}
Gate:

Response contains "result":"0x..." (subscription id).

Step 8 — Private Key Gate (T2)
Command:

bash
Copier le code
read -s -p "NOOR_PRIVATE_KEY (0x + 64 hex, len=66): " NOOR_PRIVATE_KEY; echo; export NOOR_PRIVATE_KEY; \
python3 - <<'PY'
import os,re,sys
k=os.environ.get("NOOR_PRIVATE_KEY","")
ok=bool(re.fullmatch(r"0x[0-9a-fA-F]{64}", k))
print("NOOR_PRIVATE_KEY_OK" if ok else f"NOOR_PRIVATE_KEY_BAD len={len(k)}")
sys.exit(0 if ok else 1)
PY
Gate:

Output is NOOR_PRIVATE_KEY_OK.

Step 9 — Trigger a log-producing transaction (T2)
In a separate T2 terminal (keep wscat open), run:

Command:

bash
Copier le code
cd /workspaces/noorchain-core && NOOR_RPC_URL=http://127.0.0.1:8545 node scripts/submit-snapshot.mjs
Gate:

Script prints receipt status: 1 and SUBMIT_OK.

Step 10 — Confirm WS log streaming (leader) (T2)
In the wscat window, confirm you receive at least one message with:

method: "eth_subscription"

params.subscription: "<same 0x.. id>"

params.result includes address, topics, blockNumber, transactionHash, logIndex, removed:false

Example shape (fields may vary):

json
Copier le code
{
  "jsonrpc":"2.0",
  "method":"eth_subscription",
  "params":{
    "subscription":"0x1",
    "result":{
      "address":"0x...",
      "topics":["0x..."],
      "blockNumber":"0x...",
      "transactionHash":"0x...",
      "logIndex":"0x0",
      "removed":false
    }
  }
}
Gate:

At least one valid eth_subscription notification is received after the tx is mined.

Step 11 — Confirm follower WS is leader-only (pre-M18) (T2)
Command:

bash
Copier le code
npx -y wscat -c ws://127.0.0.1:8546 -x '{"jsonrpc":"2.0","id":1,"method":"eth_subscribe","params":["logs",{}]}'
Gate:

Response contains:

error.code = -32000

error.message = "leader-only"

Step 12 — Unsubscribe (optional hygiene) (T2)
In the leader wscat session:

json
Copier le code
{"jsonrpc":"2.0","id":2,"method":"eth_unsubscribe","params":["0x1"]}
Gate:

Response is true.

4. Expected Results Summary
Pass criteria:

Leader WS subscribe returns a subscription id.

New logs are streamed via eth_subscription after a log-producing tx.

Follower WS subscribe returns -32000 leader-only when FollowRPC is set.

5. Troubleshooting
No WS messages received

Ensure the tx actually produced logs (receipt status must be 1, and the contract emits events).

Keep wscat session open while sending the transaction.

Verify leader eth_getLogs returns results for the expected block range.

Follower returns ECONNREFUSED

Follower is not running on 8546. Start it with -follow-rpc http://127.0.0.1:8545.

NOOR_PRIVATE_KEY errors

If signer list is empty or scripts fail with TypeError on address, the key is missing or malformed.

Enforce the Private Key Gate (len=66, 0x prefixed).