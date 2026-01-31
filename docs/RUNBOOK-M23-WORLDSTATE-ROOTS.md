# RUNBOOK — M23 WORLD-STATE ROOTS (client-grade) — roots/bloom/txroot persisted + follower parity

## Scope

This runbook validates **client-grade block metadata** on NOORCHAIN 2.1:
- `stateRoot` is non-zero and stable per block
- `receiptsRoot` is derived and persisted per block
- `logsBloom` is derived and persisted per block
- `transactionsRoot` is derived from **raw signed tx RLP** (txpool) and persisted per block
- JSON-RPC `eth_getBlockByNumber` returns stable values sourced from persisted metadata (not `time.Now()`).
- **Follower parity**: follower returns the same block fields as leader for the same height.

This runbook also includes mandatory **Private Key Gate** and **PoSS Registry Gate**.

## Invariants

- Conversations: irrelevant. This file is an operational artifact.
- Chain: `chain-id = noorchain-2-1-local`
- EVM Chain ID (tooling): `2121` (hex: `0x849`)
- Two nodes: leader + follower.
- FollowRPC is authoritative on follower (proxy semantics).
- No private key is ever committed or printed.

## Terminals discipline

- **T1 = nodes only** (start/stop, process/ports checks)
- **T2 = tooling only** (curl, scripts, git)
- One command → one gate → next step.

## Ports (local)

- Leader:  P2P `127.0.0.1:30303`, RPC `127.0.0.1:8545`, Health `127.0.0.1:8081`
- Follower: P2P `127.0.0.1:30304`, RPC `127.0.0.1:8546`, Health `127.0.0.1:8082`
- Follower follow-rpc: `http://127.0.0.1:8545`

## Reference constants

Empty-trie root (expected `transactionsRoot` on empty blocks):
`0x56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421`

PoSS Registry (mandatory env override):
`0xC9F398646E19778F2C3D9fF32bb75E5a99FD4E56`

---

## Step 0 — Repo gate (T2)

