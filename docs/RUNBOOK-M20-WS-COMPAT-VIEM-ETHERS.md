# RUNBOOK — M20 — WebSocket Compatibility (viem + ethers) + Evidence Pack

Scope: NOORCHAIN 2.1 / noorcore JSON-RPC over WebSocket.
Goal: Prove wallet-grade compatibility for WebSocket clients using:
- viem (WS JSON-RPC request path)
- ethers v6 (WebSocketProvider: newHeads + logs)
Including follower semantics via FollowRPC proxy.

This runbook is written for the local mainnet-like pack:
- leader RPC: 127.0.0.1:8545
- follower RPC: 127.0.0.1:8546 (FollowRPC -> leader)
- WS uses the same host/port as JSON-RPC unless otherwise configured by your build.

Operational discipline:
- T1 = nodes only
- T2 = tooling only
- One command per step.
- Each step has an explicit gate (PASS/FAIL).
- Do not proceed if a gate fails.

Security posture:
- Never commit private keys.
- Do not paste private keys into logs, files, or screenshots.
- Evidence pack must not contain secrets.

---

## 0) Preconditions

Required:
- Local pack already running (leader + follower).
- M18 validated: follower WS proxy via FollowRPC is enabled and stable.
- M19.2 validated: WS hardening caps/backpressure/GC enabled.

Assumed paths:
- repo root: `/workspaces/noorchain-core`
- dApps/tooling workspace: `/workspaces/noorchain-core/dapps/curators-hub-v0`
- canonical submit script: `/workspaces/noorchain-core/scripts/submit-snapshot.mjs`

---

## 1) Private Key Gate (MANDATORY)

Rationale: Hardhat/viem/ethers-based scripts require a runtime private key.
If missing, signers may be empty and deploy/submit scripts may fail.

Format:
- 0x-prefixed hex
- Length: 66 characters (0x + 64 hex)

### Step 1 — Check key is set and valid (T2)

