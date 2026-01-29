# RUNBOOK — M19.2 — WebSocket Hardening (Global Caps, Backpressure, GC)

Document status: operational runbook  
Audience: core operators / devs  
Scope: NOORCHAIN 2.1 / noorcore JSON-RPC over WebSocket (leader + follower)

## 0) Objective

Validate M19.2 hardening for the WebSocket server:

- Global connection/subscription caps
- Per-connection bounded outbox + backpressure behavior
- Garbage collection (GC) for idle connections/subscriptions
- Clean shutdown semantics (no 1006 / no uncontrolled drops)
- Multi-node behavior: leader (WS enabled) + follower (FollowRPC enabled)

This runbook is **validation-first**: every step has a gate.

## 1) Non-negotiable invariants

- Two nodes only:
  - Leader: RPC+WS `127.0.0.1:8545`, P2P `127.0.0.1:30303`, health `127.0.0.1:8081`
  - Follower: RPC+WS `127.0.0.1:8546`, P2P `127.0.0.1:30304`, health `127.0.0.1:8082`, FollowRPC `http://127.0.0.1:8545`
- Never run two `noorcore` on the same `-data-dir` (LevelDB LOCK).
- Operational discipline:
  - T1 = nodes only
  - T2 = tooling only
  - One command → one gate → next step

## 2) Preconditions

- Repo root: `/workspaces/noorchain-core`
- Binary build method (mandatory): `go build -o noorcore ./core`
- Node tooling:
  - `curl`, `ss`, `pgrep`
  - `node` available
  - Node package `ws` available (see step 6.A gate)

## 3) Ports and processes baseline

