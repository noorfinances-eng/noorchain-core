# NOORCHAIN 2.1 (evm-l1) — RUNBOOK M11 dApps v0

Scope : opérations propres (stop/build/start/verify) + déploiement/validation dApp Curators Hub v0 via RPC leader (node1).
Invariants : chainId EVM = 2121 (0x849). Multi-nœuds : node1 leader, node2 follower.

## Discipline
- T1 = nodes uniquement
- T2 = tooling uniquement
- Une commande à la fois, valider les gates.

## Ports standard
- node1 (leader) : P2P 127.0.0.1:30303 | RPC 127.0.0.1:8545 | Health 127.0.0.1:8081
- node2 (follower) : P2P 127.0.0.1:30304 | RPC 127.0.0.1:8546 | Health 127.0.0.1:8082 | Follow RPC http://127.0.0.1:8545

## Logs
- node1 : /tmp/noorcore_node1.log
- node2 : /tmp/noorcore_node2.log

---

# 1) STOP (clean)
Terminal : T2
Gate : STOP_OK puis PORTS_OK

```bash
pkill -INT -f "(^|/)(noorcore)( |$)" || true
sleep 1
pgrep -a noorcore || echo "STOP_OK"
ss -ltnp | egrep ":(30303|30304|8545|8546|8081|8082)\b" || echo "PORTS_OK"
```

---

# 2) BUILD (clean)
Terminal : T2
Gate : BUILD_OK

```bash
cd /workspaces/noorchain-core
go build -o noorcore ./core
echo BUILD_OK
```

---

# 3) START node1 + node2 (leader/follower)
Terminal : T2
Gate : 2 PIDs noorcore + ports LISTEN OK

```bash
cd /workspaces/noorchain-core

./noorcore -chain-id noorchain-2-1-local -data-dir ./data/node1 \
  -p2p-addr 127.0.0.1:30303 -rpc-addr 127.0.0.1:8545 -health-addr 127.0.0.1:8081 \
  -boot-peers 127.0.0.1:30304 > /tmp/noorcore_node1.log 2>&1 &

sleep 0.8

./noorcore -chain-id noorchain-2-1-local -data-dir ./data/node2 \
  -p2p-addr 127.0.0.1:30304 -rpc-addr 127.0.0.1:8546 -health-addr 127.0.0.1:8082 \
  -follow-rpc http://127.0.0.1:8545 > /tmp/noorcore_node2.log 2>&1 &

sleep 1

pgrep -a noorcore || true
ss -ltnp | egrep ":(30303|30304|8545|8546|8081|8082)\b" || true
```

---

# 4) VERIFY — RPC sanity (leader RPC 8545)
Terminal : T2
Gates :
- eth_chainId == 0x849
- eth_getBalance répond (0x0 minimal)
- eth_getTransactionCount répond (nonce hex)

```bash
curl -s -H "content-type: application/json" \
  --data "{\"jsonrpc\":\"2.0\",\"id\":1,\"method\":\"eth_chainId\",\"params\":[]}" \
  http://127.0.0.1:8545

curl -s -H "content-type: application/json" \
  --data "{\"jsonrpc\":\"2.0\",\"id\":2,\"method\":\"eth_getBalance\",\"params\":[\"0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266\",\"latest\"]}" \
  http://127.0.0.1:8545

curl -s -H "content-type: application/json" \
  --data "{\"jsonrpc\":\"2.0\",\"id\":3,\"method\":\"eth_getTransactionCount\",\"params\":[\"0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266\",\"latest\"]}" \
  http://127.0.0.1:8545
```

---

# 5) dApp — Curators Hub v0 (deploy + receipts)

## 5.1 Env (PK)
Terminal : T2

```bash
export PK=0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80
```

## 5.2 Deploy CuratorSet + PoSSRegistry
Terminal : T2
Gate :
- CuratorSet tx != PoSSRegistry tx
- CuratorSet deployed at != PoSSRegistry deployed at

```bash
cd /workspaces/noorchain-core/dapps/curators-hub-v0
export PK=0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80
node ./scripts/deploy-curators-hub.mjs
```

## 5.3 Receipts check (coller TX1/TX2)
Terminal : T2
Gate : contractAddress non-null sur les 2

```bash
TX1=0x...
TX2=0x...

for h in $TX1 $TX2; do
  echo "=== $h ==="
  curl -s -H "content-type: application/json" \
    --data "{\"jsonrpc\":\"2.0\",\"id\":1,\"method\":\"eth_getTransactionReceipt\",\"params\":[\"$h\"]}" \
    http://127.0.0.1:8545
  echo
done
```