```bash
[[ "${NOOR_PRIVATE_KEY:-}" =~ ^0x[0-9a-fA-F]{64}$ ]] && echo "PK_GATE=PASS len=66" || echo "PK_GATE=FAIL (export NOOR_PRIVATE_KEY: 0x + 64 hex chars, len=66)"
Gate:

PASS: prints PK_GATE=PASS len=66

FAIL: stop and export the key at runtime (do not commit)

Step 2 — Export key safely (runtime only) (T2)
read -s -p "NOOR_PRIVATE_KEY (paste 0x + 64 hex): " NOOR_PRIVATE_KEY; echo; export NOOR_PRIVATE_KEY
Gate:

command returns without error (the key must not be echoed)

Step 3 — Re-check key format (T2)
[[ "${NOOR_PRIVATE_KEY:-}" =~ ^0x[0-9a-fA-F]{64}$ ]] && echo "PK_GATE=PASS len=${#NOOR_PRIVATE_KEY}" || (echo "PK_GATE=FAIL len=${#NOOR_PRIVATE_KEY}"; echo "Tip: must be 0x + 64 hex chars (66 total)")
Gate:

PASS: PK_GATE=PASS len=66

2) Node Inventory and Endpoints
Step 4 — Inventory running nodes (T1)
pgrep -a noorcore || echo NO_PROCESSES
Gate:

PASS: two processes exist

follower includes -rpc-addr 127.0.0.1:8546 and -follow-rpc http://127.0.0.1:8545

leader includes -rpc-addr 127.0.0.1:8545

Reference endpoints:

HTTP leader: http://127.0.0.1:8545

WS leader: ws://127.0.0.1:8545

HTTP follower: http://127.0.0.1:8546

WS follower: ws://127.0.0.1:8546

3) Registry Address Selection (No guessing; prove it)
Rule:

The PoSSRegistry address MUST be a deployed contract in the current environment.

Determine it by proving eth_getCode != 0x.

Candidate addresses:

Environment-specific deployed PoSSRegistry (preferred)

Default fallback used by some scripts may be outdated depending on the environment

Step 5 — Prove deployed registry by eth_getCode (T2)
for A in 0x560Ff18B151045561E9393bb8FaEA15C17BC02Fc 0xC9F398646E19778F2C3D9fF32bb75E5a99FD4E56; do
  echo "ADDR=$A"
  curl -s -H 'Content-Type: application/json' \
    -d "{\"jsonrpc\":\"2.0\",\"id\":1,\"method\":\"eth_getCode\",\"params\":[\"$A\",\"latest\"]}" \
    http://127.0.0.1:8545
  echo
done
Gate:

PASS: select the address whose result is NOT "0x".

Step 6 — Export registry address for deterministic targeting (T2)
Example (if the deployed registry is 0xC9F398...):

export NOOR_POSS_REGISTRY=0xC9F398646E19778F2C3D9fF32bb75E5a99FD4E56; echo "NOOR_POSS_REGISTRY=$NOOR_POSS_REGISTRY"
Gate:

PASS: prints the exported address.

4) Derive Topic0 from a Real Receipt (canonical)
Rationale:

Avoid ABI dependency for test gating.

Use the first emitted log topic0 from a known successful submitSnapshot receipt.

Step 7 — Trigger submitSnapshot and read receipt topics (T2)
Trigger:

NOOR_RPC_URL=http://127.0.0.1:8545 node /workspaces/noorchain-core/scripts/submit-snapshot.mjs
Gate:

PASS: prints SUBMIT_OK and shows the submitSnapshot tx: 0x...

Receipt summary (replace TX hash):

TX=0x<PASTE_TX_HASH>; \
curl -s -H 'Content-Type: application/json' \
  -d "{\"jsonrpc\":\"2.0\",\"id\":1,\"method\":\"eth_getTransactionReceipt\",\"params\":[\"$TX\"]}" \
  http://127.0.0.1:8545 | \
node -e 'let s="";process.stdin.on("data",d=>s+=d);process.stdin.on("end",()=>{const j=JSON.parse(s);const r=j.result||{};console.log("status=",r.status,"to=",r.to,"logs_len=",Array.isArray(r.logs)?r.logs.length:"NA"); if(Array.isArray(r.logs)&&r.logs[0]) console.log("log0_addr=",r.logs[0].address,"topic0=",r.logs[0].topics&&r.logs[0].topics[0]);});'
Gate:

PASS:

status= 0x1

to= <registry>

logs_len >= 1

topic0= 0x... (record it; used in WS filter)

5) viem — WebSocket Compatibility
All viem steps are executed from the tooling workspace where viem is installed.

Step 8 — Ensure viem is installed (T2)
cd /workspaces/noorchain-core/dapps/curators-hub-v0 && node -p "require('viem/package.json').version"
Gate:

PASS: prints a version (e.g., 2.43.3)

If missing:

cd /workspaces/noorchain-core/dapps/curators-hub-v0 && npm ci
Gate:

PASS: completes without error.

Step 9 — viem WS newHeads subscribe/unsubscribe via leader + follower (T2)
cd /workspaces/noorchain-core/dapps/curators-hub-v0 && node - <<'NODE'
(async () => {
  const viemPkg = require("viem/package.json");
  const { createPublicClient, webSocket } = require("viem");

  async function subUnsub(url) {
    const c = createPublicClient({ transport: webSocket(url) });
    const subId = await c.request({ method: "eth_subscribe", params: ["newHeads"] });
    const unsub = await c.request({ method: "eth_unsubscribe", params: [subId] });
    return { subId, unsub };
  }

  console.log(`VIEM_VERSION=${viemPkg.version}`);

  const L = await subUnsub("ws://127.0.0.1:8545");
  console.log(`LEADER_SUB_ID=${L.subId}`);
  console.log(`LEADER_UNSUB=${String(L.unsub)}`);

  const F = await subUnsub("ws://127.0.0.1:8546");
  console.log(`FOLLOWER_SUB_ID=${F.subId}`);
  console.log(`FOLLOWER_UNSUB=${String(F.unsub)}`);

  const pass =
    typeof L.subId === "string" && L.subId.startsWith("0x") && L.unsub === true &&
    typeof F.subId === "string" && F.subId.startsWith("0x") && F.unsub === true;

  console.log(`RESULT=${pass ? "PASS" : "FAIL"}`);
  process.exit(pass ? 0 : 1);
})();
NODE
Gate:

PASS: LEADER_UNSUB=true, FOLLOWER_UNSUB=true, RESULT=PASS

6) Raw WS logs subscribe + trigger tx (Evidence-oriented)
This step proves:

eth_subscribe(logs) works over WS

follower proxy streams logs consistently (FollowRPC semantics)

filter uses address + topic0 without ABI dependency

Step 10 — WS logs on follower + trigger submitSnapshot on leader (T2)
Set:

NOOR_POSS_REGISTRY already exported

TOPIC0 obtained from Step 7 receipt

Replace TOPIC0 in the command below.

cd /workspaces/noorchain-core && node - <<'NODE'
const { spawn } = require("child_process");

const REG = (process.env.NOOR_POSS_REGISTRY || "").trim();
if (!REG) { console.log("FAIL: NOOR_POSS_REGISTRY missing"); process.exit(2); }

const TOPIC0 = "0x<PASTE_TOPIC0_FROM_RECEIPT>";
const WSURL = "ws://127.0.0.1:8546"; // follower proxy

if (typeof WebSocket === "undefined") {
  console.log("FAIL: global WebSocket missing in this Node runtime");
  process.exit(2);
}

const ws = new WebSocket(WSURL);

let subId = null;
let notif = 0;
let txCode = null;
let txOut = "";
let txErr = "";

function send(o){ ws.send(JSON.stringify(o)); }

const timer = setTimeout(() => {
  try { ws.close(); } catch {}
  console.log(`SUB_ID=${subId || "NONE"}`);
  console.log(`NOTIF_TOTAL=${notif}`);
  console.log(`TX_EXIT=${txCode === null ? "NONE" : txCode}`);
  if (txOut.trim()) console.log("[TX_STDOUT_LAST]\n" + txOut.trim().split("\n").slice(-12).join("\n"));
  if (txErr.trim()) console.log("[TX_STDERR_LAST]\n" + txErr.trim().split("\n").slice(-12).join("\n"));
  console.log(`RESULT=${(subId && notif > 0) ? "PASS" : "FAIL"}`);
  process.exit((subId && notif > 0) ? 0 : 1);
}, 25000);

ws.onopen = () => {
  console.log(`REGISTRY=${REG}`);
  console.log(`TOPIC0=${TOPIC0}`);
  send({ jsonrpc:"2.0", id:1, method:"eth_subscribe", params:["logs", { address:[REG, REG.toLowerCase()], topics:[TOPIC0] }] });
};

ws.onmessage = (ev) => {
  const data = typeof ev.data === "string" ? ev.data : Buffer.from(ev.data).toString("utf8");
  let msg; try { msg = JSON.parse(data); } catch { return; }

  if (msg.id === 1) {
    if (msg.error) {
      console.log("SUBSCRIBE_ERROR=" + (msg.error.message || JSON.stringify(msg.error)));
      clearTimeout(timer);
      try { ws.close(); } catch {}
      console.log("RESULT=FAIL");
      process.exit(1);
    }
    subId = msg.result;
    console.log("SUB_ID=" + subId);

    const child = spawn(process.execPath, ["scripts/submit-snapshot.mjs"], {
      cwd: "/workspaces/noorchain-core",
      env: { ...process.env, NOOR_RPC_URL: "http://127.0.0.1:8545", NOOR_POSS_REGISTRY: REG },
      stdio: ["ignore", "pipe", "pipe"],
    });
    child.stdout.on("data", (b) => txOut += b.toString());
    child.stderr.on("data", (b) => txErr += b.toString());
    child.on("close", (c) => { txCode = c; console.log("TX_EXIT=" + c); });
    return;
  }

  if (msg.method === "eth_subscription" && msg.params && msg.params.subscription === subId) {
    notif += 1;
    const l = msg.params.result;
    console.log(`LOG#${notif} block=${l.blockNumber} tx=${l.transactionHash} idx=${l.logIndex} addr=${l.address}`);
    if (notif >= 1) {
      send({ jsonrpc:"2.0", id:2, method:"eth_unsubscribe", params:[subId] });
      clearTimeout(timer);
      try { ws.close(); } catch {}
      console.log(`SUB_ID=${subId}`);
      console.log(`NOTIF_TOTAL=${notif}`);
      console.log(`TX_EXIT=${txCode === null ? "PENDING" : txCode}`);
      console.log("RESULT=PASS");
      process.exit(0);
    }
  }
};