### 3.A Stop all nodes (T1)
```bash
pkill -INT noorcore || true
pgrep -a noorcore || echo NO_PROCESSES
Gate: output includes NO_PROCESSES.

3.B Confirm ports are free (T1)
ss -lntp | grep -E ':(30303|30304|8545|8546|8081|8082)\b' || echo PORTS_FREE
Gate: output includes PORTS_FREE.

4) Build
4.A Build noorcore (T2)
cd /workspaces/noorchain-core && go build -o noorcore ./core
Gate: command exits 0.

4.B Sanity check binary (T2)
cd /workspaces/noorchain-core && ls -l ./noorcore && ./noorcore -h >/dev/null && echo HELP_EXIT=$?
Gate: HELP_EXIT=0.

5) Start nodes
5.A Start leader (T1)
cd /workspaces/noorchain-core && ./noorcore \
  -role leader \
  -chain-id noorchain-2-1-local \
  -data-dir ./data/node1 \
  -p2p-addr 127.0.0.1:30303 \
  -rpc-addr 127.0.0.1:8545 \
  -health-addr 127.0.0.1:8081 \
  > /tmp/noorcore_node1.log 2>&1 &
Gate (T1):

pgrep -a noorcore | grep -E -- '-role leader .* -rpc-addr 127\.0\.0\.1:8545'
ss -lntp | grep -E ':(30303|8545|8081)\b'
curl -s http://127.0.0.1:8545 -H 'content-type: application/json' --data '{"jsonrpc":"2.0","id":1,"method":"eth_chainId","params":[]}'
curl -s http://127.0.0.1:8545 -H 'content-type: application/json' --data '{"jsonrpc":"2.0","id":2,"method":"eth_blockNumber","params":[]}'
Leader process visible

Ports listening on 30303/8545/8081

eth_chainId returns 0x849

eth_blockNumber returns a hex value

5.B Start follower (T1)
cd /workspaces/noorchain-core && ./noorcore \
  -role follower \
  -chain-id noorchain-2-1-local \
  -data-dir ./data/node2 \
  -p2p-addr 127.0.0.1:30304 \
  -rpc-addr 127.0.0.1:8546 \
  -health-addr 127.0.0.1:8082 \
  -follow-rpc http://127.0.0.1:8545 \
  -boot-peers 127.0.0.1:30303 \
  > /tmp/noorcore_node2.log 2>&1 &
Gate (T1):

pgrep -a noorcore | grep -E -- '-role follower .* -rpc-addr 127\.0\.0\.1:8546'
ss -lntp | grep -E ':(30304|8546|8082)\b'
curl -s http://127.0.0.1:8546 -H 'content-type: application/json' --data '{"jsonrpc":"2.0","id":1,"method":"eth_chainId","params":[]}'
curl -s http://127.0.0.1:8546 -H 'content-type: application/json' --data '{"jsonrpc":"2.0","id":2,"method":"eth_blockNumber","params":[]}'
curl -s http://127.0.0.1:8545 -H 'content-type: application/json' --data '{"jsonrpc":"2.0","id":3,"method":"eth_blockNumber","params":[]}'
Follower process visible

Ports listening on 30304/8546/8082

Follower chainId 0x849

Follower blockNumber equals leader blockNumber (or quickly converges)

6) WebSocket gates
6.A Tooling gate: Node ws module present (T2)
node -e 'require("ws"); console.log("WS_OK")'
Gate: prints WS_OK.

If missing, install locally (T2) and re-run the gate:

npm i ws --no-save
6.B Upgrade handshake (T2)
Leader:

curl -i -N \
  -H 'Connection: Upgrade' \
  -H 'Upgrade: websocket' \
  -H 'Sec-WebSocket-Version: 13' \
  -H 'Sec-WebSocket-Key: dGhlIHNhbXBsZSBub25jZQ==' \
  http://127.0.0.1:8545 | head -n 12
Follower:

curl -i -N \
  -H 'Connection: Upgrade' \
  -H 'Upgrade: websocket' \
  -H 'Sec-WebSocket-Version: 13' \
  -H 'Sec-WebSocket-Key: dGhlIHNhbXBsZSBub25jZQ==' \
  http://127.0.0.1:8546 | head -n 12
Gate: both return HTTP/1.1 101 Switching Protocols.

7) Stress / backpressure validation
7.A High-rate request blast (no crash, clean close) (T2)
Run:

node -e '
const N=20000;
const WS=require("ws");

async function blast(url){
  return new Promise((resolve)=>{
    const ws=new WS(url);
    let sent=0, t0=Date.now();
    ws.on("open", ()=>{
      function loop(){
        for(let i=0;i<200;i++){
          ws.send(JSON.stringify({jsonrpc:"2.0",id:sent,method:"eth_chainId",params:[]}));
          sent++;
          if(sent>=N){
            console.log("SENT_OK", {url, sent, ms: Date.now()-t0});
            ws.close(1000);
            console.log("CLIENT_CLOSE", {url, sent});
            return;
          }
        }
        setImmediate(loop);
      }
      loop();
    });
    ws.on("error",(e)=>{ console.log("WS_ERR", {url, sent, err:String(e)}); });
    ws.on("close",(code,reason)=>{
      console.log("SERVER_CLOSED", {url, code, reason:String(reason), sent, ms: Date.now()-t0});
      resolve();
    });
  });
}

(async()=>{
  await blast("ws://127.0.0.1:8545");
  await blast("ws://127.0.0.1:8546");
})();
'
Gate: for both URLs:

SENT_OK ...

CLIENT_CLOSE ...

SERVER_CLOSED ... code: 1000

No 1006, no [object ErrorEvent], no hang.

7.B Backpressure timeout (forced non-reader) — optional but recommended (T2)
This test tries to block reading to force server-side outbox pressure.
Expected behavior: server closes with a controlled close code (often 1013) and not 1006.

node -e '
const WS=require("ws");
async function backpressure(url){
  return new Promise((resolve)=>{
    const ws=new WS(url);
    let sent=0, t0=Date.now();
    ws.on("open", ()=>{
      // Attempt to stop reading frames to create write backpressure.
      if(ws._socket && ws._socket.pause) ws._socket.pause();

      const payload=JSON.stringify({jsonrpc:"2.0",id:1,method:"eth_chainId",params:[]});
      function loop(){
        for(let i=0;i<500;i++){
          ws.send(payload);
          sent++;
        }
        setImmediate(loop);
      }
      loop();
    });
    ws.on("close",(code,reason)=>{
      console.log("BACKPRESSURE_CLOSED", {url, code, reason:String(reason), sent, ms: Date.now()-t0});
      resolve();
    });
    ws.on("error",(e)=>{ console.log("BACKPRESSURE_ERR", {url, sent, err:String(e)}); });
    setTimeout(()=>{ try{ ws.close(1000); }catch{} }, 5000);
  });
}
(async()=>{ await backpressure("ws://127.0.0.1:8545"); })();
'
Gate:

Connection closes within a few seconds.

Close is controlled (not 1006).

7.C Oversize message close — optional (T2)
Expected: close with 1009 (“message too big”) or immediate controlled closure.

node -e '
const WS=require("ws");
const ws=new WS("ws://127.0.0.1:8545");
ws.on("open", ()=>{
  const big="x".repeat((1<<20)+1024);
  ws.send(big);
});
ws.on("close",(code,reason)=>console.log("OVERSIZE_CLOSED",{code,reason:String(reason)}));
ws.on("error",(e)=>console.log("OVERSIZE_ERR",String(e)));
'
Gate: close is controlled (not 1006).

8) Follower semantics gate (WS methods)
If follower is configured with FollowRPC, certain WS methods may be restricted.
Validate that follower returns a JSON-RPC error for eth_subscribe / eth_unsubscribe (leader-only policy).

node -e '
const WS=require("ws");
const ws=new WS("ws://127.0.0.1:8546");
ws.on("open", ()=>ws.send(JSON.stringify({jsonrpc:"2.0",id:1,method:"eth_subscribe",params:["newHeads"]})));
ws.on("message",(m)=>{ console.log("FOLLOWER_SUBSCRIBE_RESP", m.toString()); ws.close(1000); });
ws.on("close",(c)=>process.exit(0));
'
Gate: response contains an error and indicates leader-only behavior (e.g. code -32000 with message "leader-only").

9) GC liveness gate (no leaks / no crash)
9.A Observe node stability during WS activity (T1)
After completing section 7, confirm nodes are still running and RPC works:

pgrep -a noorcore
curl -s http://127.0.0.1:8545 -H 'content-type: application/json' --data '{"jsonrpc":"2.0","id":9,"method":"eth_blockNumber","params":[]}'
curl -s http://127.0.0.1:8546 -H 'content-type: application/json' --data '{"jsonrpc":"2.0","id":10,"method":"eth_blockNumber","params":[]}'
Gate: both nodes still alive; RPC responds.

10) Evidence pack (minimum)
Capture and store:

Current commit:

git rev-parse HEAD

Binary build proof:

ls -l ./noorcore

Node startup parameters:

pgrep -a noorcore

WS stress output:

console output from section 7.A (leader + follower)

Logs (tail):

tail -n 80 /tmp/noorcore_node1.log

tail -n 80 /tmp/noorcore_node2.log

11) Clean shutdown
11.A Stop nodes (T1)
pkill -INT noorcore || true
pgrep -a noorcore || echo NO_PROCESSES
ss -lntp | grep -E ':(30303|30304|8545|8546|8081|8082)\b' || echo PORTS_FREE
Gate: NO_PROCESSES and PORTS_FREE.

12) Pass criteria
M19.2 is PASS if:

WS blast (20k) passes on leader and follower with clean 1000 close.

No uncontrolled 1006.

Backpressure + oversize tests (if executed) close in a controlled way.

Nodes remain alive and RPC responsive after WS stress.

Follower leader-only policy behaves as specified for restricted WS methods (if policy is enabled).


