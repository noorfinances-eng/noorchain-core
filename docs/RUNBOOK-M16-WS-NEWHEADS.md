0) Discipline et terminaux

T1 (Nodes only) : start/stop noorcore, logs.

T2 (Tooling only) : curl/node/python, validations.

One command → one gate : ne pas chaîner sauf section “one-liners” explicitement.

1) Preconditions
1.1 Repo / tag gate (T2)
cd /workspaces/noorchain-core
git status --porcelain
git rev-parse --short HEAD
git tag --points-at HEAD


Gate

git status --porcelain vide

HEAD = a9b5fa8

tags incluent M16-WS-NEWHEADS-STABLE

1.2 Go version gate (T2)
go version


Gate : Go utilisable (baseline projet : Go 1.25.5).

2) Paramètres réseau (référence)

Leader (node1)

P2P: 127.0.0.1:30303

RPC: 127.0.0.1:8545

Health: 127.0.0.1:8081

Data: ./data/node1

Follower (node2)

P2P: 127.0.0.1:30304

RPC: 127.0.0.1:8546

Health: 127.0.0.1:8082

Follow RPC: http://127.0.0.1:8545

Data: ./data/node2

WS endpoint :

Leader: ws://127.0.0.1:8545/ws

Follower: ws://127.0.0.1:8546/ws

3) STOP nodes (T1)
pkill -INT noorcore || true
pgrep -a noorcore || echo NO_PROCESSES


Gate : imprime NO_PROCESSES.

4) Ports free gate (T1)
ss -lntp | grep -E ':(30303|30304|8545|8546|8081|8082)\b' || echo PORTS_FREE


Gate : imprime PORTS_FREE.

5) BUILD (T2)
cd /workspaces/noorchain-core && go build -o noorcore ./core


Gate : exit 0.

6) START leader (T1)
cd /workspaces/noorchain-core && \
./noorcore \
  -chain-id noorchain-2-1-local \
  -data-dir ./data/node1 \
  -p2p-addr 127.0.0.1:30303 \
  -rpc-addr 127.0.0.1:8545 \
  -health-addr 127.0.0.1:8081 \
  -boot-peers 127.0.0.1:30304 \
  > /tmp/noorcore_node1.log 2>&1 &


Gate (T1)

pgrep -a noorcore | sed -n '1,3p'
ss -lntp | grep -E ':(30303|8545|8081)\b'


Expected : PID présent + ports LISTEN.

7) START follower (T1)
cd /workspaces/noorchain-core && \
./noorcore \
  -chain-id noorchain-2-1-local \
  -data-dir ./data/node2 \
  -p2p-addr 127.0.0.1:30304 \
  -rpc-addr 127.0.0.1:8546 \
  -health-addr 127.0.0.1:8082 \
  -boot-peers 127.0.0.1:30303 \
  -follow-rpc http://127.0.0.1:8545 \
  > /tmp/noorcore_node2.log 2>&1 &


Gate (T1)

ss -lntp | grep -E ':(30304|8546|8082)\b'


Expected : ports LISTEN.

8) RPC sanity gate (T2)
curl -s http://127.0.0.1:8545 -H 'content-type: application/json' \
  --data '{"jsonrpc":"2.0","id":1,"method":"eth_chainId","params":[]}' ; echo

curl -s http://127.0.0.1:8546 -H 'content-type: application/json' \
  --data '{"jsonrpc":"2.0","id":2,"method":"eth_chainId","params":[]}' ; echo


Gate : les deux renvoient 0x849.

9) WS Upgrade gate (leader) (T2)

Objectif : confirmer handshake WS (HTTP 101). Le curl ne lit pas de frames WS ensuite → timeout attendu.

curl -i -s -N \
  -H "Connection: Upgrade" \
  -H "Upgrade: websocket" \
  -H "Host: 127.0.0.1:8545" \
  -H "Origin: http://127.0.0.1" \
  -H "Sec-WebSocket-Key: SGVsbG8sIHdvcmxkIQ==" \
  -H "Sec-WebSocket-Version: 13" \
  http://127.0.0.1:8545/ws | head -n 20