ws.onerror = (e) => console.log("WS_ERROR=" + (e?.message || String(e)));
NODE
Gate:

PASS:

SUB_ID=0x...

at least one LOG#1 ...

RESULT=PASS

7) ethers v6 — WebSocketProvider Compatibility
Step 11 — ethers WS newHeads + logs via follower (proxy) (T2)
cd /workspaces/noorchain-core && node - <<'NODE'
const { spawn } = require("child_process");

(async () => {
  const REG = (process.env.NOOR_POSS_REGISTRY || "").trim();
  if (!REG) { console.log("FAIL: NOOR_POSS_REGISTRY missing"); process.exit(2); }

  const TOPIC0 = "0x<PASTE_TOPIC0_FROM_RECEIPT>";
  const WSURL  = "ws://127.0.0.1:8546";

  let ethers;
  try { ethers = await import("ethers"); } catch { console.log("ETHERS=NOT_FOUND"); process.exit(2); }
  const { WebSocketProvider } = ethers;

  console.log(`ETHERS_VERSION=${ethers.version || "unknown"}`);
  console.log(`REGISTRY=${REG}`);
  console.log(`TOPIC0=${TOPIC0}`);
  console.log(`WSURL=${WSURL}`);

  const provider = new WebSocketProvider(WSURL);

  let blockSeen = false;
  let logSeen = false;
  let txExit = null;
  let txOut = "";
  let txErr = "";

  const filter = { address: [REG, REG.toLowerCase()], topics: [TOPIC0] };

  provider.on("block", (bn) => { blockSeen = true; console.log(`NEWHEAD block=${bn}`); });
  provider.on(filter, (log) => { logSeen = true; console.log(`LOG addr=${log.address} block=${log.blockNumber} tx=${log.transactionHash} idx=${log.index}`); });

  const child = spawn(process.execPath, ["scripts/submit-snapshot.mjs"], {
    cwd: "/workspaces/noorchain-core",
    env: { ...process.env, NOOR_RPC_URL: "http://127.0.0.1:8545", NOOR_POSS_REGISTRY: REG },
    stdio: ["ignore", "pipe", "pipe"],
  });
  child.stdout.on("data", (b) => txOut += b.toString());
  child.stderr.on("data", (b) => txErr += b.toString());
  child.on("close", (c) => { txExit = c; console.log(`TX_EXIT=${c}`); });

  const t0 = Date.now();
  while (Date.now() - t0 < 25000) {
    if (blockSeen && logSeen && txExit === 0) break;
    await new Promise((r) => setTimeout(r, 250));
  }

  try { provider.removeAllListeners(); } catch {}
  try { await provider.destroy?.(); } catch {}

  console.log(`BLOCK_SEEN=${blockSeen}`);
  console.log(`LOG_SEEN=${logSeen}`);
  console.log(`TX_EXIT=${txExit === null ? "NONE" : txExit}`);

  const pass = (blockSeen && logSeen && txExit === 0);
  console.log(`RESULT=${pass ? "PASS" : "FAIL"}`);
  process.exit(pass ? 0 : 1);
})();
NODE
Gate:

