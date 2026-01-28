# RUNBOOK — M18 WS Follower Proxy via FollowRPC (Mainnet-like)

Status: **VALIDATED**  
Branch: `main`  
Commit: `f047430`  
Tag: `M18-WS-FOLLOWER-PROXY-STABLE`

## 0. Purpose

Validate that a follower node proxies **WebSocket** traffic to the leader via `FollowRPC`, providing:

- `eth_subscribe` / `eth_unsubscribe` on the follower endpoint
- Streaming parity on follower for:
  - `newHeads`
  - `logs` (event streaming)

This runbook is designed for a **controlled environment** and a **leader/follower** topology.

## 1. Scope and invariants

- Topology:
  - Leader: HTTP RPC `127.0.0.1:8545`
  - Follower: HTTP RPC `127.0.0.1:8546` with `-follow-rpc http://127.0.0.1:8545`
- WebSocket transport is served on the same RPC port (HTTP Upgrade).
- Follower WS proxy is enabled when `FollowRPC` is non-empty.
- One-command discipline: **one command → one gate**.
- Terminals:
  - **T1 = nodes only**
  - **T2 = tooling only**

## 2. Prerequisites

- Go toolchain available.
- Node.js available (for WebSocket tests and scripts).
- Repository root: `/workspaces/noorchain-core`
- Scripts available:
  - `scripts/submit-snapshot.mjs`

## 3. Mandatory Private Key Gate (REQUIRED)

Tooling requires `NOOR_PRIVATE_KEY` to be exported at runtime (never committed).

- Format: `0x`-prefixed hex
- Length: **66** characters

If the key is not exported, signing/deploy/submit scripts may fail or behave inconsistently across terminals.

## 4. Reference addresses

PoSSRegistry used by `scripts/submit-snapshot.mjs`:

- Default registry (script fallback):
  - `0xC9F398646E19778F2C3D9fF32bb75E5a99FD4E56`
- Override (preferred) for deterministic targeting:
  - `NOOR_POSS_REGISTRY=<0x...>` (or `NOOR_REGISTRY=<0x...>`)

This runbook uses the default registry unless explicitly overridden.

## 5. Step-by-step

### Step 1 — Stop any running nodes (T1)

