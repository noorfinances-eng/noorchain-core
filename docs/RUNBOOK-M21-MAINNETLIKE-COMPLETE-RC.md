# RUNBOOK — M21 Mainnet-like Complete (RC Freeze)

**Project:** NOORCHAIN 2.1 (sovereign EVM L1)  
**Branch:** `main`  
**Prereq Tag:** `M20-WS-COMPAT-VIEM-ETHERS-STABLE`  
**Outcome Tag (this runbook):** `M21-MAINNETLIKE-COMPLETE-RC`  
**Last validated:** 2026-01-29 (UTC)  
**Scope:** Final RC gates for “Mainnet-like Complete”: HTTP+WS parity (leader/follower), restart invariants, final evidence pack, freeze tag.

---

## 0) Non-negotiable invariants

- **Terminal discipline**
  - **T1 = nodes only** (start/stop/run logs)
  - **T2 = tooling only** (curl/node/scripts/git)
  - One command per step. Execute → verify gate → next step.

- **Leader/Follower**
  - Leader RPC: `127.0.0.1:8545` (HTTP + WS)
  - Follower RPC: `127.0.0.1:8546` (HTTP + WS, **proxy via `-follow-rpc`**)

- **Chain ID**
  - JSON-RPC `eth_chainId` MUST return `0x849` (2121).

- **Private Key Gate (mandatory before any tooling tx)**
  - `NOOR_PRIVATE_KEY` must be exported at runtime, **0x-prefixed, len=66**.
  - Never commit/store keys.

- **PoSS Registry env (mandatory for scripts)**
  - `NOOR_POSS_REGISTRY=0xC9F398646E19778F2C3D9fF32bb75E5a99FD4E56`

- **Safety**
  - Never run two `noorcore` processes on the same `-data-dir` (LevelDB LOCK).
  - Avoid piping `submit-snapshot.mjs` into `head` (causes Node `EPIPE`). Log to file instead.

---

## 1) Parameters (local RC)

- **Leader**
  - `-data-dir ./data/node1`
  - `-p2p-addr 127.0.0.1:30303`
  - `-rpc-addr 127.0.0.1:8545`
  - `-health-addr 127.0.0.1:8081`
- **Follower**
  - `-data-dir ./data/node2`
  - `-p2p-addr 127.0.0.1:30304`
  - `-rpc-addr 127.0.0.1:8546`
  - `-health-addr 127.0.0.1:8082`
  - `-boot-peers 127.0.0.1:30303`
  - `-follow-rpc http://127.0.0.1:8545`

---

## 2) Step-by-step (M21 RC gates)

