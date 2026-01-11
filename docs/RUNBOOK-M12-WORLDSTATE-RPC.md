NOORCHAIN 2.1 — M12 Runbook (World State + RPC Parity Foundations)
0) Scope and Objective

M12 objective (non-ambiguous): implement an Ethereum-compatible persistent world state (accounts/balances/nonces/code/storage) backed by go-ethereum StateDB + trie/triedb/rawdb, persist a stateRoot per block, and evolve JSON-RPC so “reads” become real state reads (not shims). Ensure correctness under leader/follower (follower routes leader-only reads).

M12 closure status (as executed):

M12.2: Geth state store integrated + commit/roots, RPC reads for nonce/balance.

M12.3: Follower routing for world-state reads (leader-only).

M12.5–M12.6: Dev alloc bootstrap + tooling hardening for local tests (Hardhat/v3).

M12.7: PoSS snapshot ABI fidelity fixed for getSnapshot (calldata decode + publisher) + full RLP meta persisted; Hardhat script supports NOOR_* overrides.

Tags used across the path (project history): M12.2-WORLDSTATE-RPC-NONCE-BALANCE, M12.6-VALIDATED, M12.7-VALIDATED (and earlier M12 partial tags as applicable).

1) Invariants (Do Not Break)

Consensus/security remains separate from PoSS (PoSS is application layer).

Leader/follower topology stays stable: follower proxies leader-only reads using FollowRPC as the authoritative signal.

Operational discipline:

T1 = nodes only (leader/follower)

T2 = tooling only (curl, node, hardhat, git)

“1 command → 1 edit → 1 gate”

Never run two nodes on the same -data-dir (LevelDB lock will fail).

Do not commit or store private keys.

2) Preconditions
2.1 Repo and Branch
cd /workspaces/noorchain-core
git checkout evm-l1
git pull

2.2 Build toolchain

Go: per repo (go env should match project standard)

Node.js + npm for scripts/Hardhat

Ensure local artifacts exist when needed (artifacts/contracts/...)

2.3 Network parameters (canonical local)

chainId = 2121 (0x849)

Leader:

P2P: 127.0.0.1:30303

RPC: 127.0.0.1:8545

Health: 127.0.0.1:8081

Follower:

P2P: 127.0.0.1:30304

RPC: 127.0.0.1:8546

Health: 127.0.0.1:8082

-follow-rpc http://127.0.0.1:8545

3) Standard Operations
3.1 Stop nodes (T1)
pkill -INT noorcore || true


Gate (T2): verify ports are closed

ss -lntp | grep -E ':(30303|30304|8545|8546|8081|8082)\b' || true

3.2 Build (T2)

Always build the binary explicitly:

cd /workspaces/noorchain-core && go build -o noorcore ./core


Gate: binary exists

ls -la /workspaces/noorchain-core/noorcore

3.3 Start leader (T1)
cd /workspaces/noorchain-core && \
./noorcore \
  -chain-id noorchain-2-1-local \
  -data-dir ./data/node1 \
  -p2p-addr 127.0.0.1:30303 \
  -rpc-addr 127.0.0.1:8545 \
  -health-addr 127.0.0.1:8081 \
  -alloc-file ./alloc-dev.json \
  > /tmp/noorcore_node1.log 2>&1 &


Gate (T2):

pgrep -a noorcore | head -n 5
ss -lntp | grep -E ':(30303|8545|8081)\b'
curl -s -H 'content-type: application/json' \
  -d '{"jsonrpc":"2.0","id":1,"method":"eth_chainId","params":[]}' \
  http://127.0.0.1:8545


Expected: "result":"0x849".

3.4 Start follower (T1)
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


Gate (T2):

ss -lntp | grep -E ':(30304|8546|8082)\b'
curl -s -H 'content-type: application/json' \
  -d '{"jsonrpc":"2.0","id":1,"method":"eth_chainId","params":[]}' \
  http://127.0.0.1:8546

4) M12.0–M12.1 — Storage Foundation (Geth DB isolation)
Intent

Introduce a dedicated geth DB under <data-dir>/db/geth and wire node lifecycle open/close.

Validation gates

Node starts cleanly.

No DB lock issues (one process per data-dir).

Logs indicate DB opened and closed cleanly on shutdown.

5) M12.1–M12.2 — Block Metadata + Roots Exposure (blkmeta)
Intent

Persist per-block metadata, including:

