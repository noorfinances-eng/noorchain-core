# NOORCHAIN 2.1 (evm-l1) — RUNBOOK M11 dApps v0
Date: 2025-12-30
Branch: evm-l1

## Terminal discipline
- T1 = nodes (node1/node2 only)
- T2 = tooling (git, build, curl, dApps)
- T3 = optional (tail logs, extra checks)
Rule: one step / one command at a time.

## Ports (default)
- Node1: P2P 30303, RPC 8545, HEALTH 8081
- Node2: P2P 30304, RPC 8546, HEALTH 8082
Health path is **/healthz** (NOT /).

---

## A) STOP (clean)
(T2)

```bash
pkill -INT -f '(^|/)(noorcore)( |$)' || true
sleep 1
pgrep -a noorcore || true
ss -ltnp | egrep ':30303|:30304|:8545|:8546|:8081|:8082' || true
Gate expected:

no noorcore process

no LISTEN on 30303/30304/8545/8546/8081/8082

B) BUILD (clean)
(T2)

bash
Copier le code
cd /workspaces/noorchain-core
go build -o noorcore ./core
Gate expected:

BUILD_OK (no errors)

C) START (2 nodes mainnet-like)
(T1)

Start node1 (leader)
bash
Copier le code
cd /workspaces/noorchain-core
./noorcore \
  -chain-id noorchain-2-1-local \
  -data-dir ./data/node1 \
  -p2p-addr 127.0.0.1:30303 \
  -rpc-addr 127.0.0.1:8545 \
  -health-addr 127.0.0.1:8081 \
  -boot-peers 127.0.0.1:30304 \
  > ./data/node1/noorcore.log 2>&1 &
echo "PID=$(pgrep -n noorcore)"
Start node2 (follower)
bash
Copier le code
cd /workspaces/noorchain-core
./noorcore \
  -chain-id noorchain-2-1-local \
  -data-dir ./data/node2 \
  -p2p-addr 127.0.0.1:30304 \
  -rpc-addr 127.0.0.1:8546 \
  -health-addr 127.0.0.1:8082 \
  -follow-rpc http://127.0.0.1:8545 \
  > /tmp/noorcore_node2.log 2>&1 &
echo "PID=$(pgrep -n noorcore)"
D) VERIFY (nodes)
(T2)

Processes + ports
bash
Copier le code
pgrep -a noorcore
ss -ltnp | egrep ':30303|:30304|:8545|:8546|:8081|:8082'
Health endpoints
bash
Copier le code
curl -fsS http://127.0.0.1:8081/healthz; echo
curl -fsS http://127.0.0.1:8082/healthz; echo
RPC (node1)
bash
Copier le code
curl -s http://127.0.0.1:8545 -H 'content-type: application/json' \
  --data '{"jsonrpc":"2.0","id":1,"method":"eth_chainId","params":[]}'
echo
curl -s http://127.0.0.1:8545 -H 'content-type: application/json' \
  --data '{"jsonrpc":"2.0","id":2,"method":"eth_blockNumber","params":[]}'
echo
curl -s http://127.0.0.1:8545 -H 'content-type: application/json' \
  --data '{"jsonrpc":"2.0","id":3,"method":"eth_getBalance","params":["0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266","latest"]}'
echo
curl -s http://127.0.0.1:8545 -H 'content-type: application/json' \
  --data '{"jsonrpc":"2.0","id":4,"method":"eth_getTransactionCount","params":["0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266","latest"]}'
echo
Gate expected:

eth_chainId = 0x849 (2121)

healthz returns "ok"

RPC answers without errors

E) dApp (Curators Hub v0)
Path: /workspaces/noorchain-core/dapps/curators-hub-v0

Install deps (only if needed)
(T2)

bash
Copier le code
cd /workspaces/noorchain-core/dapps/curators-hub-v0
npm ci
1) Deploy contracts (writes deployments json)
(T2)

bash
Copier le code
cd /workspaces/noorchain-core/dapps/curators-hub-v0
node ./scripts/deploy-curators-hub.mjs
Expected:

prints chainId 2121

prints CuratorSet address + PoSSRegistry address

creates/updates: deployments/noorchain-2.1-local.json

2) Submit PoSS snapshot (requires PK env)
(T2)

bash
Copier le code
cd /workspaces/noorchain-core/dapps/curators-hub-v0
export PK=0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80
npx hardhat run ./scripts/send-op-tx.ts --network localhost
Expected:

receipt.status: success

snapshotCount increments

3) Read PoSS state (snapshotCount/latest/getSnapshot)
(T2)

bash
Copier le code
cd /workspaces/noorchain-core/dapps/curators-hub-v0
node ./scripts/read-poss.mjs
Expected:

snapshotCount >= 1

latestSnapshotId matches

getSnapshot(latest) shows uri, timestamps, publisher

F) Troubleshooting quick hits
Health 404
Use /healthz, not /

Hardhat network errors
Use:

--network localhost
and ensure hardhat.config.ts defines localhost with chainId 2121.

Nonce / duplicate contract addresses
Verify:

bash
Copier le code
curl -s http://127.0.0.1:8545 -H 'content-type: application/json' \
  --data '{"jsonrpc":"2.0","id":1,"method":"eth_getTransactionCount","params":["0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266","latest"]}'
echo
If something feels stale
STOP all nodes, verify ports closed, rebuild, start again.