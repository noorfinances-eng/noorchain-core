RUNBOOK — M18 WS Follower Proxy via FollowRPC (Mainnet-like)

Status: VALIDATED
Branch: main
Commit: f047430
Tag: M18-WS-FOLLOWER-PROXY-STABLE

0. Purpose

Validate that a follower node proxies WebSocket traffic to the leader via FollowRPC, providing on the follower endpoint:

eth_subscribe / eth_unsubscribe

Streaming parity for:

newHeads

logs (event streaming)

This runbook targets a controlled leader/follower topology.

1. Scope and invariants

Topology:

Leader HTTP/WS RPC: 127.0.0.1:8545

Follower HTTP/WS RPC: 127.0.0.1:8546 with -follow-rpc http://127.0.0.1:8545

WebSocket transport is served on the same RPC port (HTTP Upgrade).

Proxy mode is enabled when FollowRPC is non-empty.

One-command discipline: one command → one gate.

Terminals:

T1 = nodes only

T2 = tooling only

2. Prerequisites

Go toolchain available.

Node.js available.

Repo root: /workspaces/noorchain-core

Node dependency: ws available in the environment (used in tests below).

Script available (for log trigger):

scripts/submit-snapshot.mjs

3. Mandatory Private Key Gate (REQUIRED)

Tooling requires NOOR_PRIVATE_KEY exported at runtime (never committed).

Format: 0x-prefixed hex

Length: 66 characters

If the key is missing, submit/deploy scripts may fail (e.g., empty signers list).

4. Reference addresses

PoSSRegistry address used by scripts/submit-snapshot.mjs must be a deployed contract.

Default (script fallback, may be outdated depending on the environment):

0xC9F398646E19778F2C3D9fF32bb75E5a99FD4E56

Override (preferred for deterministic targeting):

NOOR_POSS_REGISTRY=<0x...> (or NOOR_REGISTRY=<0x...>)

This runbook uses NOOR_POSS_REGISTRY when set, otherwise falls back to the default.

5. Step-by-step
Step 1 — Stop any running nodes (T1)
pkill -INT noorcore || true


Gate: command returns (ports/processes are verified next).

Step 2 — Verify processes and ports are free (T1)
pgrep -a noorcore || echo NO_PROCESSES
ss -lntp | grep -E ':(30303|30304|8545|8546|8081|8082)\b' || echo PORTS_FREE


Gate:

NO_PROCESSES

PORTS_FREE

Step 3 — Build the binary (T1)
cd /workspaces/noorchain-core && go build -o noorcore ./core


Gate: build succeeds (no errors).

Step 4 — Start leader (T1)
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

ss -lntp | grep -E ':(30303|8545|8081)\b'

Step 5 — Start follower with FollowRPC (T1)
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

ss -lntp | grep -E ':(30304|8546|8082)\b'

Step 6 — HTTP sanity checks (T2)
curl -sS http://127.0.0.1:8545 -H 'content-type: application/json' --data '{"jsonrpc":"2.0","id":1,"method":"eth_chainId","params":[]}'
curl -sS http://127.0.0.1:8546 -H 'content-type: application/json' --data '{"jsonrpc":"2.0","id":2,"method":"eth_chainId","params":[]}'


Gate: both return a valid result (same chain id).

Step 7 — Private Key Gate (T2) (MANDATORY)
python - <<'PY'
import os
k=os.environ.get("NOOR_PRIVATE_KEY","")
print("NOOR_PRIVATE_KEY_LEN=", len(k))
print("OK" if (len(k)==66 and k.startswith("0x")) else "BAD")
PY


Gate:

NOOR_PRIVATE_KEY_LEN= 66

OK

If BAD, set it (runtime only), then re-run Step 7:

read -s -p "NOOR_PRIVATE_KEY (0x...): " NOOR_PRIVATE_KEY; echo; export NOOR_PRIVATE_KEY

Step 8 — Prove registry address is a contract (T2)
ADDR="${NOOR_POSS_REGISTRY:-0xC9F398646E19778F2C3D9fF32bb75E5a99FD4E56}"; \
echo "REG=$ADDR"; \
curl -sS http://127.0.0.1:8545 -H 'content-type: application/json' --data '{"jsonrpc":"2.0","id":1,"method":"eth_getCode","params":["'"$ADDR"'","latest"]}'


Gate: result is not "0x".

If "0x", export the correct deployed registry address (then re-run Step 8):

export NOOR_POSS_REGISTRY=0xYOUR_DEPLOYED_REGISTRY

Step 9 — Proxy proof (follower establishes upstream to leader) (T2)

This step opens a follower WS subscription, then prints any ESTAB TCP connection(s) to 127.0.0.1:8545 (upstream).

