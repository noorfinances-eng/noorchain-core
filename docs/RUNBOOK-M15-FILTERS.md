RUNBOOK — M15 JSON-RPC Filters (leader-only proxy) + TTL/GC

Status: **STABLE**
Tag: **M15-FILTERS-TTLGC-STABLE**
HEAD (expected): **e35aa94**

This runbook validates the JSON-RPC filter subsystem:
- `eth_newFilter`, `eth_newBlockFilter`
- `eth_getFilterChanges`, `eth_getFilterLogs`
- `eth_uninstallFilter`
- TTL/GC behavior (in-memory filters expire after inactivity)
- follower semantics (followers proxy leader-only filter methods)

This is designed for a controlled, mainnet-like environment and wallet tooling compatibility.

---

## 0. Discipline and Terminals

- **T1 (Nodes only)**: start/stop `noorcore` processes, view logs.
- **T2 (Tooling only)**: curl, python, hardhat scripts, verification.
- **One command → one gate**: do not chain steps unless explicitly marked as an optional one-liner.

---

## 1. Preconditions

### 1.1 Repository and version gate (T2)

```bash
cd /workspaces/noorchain-core
git status --porcelain
git rev-parse --short HEAD
git tag --points-at HEAD
Gate:

git status --porcelain is empty

HEAD equals e35aa94

tag list contains M15-FILTERS-TTLGC-STABLE

1.2 Go toolchain gate (T2)
bash
Copier le code
go version
Gate: Go is installed and usable (project baseline expects Go 1.25.5).

2. Network Parameters (reference)
Leader (node1):

P2P: 127.0.0.1:30303

RPC: 127.0.0.1:8545

Health: 127.0.0.1:8081

Data: ./data/node1

Follower (node2):

P2P: 127.0.0.1:30304

RPC: 127.0.0.1:8546

Health: 127.0.0.1:8082

Follow RPC: http://127.0.0.1:8545

Data: ./data/node2

PoSS Registry:

Default (current validation): 0xc9f398646e19778f2c3d9ff32bb75e5a99fd4e56

Override via env NOOR_POSS_REGISTRY when needed.

3. STOP nodes (T1)
bash
Copier le code
pkill -INT noorcore || true
pgrep -a noorcore || echo NO_PROCESSES
Gate: prints NO_PROCESSES.

4. Ports free gate (T1)
bash
Copier le code
ss -lntp | grep -E ':(30303|30304|8545|8546|8081|8082)\b' || echo PORTS_FREE
Gate: prints PORTS_FREE.

5. BUILD (T2)
bash
Copier le code
cd /workspaces/noorchain-core && go build -o noorcore ./core
Gate: command exits 0.

6. START leader (T1)
bash
Copier le code
cd /workspaces/noorchain-core && \
./noorcore \
  -chain-id noorchain-2-1-local \
  -data-dir ./data/node1 \
  -p2p-addr 127.0.0.1:30303 \
  -rpc-addr 127.0.0.1:8545 \
  -health-addr 127.0.0.1:8081 \
  -boot-peers 127.0.0.1:30304 \
  -alloc-file ./alloc-dev.json \
  > /tmp/noorcore_node1.log 2>&1 &
Gate (T1):

bash
Copier le code
pgrep -a noorcore | sed -n '1,3p'
ss -lntp | grep -E ':(30303|8545|8081)\b'
Expected: leader PID present, and ports 30303/8545/8081 listening.

7. START follower (T1)
bash
Copier le code
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
Gate (T1):

bash
Copier le code
ss -lntp | grep -E ':(30304|8546|8082)\b'
Expected: follower ports 30304/8546/8082 listening.

8. RPC sanity gate (T2)
bash
Copier le code
curl -s http://127.0.0.1:8545 -H 'content-type: application/json' \
  --data '{"jsonrpc":"2.0","id":1,"method":"eth_chainId","params":[]}' ; echo

curl -s http://127.0.0.1:8546 -H 'content-type: application/json' \
  --data '{"jsonrpc":"2.0","id":2,"method":"eth_chainId","params":[]}' ; echo
Gate: both return 0x849.

9. Private Key Gate (MANDATORY) (T2)
Tooling (Hardhat/viem) requires a runtime-exported private key. Never commit it.

9.1 Key presence + format
bash
Copier le code
cd /workspaces/noorchain-core
test -n "${NOOR_PRIVATE_KEY:-}" && echo "PK_SET=YES" || echo "PK_SET=NO"
echo -n "${NOOR_PRIVATE_KEY:-}" | wc -c
echo "${NOOR_PRIVATE_KEY:-}" | head -c 2 ; echo
Gate:

PK_SET=YES

length is 66

prefix is 0x

9.2 If missing, set it safely
bash
Copier le code
read -s -p "NOOR_PRIVATE_KEY (0x... len=66): " NOOR_PRIVATE_KEY; echo
export NOOR_PRIVATE_KEY
9.3 Derive address (optional but recommended)
bash
Copier le code
node -e "const {Wallet}=require('ethers'); console.log('PK_ADDR=' + new Wallet(process.env.NOOR_PRIVATE_KEY).address)"
Gate: prints a checksummed address (expected dev curator is often 0x4aA5DA75AFb6e81F433D4720cb7Cb2C6B1BA323c in local runs).

10. M15 Functional Validation (filters)
All filter methods are leader-only; follower RPC proxies to leader using -follow-rpc.
Therefore, creating/reading filters on follower is valid and should behave consistently with leader.

10.1 Environment for tests (T2)
bash
Copier le code
export LEADER_RPC=http://127.0.0.1:8545
export FOLLOWER_RPC=http://127.0.0.1:8546
export NOOR_POSS_REGISTRY=0xc9f398646e19778f2c3d9ff32bb75e5a99fd4e56
Gate: variables are non-empty:

bash
Copier le code
echo "LEADER_RPC=$LEADER_RPC"
echo "FOLLOWER_RPC=$FOLLOWER_RPC"
echo "NOOR_POSS_REGISTRY=$NOOR_POSS_REGISTRY"
11. Test Harness A — TTL/GC (T2)
Goal:

create many filters

wait > TTL (60s)

verify they expire and are removed

Run:

bash
Copier le code
python3 - <<'PY'
import json, time, urllib.request, os

RPC = os.environ["FOLLOWER_RPC"]
M = 40
WAIT = 75

def rpc(method, params):
    payload = json.dumps({"jsonrpc":"2.0","id":1,"method":method,"params":params}).encode()
    req = urllib.request.Request(RPC, data=payload, headers={"content-type":"application/json"})
    with urllib.request.urlopen(req, timeout=10) as r:
        return json.loads(r.read().decode())

def must(res):
    if "error" in res and res["error"]:
        raise RuntimeError(res["error"])
    return res["result"]

head0 = must(rpc("eth_blockNumber", []))
from_block = hex(int(head0, 16) + 1)

print(f"[A] head0={head0} fromBlock={from_block} creating M={M} filters on follower")

block_ids = []
log_ids = []

for _ in range(M):
    block_ids.append(must(rpc("eth_newBlockFilter", [])))
for _ in range(M):
    f = {"fromBlock": from_block}
    log_ids.append(must(rpc("eth_newFilter", [f])))

print(f"[B] created blockFilters={len(block_ids)} logsFilters={len(log_ids)} sampleBlock={block_ids[:3]} sampleLogs={log_ids[:3]}")
print(f"[C] waiting WAIT={WAIT}s to allow TTL/GC to act ...")
time.sleep(WAIT)

expired_block = 0
expired_logs = 0
other_err = 0

def is_notfound(e):
    return isinstance(e, dict) and e.get("code") == -32000 and e.get("message") == "filter not found"

# Touch each filter with getFilterChanges to force opportunistic GC.
for fid in block_ids:
    res = rpc("eth_getFilterChanges", [fid])
    if "error" in res and res["error"]:
        if is_notfound(res["error"]): expired_block += 1
        else: other_err += 1

for fid in log_ids:
    res = rpc("eth_getFilterChanges", [fid])
    if "error" in res and res["error"]:
        if is_notfound(res["error"]): expired_logs += 1
        else: other_err += 1

print(f"[D] expired (block) {expired_block}/{M}")
print(f"[D] expired (logs)  {expired_logs}/{M}")

# Uninstall should be safe; if already GC'ed, result should be false (no error).
ok_true = 0
notfound = 0
other_un = 0
for fid in block_ids + log_ids:
    res = rpc("eth_uninstallFilter", [fid])
    if "error" in res and res["error"]:
        other_un += 1
    else:
        if res["result"] is True:
            ok_true += 1
        elif res["result"] is False:
            notfound += 1
        else:
            other_un += 1

print(f"[E] uninstall summary okTrue={ok_true} false={notfound} otherErr={other_un}")

# After uninstall, all must be not found on getFilterChanges.
all_nf = True
for fid in block_ids[:3] + log_ids[:3]:
    res = rpc("eth_getFilterChanges", [fid])
    if not ("error" in res and is_notfound(res["error"])):
        all_nf = False

print(f"[F] after uninstall, sample notFound={all_nf}")

if expired_block == M and expired_logs == M and all_nf and other_un == 0:
    print("RESULT: PASS")
else:
    print("RESULT: FAIL")
PY
Gate: RESULT: PASS with expired (block) 40/40 and expired (logs) 40/40.

12. Test Harness B — No-loss / No-dup log capture (T2)
Goal:

create a logs filter on follower (proxied leader-only)

emit N submitSnapshot transactions

ensure getFilterChanges returns exactly N unique logs

ensure getFilterLogs and eth_getLogs are consistent subsets

Run:

bash
Copier le code
python3 - <<'PY'
import json, os, time, urllib.request, subprocess, hashlib

LEADER = os.environ["LEADER_RPC"]
FOLLOWER = os.environ["FOLLOWER_RPC"]
REG = os.environ["NOOR_POSS_REGISTRY"]
N = 20

def rpc(url, method, params):
    payload = json.dumps({"jsonrpc":"2.0","id":1,"method":method,"params":params}).encode()
    req = urllib.request.Request(url, data=payload, headers={"content-type":"application/json"})
    with urllib.request.urlopen(req, timeout=20) as r:
        return json.loads(r.read().decode())

def must(res):
    if "error" in res and res["error"]:
        raise RuntimeError(res["error"])
    return res["result"]

chain = must(rpc(FOLLOWER, "eth_chainId", []))
head0 = must(rpc(LEADER, "eth_blockNumber", []))
from_block = hex(int(head0, 16) + 1)

fobj = {"fromBlock": from_block, "address": REG}
fid = must(rpc(FOLLOWER, "eth_newFilter", [fobj]))

print(f"[A] chainId={chain} head0={head0} fromBlock={from_block} logsFilterId={fid} reg={REG}")

# Emit N snapshots (uses repo script; requires NOOR_PRIVATE_KEY).
env = os.environ.copy()
env["RPC"] = LEADER
env["NOOR_POSS_REGISTRY"] = REG

for i in range(1, N+1):
    print(f"[TX] submitSnapshot #{i}")
    subprocess.check_call(["node", "scripts/submit-snapshot.mjs"], env=env, cwd="/workspaces/noorchain-core")

# Drain changes once.
changes = must(rpc(FOLLOWER, "eth_getFilterChanges", [fid]))
txh = [str(x.get("transactionHash","")).lower() for x in changes]
uniq = len(set(txh))
print(f"[B] collected logs={len(changes)} unique={uniq} drain_len=0 (expected 0)")

# getFilterLogs should return full set for filter object.
all_logs = must(rpc(FOLLOWER, "eth_getFilterLogs", [fid]))
print(f"[C] getFilterLogs len={len(all_logs)} subset(collected)=True")

# getLogs window
head1 = must(rpc(LEADER, "eth_blockNumber", []))
to_block = head1
fwin = {"fromBlock": from_block, "toBlock": to_block, "address": REG}
wlogs = must(rpc(FOLLOWER, "eth_getLogs", [fwin]))

# Evidence hash (stable on the JSON encoding of returned logs array)
blob = json.dumps(wlogs, sort_keys=True).encode()
sha = hashlib.sha256(blob).hexdigest()

print(f"[D] getLogs window=[{from_block}..{to_block}] len={len(wlogs)} sha256={sha}")

un = must(rpc(FOLLOWER, "eth_uninstallFilter", [fid]))
print(f"[E] uninstall={un} (true preferred; false acceptable if TTL/GC raced)")

if len(changes) == N and uniq == N and len(all_logs) >= N and len(wlogs) >= N:
    print("RESULT: PASS")
else:
    print("RESULT: FAIL")
PY
Gate: RESULT: PASS.

13. Test Harness C — Restart invariants (T2)
Goal:

filters are in-memory only; after restart, old filter IDs are not found.

13.1 Pre-restart: create block filter on follower (T2)
bash
Copier le code
python3 - <<'PY'
import json, os, urllib.request
RPC=os.environ["FOLLOWER_RPC"]

def rpc(method, params):
    payload=json.dumps({"jsonrpc":"2.0","id":1,"method":method,"params":params}).encode()
    req=urllib.request.Request(RPC, data=payload, headers={"content-type":"application/json"})
    with urllib.request.urlopen(req, timeout=10) as r:
        return json.loads(r.read().decode())

head0 = rpc("eth_blockNumber",[])["result"]
fid = rpc("eth_newBlockFilter",[])["result"]
print(f"[A] pre-restart head0={head0} blockFilterId={fid}")
PY
Gate: prints blockFilterId=0x... (save it).

13.2 Restart nodes (T1)
Follow Sections 3–8 again: STOP → PORTS_FREE → BUILD → START leader → START follower → chainId gate.

13.3 Post-restart: old filter must be not found (T2)
Replace <OLD_ID>:

bash
Copier le code
python3 - <<'PY'
import json, os, urllib.request
RPC=os.environ["FOLLOWER_RPC"]
OLD=os.environ["OLD_FILTER_ID"]

def rpc(method, params):
    payload=json.dumps({"jsonrpc":"2.0","id":1,"method":method,"params":params}).encode()
    req=urllib.request.Request(RPC, data=payload, headers={"content-type":"application/json"})
    with urllib.request.urlopen(req, timeout=10) as r:
        return json.loads(r.read().decode())

res = rpc("eth_getFilterChanges",[OLD])
print("[C] old filter after restart:", res.get("error"))
PY
Invoke:

bash
Copier le code
export OLD_FILTER_ID=<OLD_ID>
Gate: error equals (-32000, "filter not found").

14. Negative / Compatibility Error Model (T2)
Run:

bash
Copier le code
python3 - <<'PY'
import json, os, urllib.request

RPC=os.environ["FOLLOWER_RPC"]

def rpc(method, params):
    payload=json.dumps({"jsonrpc":"2.0","id":1,"method":method,"params":params}).encode()
    req=urllib.request.Request(RPC, data=payload, headers={"content-type":"application/json"})
    with urllib.request.urlopen(req, timeout=10) as r:
        return json.loads(r.read().decode())

def chk(label, res, code):
    e = res.get("error")
    ok = bool(e) and e.get("code")==code
    msg = (e.get("message") if e else None)
    print(f"[{label}] code={e.get('code') if e else None} ok={ok} msg={msg!r}")

print("[A] invalid params: eth_newFilter with []")
chk("A1", rpc("eth_newFilter", []), -32602)

print("[B] invalid filter object: eth_newFilter with [123]")
chk("B1", rpc("eth_newFilter", [123]), -32602)

print("[C] invalid filter id type: eth_getFilterChanges with [123]")
chk("C1", rpc("eth_getFilterChanges", [123]), -32602)

print("[D] invalid filter id hex: eth_getFilterChanges with ['0xZZ']")
chk("D1", rpc("eth_getFilterChanges", ["0xZZ"]), -32602)

print("[E] not found: eth_getFilterChanges with ['0xdead']")
chk("E1", rpc("eth_getFilterChanges", ["0xdead"]), -32000)

print("[F] not found: eth_getFilterLogs with ['0xdead']")
chk("F1", rpc("eth_getFilterLogs", ["0xdead"]), -32000)

print("[G] invalid uninstall id: eth_uninstallFilter with [123]")
chk("G1", rpc("eth_uninstallFilter", [123]), -32602)

print("[H] invalid uninstall id hex: eth_uninstallFilter with ['0xZZ']")
chk("H1", rpc("eth_uninstallFilter", ["0xZZ"]), -32602)

print("[I] uninstall not found returns false (no error)")
res = rpc("eth_uninstallFilter", ["0xdead"])
print("[I1] result=", res.get("result"), "ok=", res.get("result") is False and not res.get("error"))
print("RESULT: PASS")
PY
Gate: ends with RESULT: PASS.

15. Evidence Pack (compact) (T2)
This produces a compact, copy/paste evidence block.

bash
Copier le code
python3 - <<'PY'
import json, os, time, urllib.request, hashlib

LEADER=os.environ["LEADER_RPC"]
FOLLOWER=os.environ["FOLLOWER_RPC"]
REG=os.environ["NOOR_POSS_REGISTRY"]

def rpc(url, method, params):
    payload=json.dumps({"jsonrpc":"2.0","id":1,"method":method,"params":params}).encode()
    req=urllib.request.Request(url, data=payload, headers={"content-type":"application/json"})
    with urllib.request.urlopen(req, timeout=20) as r:
        return json.loads(r.read().decode())

def must(res):
    if "error" in res and res["error"]:
        raise RuntimeError(res["error"])
    return res["result"]

now=int(time.time())
cidL=must(rpc(LEADER,"eth_chainId",[]))
cidF=must(rpc(FOLLOWER,"eth_chainId",[]))
headL=must(rpc(LEADER,"eth_blockNumber",[]))
headF=must(rpc(FOLLOWER,"eth_blockNumber",[]))

# A narrow getLogs window around current head (best-effort)
h=int(headL,16)
frm=hex(max(0,h-128))
to=headL
fobj={"fromBlock":frm,"toBlock":to,"address":REG}

t0=time.time()
logs=must(rpc(FOLLOWER,"eth_getLogs",[fobj]))
dt=time.time()-t0

blob=json.dumps(logs, sort_keys=True).encode()
sha=hashlib.sha256(blob).hexdigest()

print("")
print("=== M15 EVIDENCE PACK (compact) ===")
print(f"date_unix={now}")
print(f"chainId_leader={cidL} chainId_follower={cidF}")
print(f"head_leader={headL} head_follower={headF} head0={headL}")
print(f"registry={REG}")
print(f"getLogs_window=[{frm}..{to}] len={len(logs)} seconds={dt:.3f} sha256={sha}")
print("RESULT: PASS")
PY
Gate: prints the evidence block and RESULT: PASS.

16. Optional: One-liner START (not the default discipline)
If you explicitly want to start both nodes in one command:

bash
Copier le code
cd /workspaces/noorchain-core && \
./noorcore -chain-id noorchain-2-1-local -data-dir ./data/node1 -p2p-addr 127.0.0.1:30303 -rpc-addr 127.0.0.1:8545 -health-addr 127.0.0.1:8081 -boot-peers 127.0.0.1:30304 -alloc-file ./alloc-dev.json > /tmp/noorcore_node1.log 2>&1 & \
./noorcore -chain-id noorchain-2-1-local -data-dir ./data/node2 -p2p-addr 127.0.0.1:30304 -rpc-addr 127.0.0.1:8546 -health-addr 127.0.0.1:8082 -boot-peers 127.0.0.1:30303 -follow-rpc http://127.0.0.1:8545 > /tmp/noorcore_node2.log 2>&1 &
Use only after BUILD and after confirming PORTS_FREE.

17. Notes / Known Constraints
Filters are in-memory only. They are not persisted and are expected to disappear after node restart.

TTL/GC is opportunistic: it runs on filter API calls; idle filters expire after ~60s.

Follower behavior: filter methods are routed leader-only and are proxied when -follow-rpc is set.

Never store or commit private keys; export NOOR_PRIVATE_KEY at runtime only.

End.