stateRoot (initially placeholder → later real)

receiptsRoot

logsBloom

and expose them via eth_getBlockByNumber.

Gate checks (T2)

Latest block fields:

curl -s -H 'content-type: application/json' \
  -d '{"jsonrpc":"2.0","id":1,"method":"eth_getBlockByNumber","params":["latest",false]}' \
  http://127.0.0.1:8545 | head -c 500; echo


Expect: non-null block, fields include stateRoot, receiptsRoot, logsBloom.

Query above head returns null:

curl -s -H 'content-type: application/json' \
  -d '{"jsonrpc":"2.0","id":1,"method":"eth_getBlockByNumber","params":["0xFFFFFFFF",false]}' \
  http://127.0.0.1:8545


Expect: "result":null.

Follower proxy works:

curl -s -H 'content-type: application/json' \
  -d '{"jsonrpc":"2.0","id":1,"method":"eth_getBlockByNumber","params":["latest",false]}' \
  http://127.0.0.1:8546 | head -c 500; echo

6) M12.2 — World-State Commit (StateDB + triedb) + Real Reads (nonce/balance)
Intent

Commit real geth StateDB per block.

Persist “head stateRoot” (e.g., under a NOOR KV like stateroot/v1/head).

Make eth_getTransactionCount and eth_getBalance read from StateDB.

Minimal “proof writes” were used to validate persistence (later replaced in M13 by full EVM execution).

Gate checks (T2)

Nonce reads from state:

curl -s -H 'content-type: application/json' \
  -d '{"jsonrpc":"2.0","id":1,"method":"eth_getTransactionCount","params":["0x4aA5DA75AFb6e81F433D4720cb7Cb2C6B1BA323c","latest"]}' \
  http://127.0.0.1:8545


Balance reads from state:

curl -s -H 'content-type: application/json' \
  -d '{"jsonrpc":"2.0","id":1,"method":"eth_getBalance","params":["0x4aA5DA75AFb6e81F433D4720cb7Cb2C6B1BA323c","latest"]}' \
  http://127.0.0.1:8545


Persistence across restart:

Stop nodes, restart leader only, re-run the two queries.
Expected: values persist.

7) M12.3 — Follower Routing for World-State Reads (Leader-only)
Intent

On follower, route state-critical reads via leader whenever FollowRPC != "":

eth_getTransactionCount

eth_getBalance
(and later extended similarly for eth_getCode, eth_getStorageAt)

Gate checks (T2)

Compare leader vs follower results:

curl -s -H 'content-type: application/json' \
  -d '{"jsonrpc":"2.0","id":1,"method":"eth_getTransactionCount","params":["0x4aA5DA75AFb6e81F433D4720cb7Cb2C6B1BA323c","latest"]}' \
  http://127.0.0.1:8545
curl -s -H 'content-type: application/json' \
  -d '{"jsonrpc":"2.0","id":2,"method":"eth_getTransactionCount","params":["0x4aA5DA75AFb6e81F433D4720cb7Cb2C6B1BA323c","latest"]}' \
  http://127.0.0.1:8546


Expected: identical outputs.

8) M12.5–M12.6 — Dev Alloc Bootstrap + Tooling Hardening
Intent

Ensure local developer address can transact:

alloc-dev.json provides initial funded balance for a dev address (e.g. 0x4aA5...).

Alloc applies at boot and writes a guard key (e.g., alloc/v1/applied) to prevent re-applying.

Apply alloc (T1 leader start)

Start leader with:

-alloc-file ./alloc-dev.json

Gate (T2)

Confirm log evidence:

grep -n "alloc:" /tmp/noorcore_node1.log | tail -n 20


Expect lines similar to “applying …” then “applied | new head root …”.

Confirm funded balance:

curl -s -H 'content-type: application/json' \
  -d '{"jsonrpc":"2.0","id":1,"method":"eth_getBalance","params":["0x4aA5DA75AFb6e81F433D4720cb7Cb2C6B1BA323c","latest"]}' \
  http://127.0.0.1:8545

Re-apply alloc (only if needed)

Stop node

Remove alloc guard key (method depends on your NOOR DB tooling; simplest is delete the specific key if you have a helper; otherwise reset the data-dir for local dev)

Restart leader with -alloc-file

Tooling pitfalls (Hardhat v3)

ethers is not global; use network.connect() patterns where required.