PASS:

at least one NEWHEAD block=...

at least one LOG addr=...

TX_EXIT=0

RESULT=PASS

8) Evidence Pack (Minimal)
Step 12 — Write evidence file (T2)
cd /workspaces/noorchain-core && mkdir -p artifacts/m20 && cat > artifacts/m20/EVIDENCE_M20_WS_COMPAT.txt <<'EOF'
NOORCHAIN 2.1 — M20 — WS Compatibility (viem + ethers) — Evidence Pack (minimal)

Fill:
- Date:
- Branch:
- Commit:
- Tags:

Endpoints
- HTTP leader:  http://127.0.0.1:8545
- WS leader:    ws://127.0.0.1:8545
- HTTP follower: http://127.0.0.1:8546
- WS follower:   ws://127.0.0.1:8546

Registry (proved by eth_getCode != 0x)
- NOOR_POSS_REGISTRY=<0x...>

Topic0 (from receipt.logs[0].topics[0])
- <0x...>

Tooling versions
- viem: <version>
- ethers: <version>
- node: <node -v>

Transcripts
- viem newHeads subscribe/unsubscribe: RESULT=PASS
- WS logs subscribe + tx trigger via follower: RESULT=PASS
- ethers WebSocketProvider newHeads + logs via follower: RESULT=PASS
EOF
echo "WROTE=artifacts/m20/EVIDENCE_M20_WS_COMPAT.txt"
Gate:

PASS: prints WROTE=artifacts/m20/EVIDENCE_M20_WS_COMPAT.txt

9) Completion Criteria (M20 PASS)
M20 is PASS when all are true:

Private Key Gate PASS (len=66).

Registry selection is proven by eth_getCode != 0x and exported via NOOR_POSS_REGISTRY.

Receipt proof: submitSnapshot emits >=1 log and topic0 is recorded.

viem WS newHeads subscribe/unsubscribe PASS on leader and follower.

WS logs subscribe + tx trigger PASS on follower (proxy).

ethers v6 WebSocketProvider receives newHeads and logs on follower (proxy) with tx trigger PASS.

Evidence file written under artifacts/m20/.

End.