### Step 1 — Repo gate (T2)
**Command (T2):**
```bash
cd /workspaces/noorchain-core && git status -sb && git tag --list "M20-WS-COMPAT-VIEM-ETHERS-STABLE" && git rev-parse --short HEAD
Gate:

## main...origin/main (or explicit ahead/behind)

tag M20-WS-COMPAT-VIEM-ETHERS-STABLE exists

HEAD printed (short hash)

Step 2 — Clean process/ports gate (T1)
Command (T1):

bash
Copier le code
pgrep -a noorcore || echo NO_PROCESSES; ss -lntp | grep -E ':(30303|30304|8545|8546|8081|8082)\b' || echo PORTS_FREE
Gate:

NO_PROCESSES and PORTS_FREE OR you see exactly what is already running (do not kill blindly).

Step 3 — Build gate (T2)
Command (T2):

bash
Copier le code
cd /workspaces/noorchain-core && go build -o noorcore ./core && ls -l ./noorcore && ./noorcore -h >/dev/null && echo HELP_EXIT=$?
Gate:

./noorcore exists

HELP_EXIT=0

Step 4 — Start leader (T1)
Command (T1):

bash
Copier le code
cd /workspaces/noorchain-core && ./noorcore -chain-id noorchain-2-1-local -data-dir ./data/node1 -role leader -p2p-addr 127.0.0.1:30303 -rpc-addr 127.0.0.1:8545 -health-addr 127.0.0.1:8081 > /tmp/noorcore_node1.log 2>&1 &
Gate (next step validates).

Step 5 — Leader up gate (T1)
Command (T1):

bash
Copier le code
pgrep -a noorcore | head -n 5; ss -lntp | grep -E ':(30303|8545|8081)\b'; curl -s -X POST http://127.0.0.1:8545 -H 'content-type: application/json' --data '{"jsonrpc":"2.0","id":1,"method":"eth_chainId","params":[]}' ; echo; curl -s -X POST http://127.0.0.1:8545 -H 'content-type: application/json' --data '{"jsonrpc":"2.0","id":2,"method":"eth_blockNumber","params":[]}' ; echo
Gate:

ports LISTEN on 30303/8545/8081

eth_chainId = 0x849

eth_blockNumber returns hex (any value)

Step 6 — Start follower (T1)
Command (T1):

bash
Copier le code
cd /workspaces/noorchain-core && ./noorcore -chain-id noorchain-2-1-local -data-dir ./data/node2 -role follower -p2p-addr 127.0.0.1:30304 -rpc-addr 127.0.0.1:8546 -health-addr 127.0.0.1:8082 -boot-peers 127.0.0.1:30303 -follow-rpc http://127.0.0.1:8545 > /tmp/noorcore_node2.log 2>&1 &
Gate (next step validates).

Step 7 — Follower up gate (T1)
Command (T1):

bash
Copier le code
pgrep -a noorcore | head -n 10; ss -lntp | grep -E ':(30304|8546|8082)\b'; curl -s -X POST http://127.0.0.1:8546 -H 'content-type: application/json' --data '{"jsonrpc":"2.0","id":1,"method":"eth_chainId","params":[]}' ; echo; curl -s -X POST http://127.0.0.1:8546 -H 'content-type: application/json' --data '{"jsonrpc":"2.0","id":2,"method":"eth_blockNumber","params":[]}' ; echo
Gate:

ports LISTEN on 30304/8546/8082

eth_chainId = 0x849

follower eth_blockNumber returns hex (close to leader)

Step 8 — HTTP parity gate (roots/bloom at latest) (T2)
Command (T2):

bash
Copier le code
H=$(curl -s -X POST http://127.0.0.1:8545 -H 'content-type: application/json' --data '{"jsonrpc":"2.0","id":1,"method":"eth_blockNumber","params":[]}' | sed -n 's/.*"result":"\([^"]*\)".*/\1/p'); DATA=$(printf '{"jsonrpc":"2.0","id":2,"method":"eth_getBlockByNumber","params":["%s",false]}' "$H"); echo "HEIGHT=$H"; for u in 8545 8546; do echo "=== :$u ==="; curl -s -X POST http://127.0.0.1:$u -H 'content-type: application/json' --data "$DATA" | tr -d '\n' | grep -Eo '"number":"[^"]+"|"hash":"[^"]+"|"stateRoot":"[^"]+"|"receiptsRoot":"[^"]+"|"logsBloom":"[^"]+"'; echo; done
Gate:

number/hash/stateRoot/receiptsRoot/logsBloom identical on :8545 and :8546

Step 9 — WS parity gate: newHeads (leader vs follower) (T2)
Command (T2):

bash
Copier le code
node - <<'NODE'
const N = 5, TIMEOUT_MS = 20000;
const EP = [{name:"LEADER",url:"ws://127.0.0.1:8545"},{name:"FOLLOWER",url:"ws://127.0.0.1:8546"}];
function subHeads(ep){return new Promise((res,rej)=>{const out=[];let subId=null,done=false;const ws=new WebSocket(ep.url);
const timer=setTimeout(()=>{if(done)return;done=true;try{ws.close()}catch{};rej(new Error(`${ep.name}: timeout ${out.length}/${N}`));},TIMEOUT_MS);
ws.onopen=()=>ws.send(JSON.stringify({jsonrpc:"2.0",id:1,method:"eth_subscribe",params:["newHeads"]}));
ws.onmessage=(ev)=>{let msg;try{msg=JSON.parse(ev.data)}catch{return;}
if(msg.id===1&&msg.result){subId=msg.result;return;}
if(msg.method==="eth_subscription"&&msg.params&&msg.params.subscription===subId){
const r=msg.params.result||{}; if(r.number&&r.hash){out.push({number:r.number,hash:r.hash});
if(out.length>=N&&!done){done=true;clearTimeout(timer);try{ws.send(JSON.stringify({jsonrpc:"2.0",id:2,method:"eth_unsubscribe",params:[subId]}))}catch{};try{ws.close()}catch{};res(out);}}}};
ws.onerror=()=>{if(done)return;done=true;clearTimeout(timer);rej(new Error(`${ep.name}: ws error`));};
});}
(async()=>{const [L,F]=await Promise.all(EP.map(subHeads));
const li=L.map(x=>x.hash.toLowerCase()), fi=F.map(x=>x.hash.toLowerCase());
const start=li.indexOf(fi[0]); let pass=start>=0;
if(pass){for(let i=0;i<fi.length;i++){if(li[start+i]!==fi[i]){pass=false;break;}}}
console.log("LEADER:",L); console.log("FOLLOWER:",F);
console.log(pass?"RESULT: PASS (aligned by hash)":"RESULT: FAIL (mismatch)");
process.exit(pass?0:1);
})().catch(e=>{console.error("RESULT: ERROR",e.message||e);process.exit(2);});
NODE
Gate:

RESULT: PASS (aligned by hash)

Step 10 — Private Key Gate (mandatory) (T2)
Command (T2):

bash
Copier le code
read -s -p "NOOR_PRIVATE_KEY (0x + 64 hex): " NOOR_PRIVATE_KEY; echo; export NOOR_PRIVATE_KEY; export NOOR_POSS_REGISTRY=0xC9F398646E19778F2C3D9fF32bb75E5a99FD4E56; node - <<'NODE'
const pk=process.env.NOOR_PRIVATE_KEY||""; console.log("PK_LEN=",pk.length,"PK_OK=",/^0x[0-9a-fA-F]{64}$/.test(pk));
console.log("NOOR_POSS_REGISTRY=",process.env.NOOR_POSS_REGISTRY||"");
NODE
Gate:

PK_LEN= 66 and PK_OK= true

registry address printed and correct

Step 11 — WS parity gate: logs (listener background) (T2)
Command (T2):

bash
Copier le code
rm -f /tmp/wslogs_parity.log; node - <<'NODE' > /tmp/wslogs_parity.log 2>&1 &
const REG=(process.env.NOOR_POSS_REGISTRY||"0x").toLowerCase();
const EP=[{name:"LEADER",url:"ws://127.0.0.1:8545"},{name:"FOLLOWER",url:"ws://127.0.0.1:8546"}];
function start(ep){const ws=new WebSocket(ep.url);let subId=null;
ws.onopen=()=>ws.send(JSON.stringify({jsonrpc:"2.0",id:1,method:"eth_subscribe",params:["logs",{address:REG}]}));
ws.onmessage=(ev)=>{let msg;try{msg=JSON.parse(ev.data)}catch{return;}
if(msg.id===1&&msg.result&&!subId){subId=msg.result;console.log(`SUB_OK ${ep.name} ${subId}`);return;}
if(msg.method==="eth_subscription"&&msg.params&&msg.params.subscription===subId){
const r=msg.params.result||{}; if((r.address||"").toLowerCase()!==REG) return;
console.log(JSON.stringify({ep:ep.name,blockNumber:r.blockNumber,blockHash:r.blockHash,tx:r.transactionHash,topic0:Array.isArray(r.topics)?r.topics[0]:null}));
}};
ws.onerror=()=>console.log(`WS_ERR ${ep.name}`); ws.onclose=()=>console.log(`WS_CLOSE ${ep.name}`);}
for(const ep of EP) start(ep); setInterval(()=>{},1<<30);
NODE
sleep 0.5; grep -E 'SUB_OK (LEADER|FOLLOWER)' /tmp/wslogs_parity.log || (echo "NO_SUB_OK"; tail -n 60 /tmp/wslogs_parity.log)
Gate:

SUB_OK LEADER ... and SUB_OK FOLLOWER ...

Step 12 — Trigger 2 PoSS events (2× submitSnapshot) (T2)
Command (T2):

bash
Copier le code
cd /workspaces/noorchain-core && node scripts/submit-snapshot.mjs && node scripts/submit-snapshot.mjs && echo SUBMIT_DONE=$?
Gate:

each run ends with SUBMIT_OK

SUBMIT_DONE=0

Step 13 — WS logs parity proof (T2)
Command (T2):

bash
Copier le code
echo "=== COUNT JSON EVENTS ==="; grep -c '^{.*"ep":"LEADER".*}$' /tmp/wslogs_parity.log; grep -c '^{.*"ep":"FOLLOWER".*}$' /tmp/wslogs_parity.log; echo; echo "=== LAST EVENTS ==="; grep -E '^\{.*"ep":"(LEADER|FOLLOWER)".*\}$' /tmp/wslogs_parity.log | tail -n 20
Gate:

LEADER >= 2 and FOLLOWER >= 2

matching blockHash/tx/topic0 pairs for each event (same seen by both)

Step 14 — Restart invariants (stop clean) (T1)
Command (T1):

bash
Copier le code
pkill -INT noorcore || true; sleep 0.5; pgrep -a noorcore || echo NO_PROCESSES; ss -lntp | grep -E ':(30303|30304|8545|8546|8081|8082)\b' || echo PORTS_FREE
Gate:

NO_PROCESSES

PORTS_FREE

Step 15 — Restart invariants (start leader+follower) (T1)
Command (T1):

bash
Copier le code
cd /workspaces/noorchain-core && ./noorcore -chain-id noorchain-2-1-local -data-dir ./data/node1 -role leader -p2p-addr 127.0.0.1:30303 -rpc-addr 127.0.0.1:8545 -health-addr 127.0.0.1:8081 > /tmp/noorcore_node1.log 2>&1 & ./noorcore -chain-id noorchain-2-1-local -data-dir ./data/node2 -role follower -p2p-addr 127.0.0.1:30304 -rpc-addr 127.0.0.1:8546 -health-addr 127.0.0.1:8082 -boot-peers 127.0.0.1:30303 -follow-rpc http://127.0.0.1:8545 > /tmp/noorcore_node2.log 2>&1 & echo STARTED
Gate (next step validates).

Step 16 — Post-restart up + parity gate (T1)
Command (T1):

bash
Copier le code
pgrep -a noorcore; ss -lntp | grep -E ':(30303|30304|8545|8546|8081|8082)\b'; curl -s -X POST http://127.0.0.1:8545 -H 'content-type: application/json' --data '{"jsonrpc":"2.0","id":1,"method":"eth_chainId","params":[]}' ; echo; curl -s -X POST http://127.0.0.1:8546 -H 'content-type: application/json' --data '{"jsonrpc":"2.0","id":2,"method":"eth_chainId","params":[]}' ; echo; curl -s -X POST http://127.0.0.1:8545 -H 'content-type: application/json' --data '{"jsonrpc":"2.0","id":3,"method":"eth_blockNumber","params":[]}' ; echo; curl -s -X POST http://127.0.0.1:8546 -H 'content-type: application/json' --data '{"jsonrpc":"2.0","id":4,"method":"eth_blockNumber","params":[]}' ; echo
Gate:

2 processes up, ports LISTEN

chainId 0x849 both

blockNumber returned both

Step 17 — Tooling proof post-restart (avoid EPIPE) (T2)
Command (T2):

bash
Copier le code
cd /workspaces/noorchain-core && node scripts/submit-snapshot.mjs > /tmp/m21_submit_after_restart.log 2>&1; sed -n '1,3p' /tmp/m21_submit_after_restart.log; echo '---'; tail -n 6 /tmp/m21_submit_after_restart.log
Gate:

top shows Signer(curator): ...

bottom shows receipt status: 1 and SUBMIT_OK

Step 18 — Evidence pack (minimal) (T2)
Command (T2):

bash
Copier le code
{
  echo "=== NOORCHAIN 2.1 M21 RC EVIDENCE (minimal) ===";
  date -u +"UTC %Y-%m-%dT%H:%M:%SZ";
  echo;
  echo "== git ==";
  cd /workspaces/noorchain-core && git rev-parse --abbrev-ref HEAD && git rev-parse --short HEAD && git tag --points-at HEAD || true;
  echo;
  echo "== leader/follower pids ==";
  pgrep -a noorcore;
  echo;
  echo "== chainId/blockNumber ==";
  for u in 8545 8546; do
    echo "--- :$u ---";
    curl -s -X POST http://127.0.0.1:$u -H 'content-type: application/json' --data '{"jsonrpc":"2.0","id":1,"method":"eth_chainId","params":[]}' ; echo;
    curl -s -X POST http://127.0.0.1:$u -H 'content-type: application/json' --data '{"jsonrpc":"2.0","id":2,"method":"eth_blockNumber","params":[]}' ; echo;
  done
  echo;
  echo "== latest block roots/bloom (parity snapshot) ==";
  H=$(curl -s -X POST http://127.0.0.1:8545 -H 'content-type: application/json' --data '{"jsonrpc":"2.0","id":1,"method":"eth_blockNumber","params":[]}' | sed -n 's/.*"result":"\([^"]*\)".*/\1/p');
  echo "HEIGHT=$H";
  DATA=$(printf '{"jsonrpc":"2.0","id":3,"method":"eth_getBlockByNumber","params":["%s",false]}' "$H");
  for u in 8545 8546; do
    echo "--- :$u ---";
    curl -s -X POST http://127.0.0.1:$u -H 'content-type: application/json' --data "$DATA" \
      | tr -d '\n' \
      | grep -Eo '"number":"[^"]+"|"hash":"[^"]+"|"stateRoot":"[^"]+"|"receiptsRoot":"[^"]+"|"logsBloom":"[^"]+"';
    echo;
  done
  echo;
  echo "== PoSS env + last submit ==";
  echo "NOOR_POSS_REGISTRY=$NOOR_POSS_REGISTRY";
  tail -n 6 /tmp/m21_submit_after_restart.log;
  echo;
  echo "== WS parity artifacts (counts) ==";
  [ -f /tmp/wslogs_parity.log ] && {
    echo "WS_LEADER_EVENTS=$(grep -c '^{.*\"ep\":\"LEADER\".*}$' /tmp/wslogs_parity.log)";
    echo "WS_FOLLOWER_EVENTS=$(grep -c '^{.*\"ep\":\"FOLLOWER\".*}$' /tmp/wslogs_parity.log)";
    echo "WS_LAST_EVENTS:";
    grep -E '^\{.*"ep":"(LEADER|FOLLOWER)".*\}$' /tmp/wslogs_parity.log | tail -n 6;
  } || echo "NO_WS_LOG_FILE";
} | tee /tmp/M21_RC_EVIDENCE_MINIMAL.txt
echo "EVIDENCE_WRITTEN=/tmp/M21_RC_EVIDENCE_MINIMAL.txt"
Gate:

prints EVIDENCE_WRITTEN=/tmp/M21_RC_EVIDENCE_MINIMAL.txt

evidence shows parity + post-restart submit OK

Step 19 — FREEZE (tag RC + push) (T2)
Command (T2):

bash
Copier le code
cd /workspaces/noorchain-core && git tag -a M21-MAINNETLIKE-COMPLETE-RC -m "M21 RC: HTTP+WS parity leader/follower, restart invariants PASS. Evidence: /tmp/M21_RC_EVIDENCE_MINIMAL.txt" && git push origin M21-MAINNETLIKE-COMPLETE-RC
Gate:

push confirms new tag on origin

3) Definition of Done (M21)
M21 is considered DONE when all are true:

Leader and follower both up (HTTP+WS)

eth_chainId=0x849 and blockNumber progresses

HTTP parity holds for getBlockByNumber(latest) fields: hash/stateRoot/receiptsRoot/logsBloom

WS newHeads parity holds (aligned by hash)

WS logs parity holds for PoSSRegistry address (both endpoints receive matching events)

Restart invariants hold (stop clean, start clean, parity still holds)

Tooling proof post-restart: submitSnapshot mined (SUBMIT_OK)

Evidence pack written: /tmp/M21_RC_EVIDENCE_MINIMAL.txt

Freeze tag pushed: M21-MAINNETLIKE-COMPLETE-RC

4) Post-freeze rule
After tagging M21-MAINNETLIKE-COMPLETE-RC:

No code changes unless a new milestone is explicitly opened.

Only allowed: documentation, release notes, archiving evidence packs, operational hardening notes.

Private keys remain runtime-only; never stored or committed.