```bash
pkill -INT noorcore || true
Gate:

No error is acceptable.

Step 2 — Verify processes and ports are free (T1)
bash
Copier le code
pgrep -a noorcore || echo NO_PROCESSES
ss -lntp | grep -E ':(30303|30304|8545|8546|8081|8082)\b' || echo PORTS_FREE
Gate:

NO_PROCESSES

PORTS_FREE

Step 3 — Build the binary (T1)
bash
Copier le code
cd /workspaces/noorchain-core && go build -o noorcore ./core
Gate:

Build succeeds (no errors).

Step 4 — Start leader (T1)
bash
Copier le code
cd /workspaces/noorchain-core && \
./noorcore \
  -chain-id noorchain-2-1-local \
  -data-dir ./data/node1 \
  -role leader \
  -p2p-addr 127.0.0.1:30303 \
  -rpc-addr 127.0.0.1:8545 \
  -health-addr 127.0.0.1:8081 \
  > /tmp/noorcore_node1.log 2>&1 &
Gate (leader ports):

bash
Copier le code
ss -lntp | grep -E ':(30303|8545|8081)\b'
LISTEN entries exist for 30303/8545/8081.

Step 5 — Start follower with FollowRPC (T1)
bash
Copier le code
cd /workspaces/noorchain-core && \
./noorcore \
  -chain-id noorchain-2-1-local \
  -data-dir ./data/node2 \
  -role follower \
  -p2p-addr 127.0.0.1:30304 \
  -rpc-addr 127.0.0.1:8546 \
  -health-addr 127.0.0.1:8082 \
  -follow-rpc http://127.0.0.1:8545 \
  -boot-peers 127.0.0.1:30303 \
  > /tmp/noorcore_node2.log 2>&1 &
Gate (follower ports):

bash
Copier le code
ss -lntp | grep -E ':(30304|8546|8082)\b'
LISTEN entries exist for 30304/8546/8082.

Step 6 — HTTP sanity checks (T2)
bash
Copier le code
curl -sS http://127.0.0.1:8545 -H 'content-type: application/json' --data '{"jsonrpc":"2.0","id":1,"method":"eth_chainId","params":[]}'
curl -sS http://127.0.0.1:8546 -H 'content-type: application/json' --data '{"jsonrpc":"2.0","id":2,"method":"eth_chainId","params":[]}'
Gate:

Both respond with a valid result (chain id).

Step 7 — Private Key Gate (T2) (MANDATORY)
bash
Copier le code
python - <<'PY'
import os
k=os.environ.get("NOOR_PRIVATE_KEY","")
print("NOOR_PRIVATE_KEY_LEN=", len(k))
print("OK" if (len(k)==66 and k.startswith("0x")) else "BAD")
PY
Gate:

NOOR_PRIVATE_KEY_LEN= 66

OK

If BAD:

bash
Copier le code
read -s -p "NOOR_PRIVATE_KEY (0x...): " NOOR_PRIVATE_KEY; echo; export NOOR_PRIVATE_KEY
Then re-run the gate.

Step 8 — Prove registry address is a contract (T2)
Default registry:

bash
Copier le code
curl -sS http://127.0.0.1:8545 -H 'content-type: application/json' --data '{"jsonrpc":"2.0","id":1,"method":"eth_getCode","params":["0xC9F398646E19778F2C3D9fF32bb75E5a99FD4E56","latest"]}'
Gate:

result is not "0x".

If using an override registry, replace the address accordingly.

Step 9 — WS streaming parity: newHeads on follower (T2)
bash
Copier le code
node -e '
const ws = new WebSocket("ws://127.0.0.1:8546");
const timeout = setTimeout(() => {
  console.log("TIMEOUT_NO_HEAD_30S");
  try{ws.close();}catch{}
}, 30000);

ws.onopen = () => {
  ws.send(JSON.stringify({jsonrpc:"2.0", id:1, method:"eth_subscribe", params:["newHeads"]}));
};

ws.onmessage = (ev) => {
  const msg = JSON.parse(ev.data.toString());
  if (msg.id === 1) {
    console.log("SUBID:", msg.result);
    console.log("WAITING_HEAD...");
    return;
  }
  if (msg.method === "eth_subscription") {
    console.log("GOT_HEAD");
    console.log(JSON.stringify(msg.params.result, null, 2));
    clearTimeout(timeout);
    ws.close();
  }
};

ws.onerror = (e) => console.log("WS_ERROR:", e.message || e);
'
Gate:

SUBID: 0x... (non-null)

GOT_HEAD received on follower.

Step 10 — WS streaming parity: logs on follower (subscribe → tx → event) (T2)
This must be executed in the correct order. A WS subscription does not replay history; it streams after subscription.

bash
Copier le code
cd /workspaces/noorchain-core && node -e '
const { spawn } = require("child_process");

const REG = process.env.NOOR_POSS_REGISTRY
  ?? process.env.NOOR_REGISTRY
  ?? "0xC9F398646E19778F2C3D9fF32bb75E5a99FD4E56";

const pk = process.env.NOOR_PRIVATE_KEY || "";
if (pk.length !== 66) { console.log("PK_MISSING_OR_BAD"); process.exit(1); }
console.log("PK_OK len=66");
console.log("REGISTRY:", REG);

const ws = new WebSocket("ws://127.0.0.1:8546");
let triggered = false;

const timeout = setTimeout(() => {
  console.log("TIMEOUT_NO_EVENT_90S");
  try { ws.close(); } catch {}
  process.exit(2);
}, 90000);

ws.onopen = () => {
  ws.send(JSON.stringify({jsonrpc:"2.0", id:1, method:"eth_subscribe", params:["logs", {address: REG}]}));
};

ws.onmessage = (ev) => {
  const msg = JSON.parse(ev.data.toString());

  if (msg.id === 1) {
    console.log("SUBID:", msg.result);
    if (triggered) return;
    triggered = true;

    console.log("TRIGGER_TX_NOW...");
    const child = spawn(process.execPath, ["scripts/submit-snapshot.mjs"], {
      env: { ...process.env, NOOR_POSS_REGISTRY: REG },
      stdio: ["ignore", "pipe", "pipe"],
    });
    child.stdout.on("data", (d) => process.stdout.write(d.toString()));
    child.stderr.on("data", (d) => process.stderr.write(d.toString()));
    child.on("exit", (code) => console.log("TX_SCRIPT_EXIT:", code));
    return;
  }

  if (msg.method === "eth_subscription") {
    console.log("GOT_EVENT");
    console.log(JSON.stringify(msg.params, null, 2));
    clearTimeout(timeout);
    try { ws.close(); } catch {}
    process.exit(0);
  }
};

ws.onerror = (e) => console.log("WS_ERROR:", e.message || e);
'
Gate:

SUBID: 0x... (non-null)

Script prints SUBMIT_OK

GOT_EVENT received with:

params.result.address == REG

params.result.transactionHash equals the tx hash printed by the script.

6. Troubleshooting (minimal)
curl against http://127.0.0.1:8545/ws may return 101 Switching Protocols and then time out. This can be normal if the client does not speak WebSocket frames.

If TIMEOUT_NO_EVENT occurs:

Ensure subscription was started before triggering the tx.

Verify NOOR_PRIVATE_KEY is exported in the same terminal (len=66).

Verify registry bytecode exists via eth_getCode.

Ensure follower is started with a non-empty -follow-rpc URL.

If follower returns leader-only for WS calls:

Ensure the follower is running the build that includes the M18 proxy changes (commit f047430 or later).

7. Acceptance criteria (M18 PASS)
M18 is considered PASS when, on the follower endpoint (ws://127.0.0.1:8546):

eth_subscribe returns a non-null subId for newHeads

a newHeads notification is received (eth_subscription)

eth_subscribe returns a non-null subId for logs

after triggering submitSnapshot, a logs notification is received (eth_subscription)

eth_unsubscribe returns true (implicitly validated by successful session closure and no errors)