node -e '
const {execSync}=require("child_process");
const WS=require("ws");
const ws=new WS("ws://127.0.0.1:8546");
ws.on("open",()=>ws.send(JSON.stringify({jsonrpc:"2.0",id:1,method:"eth_subscribe",params:["newHeads"]})));
ws.on("message",(m)=>{
  const j=JSON.parse(String(m));
  if(j.id===1){
    console.log("SUBID", j.result);
    setTimeout(()=>{
      const out=execSync("ss -ntp | grep -E \"ESTAB .*127\\.0\\.0\\.1:8545\" || true",{shell:true}).toString();
      console.log("UPSTREAM_ESTAB_LINES_START");
      process.stdout.write(out || "NONE\n");
      console.log("UPSTREAM_ESTAB_LINES_END");
      ws.close();
    }, 800);
  }
});
ws.on("error",(e)=>console.log("WS_ERROR", String(e)));
'


Gate:

Prints SUBID 0x...

UPSTREAM_ESTAB_LINES_START ... END contains at least one ESTAB line (or, at minimum, does not show NONE if your proxy creates a TCP connection to leader).

Step 10 — WS streaming parity: newHeads on follower (T2)
node -e '
const WS=require("ws");
let got=0;
const ws=new WS("ws://127.0.0.1:8546");
const to=setTimeout(()=>{console.log("TIMEOUT_NO_HEAD_10S",{got}); try{ws.close();}catch{}}, 10000);
ws.on("open",()=>ws.send(JSON.stringify({jsonrpc:"2.0",id:1,method:"eth_subscribe",params:["newHeads"]})));
ws.on("message",(m)=>{
  const j=JSON.parse(String(m));
  if(j.id===1){ console.log("SUBID", j.result); return; }
  if(j.method==="eth_subscription"){
    got++;
    console.log("GOT_HEAD", {number:j.params?.result?.number, hash:j.params?.result?.hash});
    if(got>=1){
      clearTimeout(to);
      ws.send(JSON.stringify({jsonrpc:"2.0",id:2,method:"eth_unsubscribe",params:[j.params.subscription]}));
    }
  }
  if(j.id===2){
    console.log("UNSUB", j.result);
    try{ws.close();}catch{}
  }
});
ws.on("error",(e)=>console.log("WS_ERROR", String(e)));
'


Gate:

SUBID 0x...

GOT_HEAD ... (at least one)

UNSUB true

Step 11 — WS streaming parity: logs on follower (subscribe → tx → event) (T2)

Subscription must start before triggering the tx.

cd /workspaces/noorchain-core && node -e '
const {spawn}=require("child_process");
const WS=require("ws");

const REG = process.env.NOOR_POSS_REGISTRY
  || process.env.NOOR_REGISTRY
  || "0xC9F398646E19778F2C3D9fF32bb75E5a99FD4E56";

const pk = process.env.NOOR_PRIVATE_KEY || "";
if(pk.length!==66 || !pk.startsWith("0x")){ console.log("PK_BAD"); process.exit(1); }
console.log("PK_OK len=66");
console.log("REGISTRY", REG);

let subId=null;
let got=false;

const ws=new WS("ws://127.0.0.1:8546");
const to=setTimeout(()=>{console.log("TIMEOUT_NO_EVENT_90S"); try{ws.close();}catch{}; process.exit(2);}, 90000);

ws.on("open",()=>{
  ws.send(JSON.stringify({jsonrpc:"2.0",id:1,method:"eth_subscribe",params:["logs",{address:REG}]}));
});

ws.on("message",(m)=>{
  const j=JSON.parse(String(m));

  if(j.id===1){
    subId=j.result;
    console.log("SUBID", subId);

    // Trigger tx after subscription ACK.
    const child=spawn(process.execPath, ["scripts/submit-snapshot.mjs"], {
      env: {...process.env, NOOR_POSS_REGISTRY: REG},
      stdio: ["ignore","pipe","pipe"],
    });
    child.stdout.on("data",(d)=>process.stdout.write(d.toString()));
    child.stderr.on("data",(d)=>process.stderr.write(d.toString()));
    child.on("exit",(code)=>console.log("TX_SCRIPT_EXIT", code));
    return;
  }

  if(j.method==="eth_subscription" && !got){
    got=true;
    console.log("GOT_EVENT");
    console.log(JSON.stringify(j.params, null, 2));

    // Unsubscribe and close.
    ws.send(JSON.stringify({jsonrpc:"2.0",id:2,method:"eth_unsubscribe",params:[j.params.subscription]}));
    return;
  }

  if(j.id===2){
    console.log("UNSUB", j.result);
    clearTimeout(to);
    try{ws.close();}catch{}
    process.exit(0);
  }
});

ws.on("error",(e)=>console.log("WS_ERROR", String(e)));
'


Gate:

SUBID 0x...

Script prints SUBMIT_OK

GOT_EVENT received with:

params.result.address == REGISTRY

params.result.transactionHash matches the tx hash printed by the script

UNSUB true

6. Acceptance criteria (M18 PASS)

M18 is PASS when, on the follower endpoint (ws://127.0.0.1:8546) with -follow-rpc enabled:

eth_subscribe(newHeads) returns a non-null subId

at least one newHeads notification is received (eth_subscription)

eth_unsubscribe(subId) returns true

eth_subscribe(logs,{address:REG}) returns a non-null subId

after triggering submit-snapshot.mjs, a logs notification is received on the follower (eth_subscription)

eth_unsubscribe(subId) returns true

Proxy proof step shows follower establishes an upstream connection to leader while subscription is active