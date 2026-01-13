# NOORCHAIN 2.1 — M9 Runbook (Contracts Execution Minimal)

## 0) Scope

M9 validates **minimal contract execution at mining** (not full EVM parity):
- `eth_sendRawTransaction` → mined into a block
- receipts persisted under `rcpt/v1/*`
- PoSS `submitSnapshot` is executed during mining via the node applyTx hook
- PoSS snapshots persisted under `poss/v1/*` (RLP)
- `eth_getTransactionReceipt` returns persisted receipts (LevelDB-first)

**Important:** M9 is a “mainnet-like minimal execution hook” milestone. Full world-state/EVM parity is handled in later milestones (M12/M13+).

## 1) Invariants / Safety

- T1 = nodes only; T2 = tooling only.
- One process per `-data-dir` (LevelDB locks otherwise).
- One command → one gate.
- Do NOT use or commit private keys.
- Keep PoSS as app-layer logic (no consensus coupling).

## 2) Preconditions

- Repo: `noorchain-core`
- Branch/tag: `evm-l1` at `M9-CONTRACTS-EXECUTION-STABLE` (or known M9 commit)
- Build uses the explicit command: `go build -o noorcore ./core`
  (Do not rely on `go build ./...` for producing the `noorcore` binary.)

## 3) Standard Ports (local)

Leader node (single-node M9 validation):
- P2P: `127.0.0.1:30303`
- RPC: `127.0.0.1:8545`

## 4) Stop / Clean (T1 → T2 gate)

### T1 — Stop node
```bash
pkill -INT noorcore || true
T2 — Gate: ports closed
bash
Copier le code
ss -lntp | grep -E ':(30303|8545)\b' || true
5) Build (T2)
bash
Copier le code
cd /workspaces/noorchain-core && go build -o noorcore ./core
Gate:

bash
Copier le code
cd /workspaces/noorchain-core && ./noorcore -version 2>/dev/null || true
ls -la /workspaces/noorchain-core/noorcore
6) Start node (T1) + RPC gates (T2)
T1 — Start leader (single node)
bash
Copier le code
cd /workspaces/noorchain-core && \
./noorcore \
  -chain-id noorchain-2-1-local \
  -data-dir ./data/node1 \
  -p2p-addr 127.0.0.1:30303 \
  -rpc-addr 127.0.0.1:8545 \
  > /tmp/noorcore_node1.log 2>&1 &
T2 — Gates (ports + chainId + blockNumber)
bash
Copier le code
ss -lntp | grep -E ':(30303|8545)\b'

curl -s -H 'content-type: application/json' \
  -d '{"jsonrpc":"2.0","id":1,"method":"eth_chainId","params":[]}' \
  http://127.0.0.1:8545

curl -s -H 'content-type: application/json' \
  -d '{"jsonrpc":"2.0","id":2,"method":"eth_blockNumber","params":[]}' \
  http://127.0.0.1:8545
Expected:

eth_chainId returns the configured chain id (project standard for local).

eth_blockNumber is available (may start at 0x0 and then increment once mining loop runs).

7) Submit PoSS snapshot transaction (tooling)
M9 validates that a PoSS submitSnapshot-like transaction:

is accepted by RPC (eth_sendRawTransaction)

is mined

triggers minimal execution hook at mining

persists receipt and PoSS snapshot RLP

Option A — Use existing repo scripts (preferred)
Use the repo’s canonical script for PoSS submission at M9 time (examples: scripts/submit-snapshot.mjs).
Run from repo root.

Gate: the command outputs a tx hash and SUBMIT_OK.

If your script needs env variables (RPC, private key, registry address), set them in the environment at runtime.
Never store keys in files.

Option B — Raw RPC path (wallet-compatible)
If you have a signed raw tx:

call eth_sendRawTransaction

then poll eth_getTransactionReceipt until status is non-null

Gate: receipt shows status=0x1 (or 1 depending encoding) and has a block number/hash.

8) Validate receipts are persisted (rcpt/v1)
M9 requires:

receipts are written into LevelDB under a stable key prefix rcpt/v1/*

JSON-RPC eth_getTransactionReceipt reads LevelDB first

T2 — Gate: receipt query works
bash
Copier le code
curl -s -H 'content-type: application/json' \
  -d '{"jsonrpc":"2.0","id":3,"method":"eth_getTransactionReceipt","params":["0x<YOUR_TX_HASH>"]}' \
  http://127.0.0.1:8545
Expected: non-null receipt; transactionHash matches; mined fields present.

9) Validate PoSS snapshot persistence (poss/v1)
M9 requires:

PoSS snapshot meta persisted under poss/v1/* (RLP)

Validation methods depend on the available RPC methods at that milestone:

If PoSS read methods exist (snapshotCount, latestSnapshotId, getSnapshot) verify via eth_call.

If not, verify via logs and/or DB inspection helpers available in repo at M9.

T2 — Gate (example if snapshotCount exists)
bash
Copier le code
curl -s -H 'content-type: application/json' \
  -d '{"jsonrpc":"2.0","id":4,"method":"eth_call","params":[{"to":"0x<PoSSRegistry>","data":"0x098ab6a1"},"latest"]}' \
  http://127.0.0.1:8545
Expected: non-zero count after submissions (depending on how many you sent).

10) Persistence Across Restart
M9 requires data survives restart:

stop node, restart with same -data-dir

confirm eth_getTransactionReceipt still returns the same receipt

confirm PoSS snapshot reads (if available) still return expected values

T1 stop
bash
Copier le code
pkill -INT noorcore || true
T1 start (same command as above)
(restart leader)

T2 gates
Re-run:

eth_getTransactionReceipt for the previous tx hash

PoSS snapshot reads (if available)

11) Troubleshooting
ECONNREFUSED / RPC down
Verify process and port:

bash
Copier le code
pgrep -a noorcore | head -n 5
ss -lntp | grep -E ':(8545)\b'
LevelDB LOCK
Two processes share the same -data-dir

Stop all nodes; confirm only one noorcore instance.

Tooling pitfalls (Hardhat v3)
ethers may not be global; use network.connect() + conn.ethers.

Tasks may not load if config format mismatches (ESM config recommended).

Avoid typing shell commands inside Hardhat REPL; exit with .exit.

12) Exit Criteria (M9 pass/fail)
M9 is PASS when all are true:

eth_sendRawTransaction path works and tx is mined.

eth_getTransactionReceipt returns a non-null receipt for mined tx.

Receipt is persisted in LevelDB and survives restart.

PoSS submit snapshot execution hook runs at mining and persists poss/v1/* snapshot meta (with a verifiable read path or equivalent validation evidence).