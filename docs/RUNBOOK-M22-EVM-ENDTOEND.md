# RUNBOOK — M22 EVM End-to-End (Deploy → Write → Read → Logs) + Follower Parity

Status: stable  
Scope: local mainnet-like pack (leader + follower), JSON-RPC HTTP  
Terminals: **T1 = nodes only**, **T2 = tooling only**  
Discipline: **one command → one gate** (do not chain commands during validation)

---

## 0) Objective

Validate an end-to-end Ethereum-compatible execution path:

- Deploy a minimal contract
- Execute a state-changing transaction (write)
- Read state via:
  - `eth_call`
  - `eth_getCode`
  - `eth_getStorageAt`
- Observe events via:
  - `eth_getLogs` (topic + address filter)
- Confirm transaction/receipt retrieval via:
  - `eth_getTransactionByHash`
  - `eth_getTransactionReceipt`
- Confirm **null semantics** for non-existent hashes:
  - Leader returns `result: null`
  - Follower returns `result: null` (**explicit result:null required by some clients**)
- Confirm **leader/follower parity** for reads and logs.

---

## 1) Preconditions

- Repo: `noorchain-core`
- Branch: `main` (recommended)
- Binaries built with `go build -o noorcore ./core`
- Ports:
  - Leader: `127.0.0.1:30303` (p2p), `127.0.0.1:8545` (rpc), `127.0.0.1:8081` (health)
  - Follower: `127.0.0.1:30304` (p2p), `127.0.0.1:8546` (rpc), `127.0.0.1:8082` (health)
- Data dirs:
  - Leader: `./data/node1`
  - Follower: `./data/node2`
- Tooling:
  - Node.js + npm
  - Hardhat (via `npx`)
- **Mandatory environment gates**:
  - `NOOR_PRIVATE_KEY` (0x-prefixed, len=66, runtime-only, never committed)
  - `NOOR_POSS_REGISTRY=0xC9F398646E19778F2C3D9fF32bb75E5a99FD4E56`

---

## 2) Stop nodes and verify ports are free (T1)

