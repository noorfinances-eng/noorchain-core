# NOORCHAIN 2.1 (evm-l1) — RUNBOOK M11 (dApps v0)
Date: 2025-12-30
Scope: Multi-node mainnet-like local stack + Curators Hub v0 dapp smoke-tests.
Terminals:
- T1 = nodes only (no editing, no tooling)
- T2 = tooling (curl, hardhat, node, git)

## 0) Constants
ChainId (EVM): 2121 (0x849)
Node1 (leader): P2P 127.0.0.1:30303 | RPC 127.0.0.1:8545 | Health 127.0.0.1:8081
Node2 (follower): P2P 127.0.0.1:30304 | RPC 127.0.0.1:8546 | Health 127.0.0.1:8082
Health endpoint: /healthz (NOT /)

Repo root: /workspaces/noorchain-core
Dapp root: /workspaces/noorchain-core/dapps/curators-hub-v0

NOTE: It is acceptable to `go build` while nodes run, but the running processes keep the old binary.
To apply changes: STOP → BUILD → START.

## 1) STOP (T2)
Clean stop all noorcore processes, then confirm ports are free.

```bash
cd /workspaces/noorchain-core
pkill -INT -f "(^|/)(noorcore)( |$)" || true
sleep 1
pgrep -a noorcore || true
ss -ltnp | egrep ":30303|:30304|:8545|:8546|:8081|:8082" || true
Expected:

no noorcore process

no LISTEN on those ports

2) BUILD (T2)
bash
Copier le code
cd /workspaces/noorchain-core
go build -o noorcore ./core
./noorcore -h >/dev/null
Expected:

build succeeds

3) START (T1)
Start leader first, then follower. Logs are persisted inside each data-dir (survive /tmp cleanups).

3.1) Node1 (leader)
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
echo "PID=$!"
3.2) Node2 (follower)
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
  > ./data/node2/noorcore.log 2>&1 &
echo "PID=$!"
Hard rule:

never run two nodes with the same -data-dir (LevelDB LOCK).

4) VERIFY stack (T2)
4.1) Processes + ports
bash
Copier le code
cd /workspaces/noorchain-core
echo "=== PROCS ==="
pgrep -a noorcore
echo
echo "=== PORTS ==="
ss -ltnp | egrep ":30303|:30304|:8545|:8546|:8081|:8082"
4.2) Health
bash
Copier le code
curl -fsS http://127.0.0.1:8081/healthz && echo
curl -fsS http://127.0.0.1:8082/healthz && echo
Expected: ok

4.3) RPC smoke (node1 leader)
bash
Copier le code
RPC=http://127.0.0.1:8545
curl -fsS -H "content-type: application/json" \
  -d "{\"jsonrpc\":\"2.0\",\"id\":1,\"method\":\"eth_chainId\",\"params\":[]}" $RPC; echo
curl -fsS -H "content-type: application/json" \
  -d "{\"jsonrpc\":\"2.0\",\"id\":2,\"method\":\"eth_blockNumber\",\"params\":[]}" $RPC; echo
Expected:

chainId = 0x849

blockNumber increases over time (non-zero eventually)

5) dApp — Curators Hub v0 (T2)
5.1) Install deps (first time only)
bash
Copier le code
cd /workspaces/noorchain-core/dapps/curators-hub-v0
npm ci
5.2) Deploy contracts
bash
Copier le code
cd /workspaces/noorchain-core/dapps/curators-hub-v0
npx hardhat run scripts/deploy-curators-hub.mjs --network localhost
Expected:

Connected chainId: 2121

CuratorSet deployed at: 0x...

PoSSRegistry deployed at: 0x... (different from CuratorSet)

5.3) Submit snapshot (PoSSRegistry.submitSnapshot)
Uses a funded dev key via env var PK.

bash
Copier le code
cd /workspaces/noorchain-core/dapps/curators-hub-v0
export PK=0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80
npx hardhat run scripts/send-op-tx.ts --network localhost
Expected:

receipt.status: success

logs print to (PoSSRegistry): <address>

5.4) Verify PoSS state via eth_call (T2)
Replace REG with the printed PoSSRegistry address.

bash
Copier le code
RPC=http://127.0.0.1:8545
REG=0x0000000000000000000000000000000000000000  # <-- set this
echo "snapshotCount:"
curl -fsS -H "content-type: application/json" -d "{\"jsonrpc\":\"2.0\",\"id\":1,\"method\":\"eth_call\",\"params\":[{\"to\":\"$REG\",\"data\":\"0x098ab6a1\"},\"latest\"]}" $RPC; echo
echo "latestSnapshotId:"
curl -fsS -H "content-type: application/json" -d "{\"jsonrpc\":\"2.0\",\"id\":2,\"method\":\"eth_call\",\"params\":[{\"to\":\"$REG\",\"data\":\"0xe484cf32\"},\"latest\"]}" $RPC; echo
Expected:

snapshotCount increments

latestSnapshotId matches

5.5) Decode getSnapshot(latest) (T2)
Replace REG and ID.

bash
Copier le code
cd /workspaces/noorchain-core/dapps/curators-hub-v0
node --input-type=module - <<'"'"'NODE'"'"'
import { decodeFunctionResult, encodeFunctionData } from "viem";
const RPC = "http://127.0.0.1:8545";
const REG = "0x0000000000000000000000000000000000000000"; // set
const id = 1n; // set

const abi = [{
  type: "function", name: "getSnapshot", stateMutability: "view",
  inputs: [{ name: "id", type: "uint256" }],
  outputs: [{
    name: "", type: "tuple", components: [
      { name: "snapshotHash", type: "bytes32" },
      { name: "uri", type: "string" },
      { name: "periodStart", type: "uint64" },
      { name: "periodEnd", type: "uint64" },
      { name: "publishedAt", type: "uint64" },
      { name: "version", type: "uint32" },
      { name: "publisher", type: "address" },
    ]
  }]
}];

const data = encodeFunctionData({ abi, functionName: "getSnapshot", args: [id] });
const body = { jsonrpc: "2.0", id: 1, method: "eth_call", params: [{ to: REG, data }, "latest"] };
const res = await fetch(RPC, { method: "POST", headers: { "content-type": "application/json" }, body: JSON.stringify(body) });
const json = await res.json();

console.log("calldata:", data);
console.log("raw:", json.result);

const decoded = decodeFunctionResult({ abi, functionName: "getSnapshot", data: json.result });
console.log("decoded:", decoded);
NODE
Expected:

raw is not 0x

decoded shows uri/publishedAt/publisher