Gate : voit HTTP/1.1 101 Switching Protocols + headers Upgrade: websocket.

10) WS newHeads subscribe/unsubscribe (leader) (T2)
10.1 One-shot Node gate (leader)
node - <<'NODE'
const ws = new WebSocket("ws://127.0.0.1:8545/ws");

let subId = null;
let unsubSent = false;

function send(obj) {
  ws.send(JSON.stringify(obj));
}

function unsub() {
  if (!subId || unsubSent) return;
  unsubSent = true;
  send({ jsonrpc: "2.0", id: 2, method: "eth_unsubscribe", params: [subId] });
}

const kill = setTimeout(() => {
  console.log("RESULT=TIMEOUT");
  try { ws.close(); } catch {}
  process.exit(2);
}, 5000);

ws.onopen = () => {
  send({ jsonrpc: "2.0", id: 1, method: "eth_subscribe", params: ["newHeads"] });
};

ws.onmessage = (ev) => {
  const msg = JSON.parse(ev.data.toString());

  if (msg.id === 1) {
    console.log("SUB_RESP=" + JSON.stringify(msg));
    subId = msg.result || null;
    if (!subId) {
      clearTimeout(kill);
      ws.close();
      process.exit(1);
    }
    return;
  }

  if (msg.method === "eth_subscription") {
    console.log("NOTIF=" + JSON.stringify(msg));
    unsub();
    return;
  }

  if (msg.id === 2) {
    console.log("UNSUB_RESP=" + JSON.stringify(msg));
    clearTimeout(kill);
    ws.close();
    process.exit(0);
  }
};

ws.onerror = (e) => {
  console.log("WS_ERROR=" + (e?.message || "unknown"));
};
NODE


Gate (PASS attendu) :

SUB_RESP contient "result":"0x..." (ex "0x1")

NOTIF contient method:"eth_subscription" avec un result.number non nul

UNSUB_RESP contient "result":true

11) Follower semantics gate (WS leader-only until M18) (T2)
node - <<'NODE'
const ws = new WebSocket("ws://127.0.0.1:8546/ws");
const kill = setTimeout(() => { console.log("RESULT=TIMEOUT"); process.exit(2); }, 4000);

ws.onopen = () => {
  ws.send(JSON.stringify({ jsonrpc:"2.0", id:1, method:"eth_subscribe", params:["newHeads"] }));
};

ws.onmessage = (ev) => {
  const msg = JSON.parse(ev.data.toString());
  console.log("RESP=" + JSON.stringify(msg));
  clearTimeout(kill);
  ws.close();

  const e = msg.error;
  if (e && e.code === -32000 && e.message === "leader-only") {
    console.log("RESULT=PASS");
    process.exit(0);
  }
  console.log("RESULT=FAIL");
  process.exit(1);
};

ws.onerror = (e) => {
  console.log("WS_ERROR=" + (e?.message || "unknown"));
};
NODE


Gate : RESULT=PASS (erreur -32000 "leader-only").

12) Evidence Pack (compact) (T2)
python3 - <<'PY'
import json, time, urllib.request, hashlib

def rpc(url, method, params):
    payload = json.dumps({"jsonrpc":"2.0","id":1,"method":method,"params":params}).encode()
    req = urllib.request.Request(url, data=payload, headers={"content-type":"application/json"})
    with urllib.request.urlopen(req, timeout=10) as r:
        return json.loads(r.read().decode())

def must(res):
    if res.get("error"):
        raise SystemExit("RPC_ERROR " + str(res["error"]))
    return res["result"]

now=int(time.time())
cid=must(rpc("http://127.0.0.1:8545","eth_chainId",[]))
head=must(rpc("http://127.0.0.1:8545","eth_blockNumber",[]))
blob=(str(cid)+str(head)).encode()
sha=hashlib.sha256(blob).hexdigest()

print("")
print("=== M16 EVIDENCE PACK (compact) ===")
print(f"date_unix={now}")
print(f"chainId={cid}")
print(f"head={head}")
print(f"sha256(chainId||head)={sha}")
print("RESULT: PASS")
PY


Gate : imprime RESULT: PASS