### Step 2.1 — Stop
```bash
pkill -INT noorcore || true
Gate: no processes remain

bash
Copier le code
pgrep -a noorcore || echo NO_PROCESSES
Step 2.2 — Ports free
bash
Copier le code
ss -lntp | grep -E ':(30303|30304|8545|8546|8081|8082)\b' || echo PORTS_FREE
Gate: PORTS_FREE

3) Build (T1)
bash
Copier le code
cd /workspaces/noorchain-core && go build -o noorcore ./core
Gate: binary exists and help exits

bash
Copier le code
cd /workspaces/noorchain-core && ls -l ./noorcore && ./noorcore -h >/dev/null && echo BUILD_OK
4) Start leader + follower (T1)
Step 4.1 — Start leader
bash
Copier le code
cd /workspaces/noorchain-core && ./noorcore \
  -role leader \
  -chain-id noorchain-2-1-local \
  -data-dir ./data/node1 \
  -p2p-addr 127.0.0.1:30303 \
  -rpc-addr 127.0.0.1:8545 \
  -health-addr 127.0.0.1:8081 \
  -boot-peers 127.0.0.1:30304 \
  > /tmp/noorcore_node1.log 2>&1 &
Gate: leader process + ports

bash
Copier le code
pgrep -a noorcore | grep -E -- '-role leader' || echo LEADER_MISSING
bash
Copier le code
ss -lntp | grep -E ':(30303|8545|8081)\b' || echo LEADER_PORTS_MISSING
Step 4.2 — Start follower
bash
Copier le code
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
Gate: follower process + ports

bash
Copier le code
pgrep -a noorcore | grep -E -- '-role follower' || echo FOLLOWER_MISSING
bash
Copier le code
ss -lntp | grep -E ':(30304|8546|8082)\b' || echo FOLLOWER_PORTS_MISSING
5) Basic RPC parity (T2)
Step 5.1 — chainId
Leader:

bash
Copier le code
curl -s -X POST 127.0.0.1:8545 -H 'content-type: application/json' \
  --data '{"jsonrpc":"2.0","id":1,"method":"eth_chainId","params":[]}' ; echo
Follower:

bash
Copier le code
curl -s -X POST 127.0.0.1:8546 -H 'content-type: application/json' \
  --data '{"jsonrpc":"2.0","id":2,"method":"eth_chainId","params":[]}' ; echo
Gate: both return 0x849

Step 5.2 — blockNumber
Leader:

bash
Copier le code
curl -s -X POST 127.0.0.1:8545 -H 'content-type: application/json' \
  --data '{"jsonrpc":"2.0","id":3,"method":"eth_blockNumber","params":[]}' ; echo
Follower:

bash
Copier le code
curl -s -X POST 127.0.0.1:8546 -H 'content-type: application/json' \
  --data '{"jsonrpc":"2.0","id":4,"method":"eth_blockNumber","params":[]}' ; echo
Gate: follower equals leader (same hex)

6) Mandatory environment gates (T2)
Step 6.1 — Private Key Gate (runtime-only)
Never print the key. Only validate the shape.

bash
Copier le code
( printf "NOOR_PRIVATE_KEY_len=%s\n" "${#NOOR_PRIVATE_KEY}"; \
  [[ "${#NOOR_PRIVATE_KEY}" -eq 66 && "${NOOR_PRIVATE_KEY:0:2}" == "0x" ]] && echo "PK_GATE=PASS" || echo "PK_GATE=FAIL" )
If FAIL (enter silently):

bash
Copier le code
read -s -p "NOOR_PRIVATE_KEY (0x..., len 66): " NOOR_PRIVATE_KEY; echo; export NOOR_PRIVATE_KEY
Re-check gate:

bash
Copier le code
( printf "NOOR_PRIVATE_KEY_len=%s\n" "${#NOOR_PRIVATE_KEY}"; \
  [[ "${#NOOR_PRIVATE_KEY}" -eq 66 && "${NOOR_PRIVATE_KEY:0:2}" == "0x" ]] && echo "PK_GATE=PASS" || echo "PK_GATE=FAIL" )
Step 6.2 — PoSS Registry Gate (exact address)
bash
Copier le code
export NOOR_POSS_REGISTRY=0xC9F398646E19778F2C3D9fF32bb75E5a99FD4E56
bash
Copier le code
( printf "NOOR_POSS_REGISTRY=%s len=%s\n" "${NOOR_POSS_REGISTRY}" "${#NOOR_POSS_REGISTRY}"; \
  [[ "${NOOR_POSS_REGISTRY}" == "0xC9F398646E19778F2C3D9fF32bb75E5a99FD4E56" ]] && echo "REG_ADDR=PASS" || echo "REG_ADDR=FAIL" )
7) Hardhat E2E deploy + write + read + logs (T2)
Step 7.1 — Prepare workspace
bash
Copier le code
mkdir -p /tmp/m22hh && cd /tmp/m22hh
Ensure ESM package type:

bash
Copier le code
cd /tmp/m22hh && npm pkg set type="module" >/dev/null && node -e 'console.log("PKG_TYPE=" + (require("./package.json").type || ""))'
Gate: PKG_TYPE=module

Step 7.2 — Compile + run script
Assumes the project already contains:

hardhat.config.mjs

scripts/deploy_and_test.mjs

a network named noorcore pointing to http://127.0.0.1:8545

bash
Copier le code
cd /tmp/m22hh && npx hardhat compile --config hardhat.config.mjs
Gate: compile succeeds

Run and capture output:

bash
Copier le code
cd /tmp/m22hh && npx hardhat run scripts/deploy_and_test.mjs \
  --network noorcore \
  --config hardhat.config.mjs \
  --show-stack-traces | tee /tmp/m22_hh_out.txt
Gate: script prints all of:

DEPLOYER=0x...

ADDR=0x...

SET_TX=0x...

BLOCK=0x...

GET=123

TOPIC0=0x...

8) Verify RPC reads + logs (leader + follower) (T2)
Extract variables:

bash
Copier le code
ADDR="$(grep -E '^ADDR=' /tmp/m22_hh_out.txt | tail -n1 | cut -d= -f2-)" && \
BLOCK="$(grep -E '^BLOCK=' /tmp/m22_hh_out.txt | tail -n1 | cut -d= -f2-)" && \
TOPIC0="$(grep -E '^TOPIC0=' /tmp/m22_hh_out.txt | tail -n1 | cut -d= -f2-)" && \
SET_TX="$(grep -E '^SET_TX=' /tmp/m22_hh_out.txt | tail -n1 | cut -d= -f2-)" && \
echo "ADDR=$ADDR" && echo "BLOCK=$BLOCK" && echo "TOPIC0=$TOPIC0" && echo "SET_TX=$SET_TX"
Step 8.1 — Leader: getCode / storage / eth_call / logs
Leader getCode:

bash
Copier le code
curl -s -X POST 127.0.0.1:8545 -H 'content-type: application/json' \
  --data "{\"jsonrpc\":\"2.0\",\"id\":1,\"method\":\"eth_getCode\",\"params\":[\"$ADDR\",\"latest\"]}" ; echo
Leader storage slot 0:

bash
Copier le code
curl -s -X POST 127.0.0.1:8545 -H 'content-type: application/json' \
  --data "{\"jsonrpc\":\"2.0\",\"id\":2,\"method\":\"eth_getStorageAt\",\"params\":[\"$ADDR\",\"0x0\",\"latest\"]}" ; echo
Leader eth_call get() (selector 0x6d4ce63c):

bash
Copier le code
curl -s -X POST 127.0.0.1:8545 -H 'content-type: application/json' \
  --data "{\"jsonrpc\":\"2.0\",\"id\":3,\"method\":\"eth_call\",\"params\":[{\"to\":\"$ADDR\",\"data\":\"0x6d4ce63c\"},\"latest\"]}" ; echo
Leader logs:

bash
Copier le code
curl -s -X POST 127.0.0.1:8545 -H 'content-type: application/json' \
  --data "{\"jsonrpc\":\"2.0\",\"id\":4,\"method\":\"eth_getLogs\",\"params\":[{\"fromBlock\":\"$BLOCK\",\"toBlock\":\"$BLOCK\",\"address\":\"$ADDR\",\"topics\":[\"$TOPIC0\"]}]}" ; echo
Gates (leader):

eth_getCode returns non-empty bytecode (starts with 0x60...)

eth_getStorageAt(slot0) returns ...007b

eth_call get() returns ...007b

eth_getLogs returns exactly one log with:

address == ADDR

blockNumber == BLOCK

topics[0] == TOPIC0

data == ...007b

transactionHash == SET_TX

Step 8.2 — Follower: getCode / storage / eth_call / logs
Follower getCode:

bash
Copier le code
curl -s -X POST 127.0.0.1:8546 -H 'content-type: application/json' \
  --data "{\"jsonrpc\":\"2.0\",\"id\":5,\"method\":\"eth_getCode\",\"params\":[\"$ADDR\",\"latest\"]}" ; echo
Follower storage slot 0:

bash
Copier le code
curl -s -X POST 127.0.0.1:8546 -H 'content-type: application/json' \
  --data "{\"jsonrpc\":\"2.0\",\"id\":6,\"method\":\"eth_getStorageAt\",\"params\":[\"$ADDR\",\"0x0\",\"latest\"]}" ; echo
Follower eth_call get():

bash
Copier le code
curl -s -X POST 127.0.0.1:8546 -H 'content-type: application/json' \
  --data "{\"jsonrpc\":\"2.0\",\"id\":7,\"method\":\"eth_call\",\"params\":[{\"to\":\"$ADDR\",\"data\":\"0x6d4ce63c\"},\"latest\"]}" ; echo
Follower logs:

bash
Copier le code
curl -s -X POST 127.0.0.1:8546 -H 'content-type: application/json' \
  --data "{\"jsonrpc\":\"2.0\",\"id\":8,\"method\":\"eth_getLogs\",\"params\":[{\"fromBlock\":\"$BLOCK\",\"toBlock\":\"$BLOCK\",\"address\":\"$ADDR\",\"topics\":[\"$TOPIC0\"]}]}" ; echo
Gates (follower): outputs match leader exactly (same code, same slot0, same call return, same log entry)

9) Verify tx/receipt + explicit null semantics (leader + follower) (T2)
Define a fake hash:

bash
Copier le code
FAKE=0x0000000000000000000000000000000000000000000000000000000000000001
echo "FAKE=$FAKE"
Step 9.1 — Leader: txByHash + receipt (REAL + FAKE)
Leader txByHash REAL:

bash
Copier le code
curl -s -X POST 127.0.0.1:8545 -H 'content-type: application/json' \
  --data "{\"jsonrpc\":\"2.0\",\"id\":11,\"method\":\"eth_getTransactionByHash\",\"params\":[\"$SET_TX\"]}" ; echo
Leader receipt REAL:

bash
Copier le code
curl -s -X POST 127.0.0.1:8545 -H 'content-type: application/json' \
  --data "{\"jsonrpc\":\"2.0\",\"id\":12,\"method\":\"eth_getTransactionReceipt\",\"params\":[\"$SET_TX\"]}" ; echo
Leader txByHash FAKE:

bash
Copier le code
curl -s -X POST 127.0.0.1:8545 -H 'content-type: application/json' \
  --data "{\"jsonrpc\":\"2.0\",\"id\":13,\"method\":\"eth_getTransactionByHash\",\"params\":[\"$FAKE\"]}" ; echo
Leader receipt FAKE:

bash
Copier le code
curl -s -X POST 127.0.0.1:8545 -H 'content-type: application/json' \
  --data "{\"jsonrpc\":\"2.0\",\"id\":14,\"method\":\"eth_getTransactionReceipt\",\"params\":[\"$FAKE\"]}" ; echo
Gates (leader):

REAL tx result is non-null and has hash == SET_TX, to == ADDR, nonce present

REAL receipt result is non-null and has:

status == 0x1

transactionHash == SET_TX

logs[0].topics[0] == TOPIC0

FAKE calls return result: null

Step 9.2 — Follower: txByHash + receipt (REAL + FAKE)
Follower txByHash REAL:

bash
Copier le code
curl -s -X POST 127.0.0.1:8546 -H 'content-type: application/json' \
  --data "{\"jsonrpc\":\"2.0\",\"id\":21,\"method\":\"eth_getTransactionByHash\",\"params\":[\"$SET_TX\"]}" ; echo
Follower receipt REAL:

bash
Copier le code
curl -s -X POST 127.0.0.1:8546 -H 'content-type: application/json' \
  --data "{\"jsonrpc\":\"2.0\",\"id\":22,\"method\":\"eth_getTransactionReceipt\",\"params\":[\"$SET_TX\"]}" ; echo
Follower txByHash FAKE:

bash
Copier le code
curl -s -X POST 127.0.0.1:8546 -H 'content-type: application/json' \
  --data "{\"jsonrpc\":\"2.0\",\"id\":23,\"method\":\"eth_getTransactionByHash\",\"params\":[\"$FAKE\"]}" ; echo
Follower receipt FAKE:

bash
Copier le code
curl -s -X POST 127.0.0.1:8546 -H 'content-type: application/json' \
  --data "{\"jsonrpc\":\"2.0\",\"id\":24,\"method\":\"eth_getTransactionReceipt\",\"params\":[\"$FAKE\"]}" ; echo
Gates (follower):

REAL responses match leader

FAKE responses return explicit result: null (not missing the result field)

10) Evidence pack (T2)
Create minimal evidence file:

bash
Copier le code
TS="$(date -u +%Y-%m-%dT%H:%M:%SZ)" && \
OUT=/tmp/M22_EVIDENCE_MINIMAL.txt && \
{
  echo "M22_EVIDENCE_MINIMAL ${TS}"
  echo "PK_LEN=${#NOOR_PRIVATE_KEY}"
  echo "NOOR_POSS_REGISTRY=${NOOR_POSS_REGISTRY}"
  echo "--- hardhat out (last) ---"
  tail -n 30 /tmp/m22_hh_out.txt
} > "${OUT}" && echo "WROTE=${OUT}"
Gate: file exists

bash
Copier le code
ls -l /tmp/M22_EVIDENCE_MINIMAL.txt
11) Exit criteria
M22 is PASS if all gates succeeded:

Leader/follower running, ports OK

eth_chainId = 0x849 on both

End-to-end contract flow:

deploy + write OK

eth_getCode non-empty

eth_getStorageAt(slot0) = ...007b

eth_call get() = ...007b

eth_getLogs returns expected log

tx/receipt:

REAL returns objects

FAKE returns explicit result: null (leader + follower)

Evidence file written.

12) Cleanup (optional) (T1)
bash
Copier le code
pkill -INT noorcore || true
bash
Copier le code
pgrep -a noorcore || echo NO_PROCESSES
bash
Copier le code
ss -lntp | grep -E ':(30303|30304|8545|8546|8081|8082)\b' || echo PORTS_FREE