If deploy scripts fail with getSigners() empty, your private key env is missing.

9) M12.7 — PoSS ABI Fidelity for getSnapshot + Full poss/v1 meta persistence
Intent

Fix getSnapshot fidelity vs Hardhat expectations:

Track core/exec/exec.go (not ignored)

Decode submitSnapshot calldata properly

Recover publisher/signer correctly

Persist full snapshot meta in poss/v1 RLP

Validate getSnapshot(2) matches Hardhat fields on leader and follower.

Tooling update: submit script honors NOOR_* env overrides

scripts/submit-snapshot.mjs supports:

NOOR_RPC_URL

NOOR_PRIVATE_KEY (or NOOR_PRIVKEY)

NOOR_POSS_REGISTRY (or NOOR_REGISTRY)

NOOR_GAS_LIMIT, NOOR_GAS_PRICE

NOOR_SNAPSHOT_URI

Gate run (T2)
cd /workspaces/noorchain-core && \
read -s -p "NOOR_PRIVATE_KEY (0x...): " NOOR_PRIVATE_KEY; echo; export NOOR_PRIVATE_KEY; \
NOOR_RPC_URL="http://127.0.0.1:8545" \
NOOR_POSS_REGISTRY="0xC9F398646E19778F2C3D9fF32bb75E5a99FD4E56" \
node scripts/submit-snapshot.mjs


Expect: receipt status: 1 and SUBMIT_OK.

ABI fidelity gates (T2)

Leader snapshot read:

curl -s -H 'content-type: application/json' \
  -d '{"jsonrpc":"2.0","id":1,"method":"eth_call","params":[{"to":"0xC9F398646E19778F2C3D9fF32bb75E5a99FD4E56","data":"<getSnapshot(0x2) calldata>"}, "latest"]}' \
  http://127.0.0.1:8545


Follower same call:

curl -s -H 'content-type: application/json' \
  -d '{"jsonrpc":"2.0","id":1,"method":"eth_call","params":[{"to":"0xC9F398646E19778F2C3D9fF32bb75E5a99FD4E56","data":"<getSnapshot(0x2) calldata>"}, "latest"]}' \
  http://127.0.0.1:8546


Expected: returned ABI-encoded fields match what Hardhat decodes (id/hash/uri/periodStart/periodEnd/version/publisher/sigs, depending on your contract ABI).

Note: in practice you likely validate via a Hardhat/ethers read (registry.getSnapshot(2)) against both RPC endpoints and compare decoded objects.

10) Git Hygiene and Release Discipline
Commit isolation

Do not mix unrelated changes. Stage only what you intend.

Example pattern:

git add path/to/file
git diff --cached
git commit -m "..."
git push

When WIP exists

Use stash before switching milestones:

git stash push -u -m "wip: ..."

11) Troubleshooting
RPC down / ECONNREFUSED

Gate:

ss -lntp | grep -E ':(8545|8546)\b' || true
curl -s -H 'content-type: application/json' \
  -d '{"jsonrpc":"2.0","id":1,"method":"eth_chainId","params":[]}' \
  http://127.0.0.1:8545

LevelDB lock

Cause: two processes using same -data-dir. Fix: stop all nodes, verify noorcore PIDs, restart with separate dirs.

Alloc not applied / balance is 0

Verify node started with -alloc-file

Check /tmp/noorcore_node1.log for alloc lines

Remember alloc may be guarded by an “applied” key; re-apply requires clearing that guard or using a fresh data-dir.

Hardhat deploy errors (NOOR_PRIVATE_KEY missing)

Symptom: getSigners() empty / deployer undefined.
Fix: set env at runtime, do not hardcode keys.

12) Exit Criteria (M12 acceptance checklist)

M12 is considered complete when:

Leader produces blocks with a persisted stateRoot derived from geth StateDB commit.

eth_getBalance and eth_getTransactionCount return StateDB-backed values and persist across restart.

Follower routes leader-only state reads using FollowRPC and matches leader responses.

PoSS getSnapshot ABI fidelity matches Hardhat decode (leader and follower).

Tooling can submit a PoSS snapshot with env overrides and returns SUBMIT_OK.

13) Next Milestone: M13 (What changes)

M13 replaces any “proof writes” with real EVM execution:

Proper CALL/CREATE execution path

Correct gas semantics

Correct logs and receipts

eth_call executes against state without mining

Extend RPC parity (code/storage/logs where required)