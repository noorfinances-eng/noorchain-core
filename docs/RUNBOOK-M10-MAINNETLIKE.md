 NOORCHAIN 2.1 — Phase 2 Mainnet-Like Runbook (M10-FINAL)

**Status:** Stable  
**Commit / Tag:** `813b98a` / `M10-MAINNETLIKE-STABLE`  
**Date:** 28 December 2025

## Overview

This runbook describes the deployment, configuration, and operational procedures for a **local mainnet-like NOORCHAIN 2.1 blockchain** with two nodes (leader/follower), fully operational RPC, P2P network, PoSS persistence, and health endpoints. It serves as a reference for development, testing, and dApp integration.

---

## 1. System Requirements

- Go 1.25.5
- LevelDB installed
- Network ports available:
  - P2P: 30303 (Node1), 30304 (Node2)
  - JSON-RPC: 8545 (Node1), 8546 (Node2)
  - Health: 8080 (Node1), 8081 (Node2)

---

## 2. Directory Structure

/workspaces/noorchain-core
├─ core/
├─ data/
│ ├─ node1/
│ └─ node2/
├─ cache/
├─ artifacts/
└─ ...

markdown
Copier le code

- `data/nodeX`: persistent blockchain and PoSS storage
- `cache/` and `artifacts/` are local development files (ignored by git)
- `.gitignore` configured for `data/`, `cache/`, `artifacts/`

---

## 3. Node Configuration

### 3.1 Node1 (Leader)

| Parameter        | Value                        |
|-----------------|------------------------------|
| chain-id        | `noorchain-2-1-local`       |
| data-dir        | `./data/node1`               |
| p2p-addr        | `127.0.0.1:30303`            |
| rpc-addr        | `127.0.0.1:8545`             |
| health-addr     | `127.0.0.1:8080`             |
| boot-peers      | `""`                          |
| role            | `leader`                     |

**Command to start Node1:**

```bash
cd /workspaces/noorchain-core
./noorcore \
  -chain-id noorchain-2-1-local \
  -data-dir ./data/node1 \
  -p2p-addr 127.0.0.1:30303 \
  -rpc-addr 127.0.0.1:8545 \
  -health-addr 127.0.0.1:8080
3.2 Node2 (Follower)
Parameter	Value
chain-id	noorchain-2-1-local
data-dir	./data/node2
p2p-addr	127.0.0.1:30304
rpc-addr	127.0.0.1:8546
health-addr	127.0.0.1:8081
boot-peers	127.0.0.1:30303
role	follower

Command to start Node2:

bash
Copier le code
./noorcore \
  -chain-id noorchain-2-1-local \
  -data-dir ./data/node2 \
  -p2p-addr 127.0.0.1:30304 \
  -rpc-addr 127.0.0.1:8546 \
  -health-addr 127.0.0.1:8081 \
  -boot-peers 127.0.0.1:30303
4. Network Validation
4.1 P2P Connectivity
bash
Copier le code
ss -ltnp | egrep ':30303|:30304'
Node1 and Node2 must be connected (peers=1 for follower)

No socket leak or CLOSE-WAIT accumulation

4.2 JSON-RPC Sanity Checks
Node1:

bash
Copier le code
curl -s -H 'content-type: application/json' \
  --data '{"jsonrpc":"2.0","id":1,"method":"eth_chainId","params":[]}' \
  http://127.0.0.1:8545

curl -s -H 'content-type: application/json' \
  --data '{"jsonrpc":"2.0","id":2,"method":"eth_blockNumber","params":[]}' \
  http://127.0.0.1:8545
Node2:

bash
Copier le code
curl -s -H 'content-type: application/json' \
  --data '{"jsonrpc":"2.0","id":3,"method":"eth_chainId","params":[]}' \
  http://127.0.0.1:8546

curl -s -H 'content-type: application/json' \
  --data '{"jsonrpc":"2.0","id":4,"method":"eth_blockNumber","params":[]}' \
  http://127.0.0.1:8546
eth_chainId must match on both nodes

eth_blockNumber shows Node2 following Node1 (smaller height expected)

4.3 Health Check
bash
Copier le code
curl http://127.0.0.1:8080/healthz
curl http://127.0.0.1:8081/healthz
Both return ok

Confirms multi-node health endpoints are operational and isolated

5. PoSS Snapshot Persistence
Path: data/nodeX/poss/v1/<registryAddr>/

Keys:

count → snapshot count

latest → latest snapshot id

snap/<id_u64_be> → serialized snapshot (RLP)

Persistent across restarts

Receipts stored under rcpt/v1/*

Validation example:

bash
Copier le code
# Node1
eth_call PoSSRegistry snapshotCount
eth_call PoSSRegistry latestSnapshotId
6. Logs / Monitoring
Node logs show tick progression:

yaml
Copier le code
[node] tick | height: N | state: RUNNING
Node1 produces blocks; Node2 follower ticks behind correctly

Check last 50 lines for stability after restart

7. Build / Repo Hygiene
.gitignore excludes:

kotlin
Copier le code
artifacts/
cache/
data/
Go build:

bash
Copier le code
go build -o noorcore ./core
Git stable:

Branch: evm-l1

Tag: M10-MAINNETLIKE-STABLE

Commit: 813b98a

8. Operational Notes
One terminal for Node1 (leader), another for tooling or Node2

Discipline: T1 = runtime / logs; T2 = tooling / curl / git

Ports must remain consistent to avoid collisions

Restart sequence:

Stop follower

Start leader if needed

Start follower with correct -boot-peers and -health-addr

Use JSON-RPC or health endpoints to verify proper sync

9. Next Steps (M11)
Deploy dApps v0: Curators Hub + modules essential

Connect wallets to JSON-RPC endpoints

Integrate PoSS interactions from dApps

Verify full transaction path and receipt persistence

End of Runbook — M10-FINAL Mainnet-Like Stack