```bash
cd /workspaces/noorchain-core && git status --porcelain
Gate: expected output is empty (or only known non-tracked dirs you intentionally ignore). If not clean, stop and resolve before proceeding.

Step 1 — Stop nodes + ports free (T1)
bash
Copier le code
pkill -INT noorcore || true
bash
Copier le code
ss -lntp | grep -E ':(30303|30304|8545|8546|8081|8082)\b' || echo PORTS_FREE
Gate: PORTS_FREE

Step 2 — Build gate (T2)
bash
Copier le code
cd /workspaces/noorchain-core && go build -o noorcore ./core && echo BUILD_OK
Gate: BUILD_OK

Step 3 — Start leader + follower (T1)
Leader:

bash
Copier le code
cd /workspaces/noorchain-core && \
./noorcore -chain-id noorchain-2-1-local -data-dir ./data/node1 -role leader \
  -p2p-addr 127.0.0.1:30303 -rpc-addr 127.0.0.1:8545 -health-addr 127.0.0.1:8081 \
  -boot-peers 127.0.0.1:30304 > /tmp/noorcore_node1.log 2>&1 & echo "LEADER_PID=$!"
Follower:

bash
Copier le code
cd /workspaces/noorchain-core && \
./noorcore -chain-id noorchain-2-1-local -data-dir ./data/node2 -role follower \
  -p2p-addr 127.0.0.1:30304 -rpc-addr 127.0.0.1:8546 -health-addr 127.0.0.1:8082 \
  -follow-rpc http://127.0.0.1:8545 -boot-peers 127.0.0.1:30303 > /tmp/noorcore_node2.log 2>&1 & echo "FOLLOWER_PID=$!"
Process/ports gate:

bash
Copier le code
pgrep -a noorcore && ss -lntp | grep -E ':(8545|8546)\b'
Gate: both RPC listeners are present (:8545 and :8546).

Step 4 — Basic RPC gates (T2)
Leader:

bash
Copier le code
curl -s -X POST http://127.0.0.1:8545 -H 'content-type: application/json' \
  --data '[{"jsonrpc":"2.0","id":1,"method":"eth_chainId","params":[]},{"jsonrpc":"2.0","id":2,"method":"eth_blockNumber","params":[]}]'
Follower:

bash
Copier le code
curl -s -X POST http://127.0.0.1:8546 -H 'content-type: application/json' \
  --data '[{"jsonrpc":"2.0","id":3,"method":"eth_chainId","params":[]},{"jsonrpc":"2.0","id":4,"method":"eth_blockNumber","params":[]}]'
Gate:

chainId is 0x849 on both

blockNumber is non-decreasing over time

Step 5 — Block metadata fields are served from persisted blkmeta (T2)
Request a concrete height H (use latest), and confirm timestamp is stable and roots/bloom exist.

bash
Copier le code
curl -s -X POST http://127.0.0.1:8545 -H 'content-type: application/json' \
  --data '{"jsonrpc":"2.0","id":10,"method":"eth_getBlockByNumber","params":["latest",false]}'
Gate:

timestamp is present and reasonable (not 0x0)

stateRoot, receiptsRoot, transactionsRoot, logsBloom are present

On an empty block, transactionsRoot may equal the empty-trie root.

Optional stability check (same height twice should be identical; avoid latest because it advances):

capture height:

bash
Copier le code
curl -s -X POST http://127.0.0.1:8545 -H 'content-type: application/json' \
  --data '{"jsonrpc":"2.0","id":11,"method":"eth_blockNumber","params":[]}'
query by explicit height (replace 0x...):

bash
Copier le code
curl -s -X POST http://127.0.0.1:8545 -H 'content-type: application/json' \
  --data '{"jsonrpc":"2.0","id":12,"method":"eth_getBlockByNumber","params":["0xHEIGHT",false]}'
Re-run the same command again.

Gate: timestamp, transactionsRoot, receiptsRoot, logsBloom, stateRoot are identical across repeated reads at the same height.

Step 6 — Mandatory Private Key Gate + PoSS Registry Gate (T2)
This step must pass before running any deploy/submit script.

Set env vars (runtime only; never commit):

bash
Copier le code
while true; do
  read -s -p "Paste NOOR_PRIVATE_KEY (0x... len=66): " NOOR_PRIVATE_KEY; echo
  export NOOR_PRIVATE_KEY
  python - <<'PY'
import os, sys
pk=os.environ.get("NOOR_PRIVATE_KEY","")
print("NOOR_PRIVATE_KEY_LEN=", len(pk))
print("NOOR_PRIVATE_KEY_0x=", pk.startswith("0x"))
sys.exit(0 if (len(pk)==66 and pk.startswith("0x")) else 1)
PY
  if [ $? -eq 0 ]; then break; fi
  echo "INVALID_KEY (need 0x-prefixed len=66). Try again."
done
export NOOR_POSS_REGISTRY=0xC9F398646E19778F2C3D9fF32bb75E5a99FD4E56
echo "NOOR_POSS_REGISTRY=$NOOR_POSS_REGISTRY"
Gate:

NOOR_PRIVATE_KEY_LEN= 66

NOOR_PRIVATE_KEY_0x= True

NOOR_POSS_REGISTRY prints the expected address (len 42)

Step 7 — Mine 1 transaction (PoSS submitSnapshot) (T2)
bash
Copier le code
cd /workspaces/noorchain-core && node scripts/submit-snapshot.mjs
Gate: script prints SUBMIT_OK and receipt status: 1 (or equivalent).

Capture tx hash from output (denote it as TXHASH).

Step 8 — Receipt gate (T2)
bash
Copier le code
curl -s -X POST http://127.0.0.1:8545 -H 'content-type: application/json' \
  --data '{"jsonrpc":"2.0","id":20,"method":"eth_getTransactionReceipt","params":["TXHASH"]}'
Gate:

status = 0x1

blockNumber is present (denote it as BH)

logs length >= 1 and includes PoSSRegistry address

Step 9 — transactionsRoot/receiptsRoot/bloom must be non-empty on the tx block (T2)
Query the block that contains the tx (BH from receipt):

bash
Copier le code
curl -s -X POST http://127.0.0.1:8545 -H 'content-type: application/json' \
  --data '{"jsonrpc":"2.0","id":30,"method":"eth_getBlockByNumber","params":["BH",false]}'
Gate (leader):

transactionsRoot != empty-trie root

receiptsRoot != empty-trie root

logsBloom is non-zero

timestamp is stable (repeatable for BH)

Step 10 — Follower parity on the same tx block (T2)
bash
Copier le code
curl -s -X POST http://127.0.0.1:8546 -H 'content-type: application/json' \
  --data '{"jsonrpc":"2.0","id":31,"method":"eth_getBlockByNumber","params":["BH",false]}'
Gate (follower parity):

The following fields match leader exactly for BH:

number, hash, parentHash

stateRoot, transactionsRoot, receiptsRoot, logsBloom

timestamp

Step 11 — Restart invariants (optional but recommended) (T1/T2)
Stop nodes:

bash
Copier le code
pkill -INT noorcore || true
Start nodes again (same commands as Step 3), then re-run:

receipt query (Step 8) for TXHASH

block query (Step 9) for BH on leader

parity query (Step 10) for BH on follower

Gate: outputs remain identical (roots/bloom/timestamp for BH).

Evidence pack (minimal)
Store the following outputs in a text file (paste manually):

Step 4 (chainId + blockNumber) leader/follower

Step 8 receipt (shows blockNumber + logs)

Step 9 leader block BH (shows non-empty txRoot/rcptRoot/bloom)

Step 10 follower block BH (shows same fields)

Location suggestion:
/tmp/M23_EVIDENCE_MINIMAL.txt

Expected result
M23 is validated when:

Block metadata is persisted and served (stable timestamp + roots/bloom)

transactionsRoot is derived from raw tx RLP and becomes non-empty on tx blocks

receiptsRoot/logsBloom become non-empty on tx blocks

follower parity matches leader for the same tx block

restart invariants preserve the same values