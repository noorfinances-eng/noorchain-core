RUNBOOK — M13 Validation (Restart Invariants + Follower Parity + eth_getLogs)

Document ID: RUNBOOK-M13-VALIDATION
Scope: Local mainnet-like pack (Leader/Follower), PoSSRegistry read/write verification, receipts/logs verification, restart invariants, follower parity.
Assumptions:

Repo: noorchain-core, branch evm-l1

Binary: ./noorcore (built from ./core)

Leader RPC: 127.0.0.1:8545

Follower RPC: 127.0.0.1:8546

ChainId: 0x849 (2121)

Data dirs: ./data/node1, ./data/node2

Health: 127.0.0.1:8081 (leader), 127.0.0.1:8082 (follower)

PoSSRegistry address used for validation: 0x560Ff18B151045561E9393bb8FaEA15C17BC02Fc

Dev curator used in examples: 0x4aA5DA75AFb6e81F433D4720cb7Cb2C6B1BA323c

Terminal Discipline

T1 = Nodes only (start/stop/ports/logs)

T2 = Tooling only (curl, scripts, verification)

One command per step. Validate gate output before continuing.

0) Clean Stop (mandatory)

T1

pkill -INT noorcore; sleep 1; pgrep -a noorcore || echo "OK: noorcore stopped"


Gate: OK: noorcore stopped

T1

ss -ltnp | grep -E '(:30303|:30304|:8545|:8546|:8081|:8082)\b' || echo "OK: ports closed"


Gate: OK: ports closed

1) Build (mandatory)

T2

cd /workspaces/noorchain-core && gofmt -w core/rpc/jsonrpc.go && go build -o noorcore ./core


Gate: build completes without errors.

2) Start Leader (node1)

T1

cd /workspaces/noorchain-core && mkdir -p ./logs && \
./noorcore -chain-id noorchain-2-1-local \
  -data-dir ./data/node1 \
  -p2p-addr 127.0.0.1:30303 \
  -rpc-addr 127.0.0.1:8545 \
  -health-addr 127.0.0.1:8081 >> ./logs/node1.log 2>&1 & echo PID=$!


Gate: PID printed.

T1

ss -ltnp | grep -E '(:30303|:8545|:8081)\b'


Gate: all three ports are LISTENing for noorcore.

T2

curl -s -X POST http://127.0.0.1:8545 -H 'content-type: application/json' \
  --data '{"jsonrpc":"2.0","id":1,"method":"eth_chainId","params":[]}' ; echo
curl -s -X POST http://127.0.0.1:8545 -H 'content-type: application/json' \
  --data '{"jsonrpc":"2.0","id":2,"method":"eth_blockNumber","params":[]}' ; echo


Gate: eth_chainId returns 0x849 and eth_blockNumber returns a non-empty hex quantity.

3) Start Follower (node2)

T1

cd /workspaces/noorchain-core && mkdir -p ./logs && \
./noorcore -chain-id noorchain-2-1-local \
  -data-dir ./data/node2 \
  -p2p-addr 127.0.0.1:30304 \
  -rpc-addr 127.0.0.1:8546 \
  -health-addr 127.0.0.1:8082 \
  -follow-rpc http://127.0.0.1:8545 \
  -boot-peers 127.0.0.1:30303 >> ./logs/node2.log 2>&1 & echo PID=$!


Gate: PID printed.

T1

ss -ltnp | grep -E '(:30304|:8546|:8082)\b'


Gate: all three follower ports are LISTENing.

T2

curl -s -X POST http://127.0.0.1:8546 -H 'content-type: application/json' \
  --data '{"jsonrpc":"2.0","id":1,"method":"eth_chainId","params":[]}' ; echo
curl -s -X POST http://127.0.0.1:8546 -H 'content-type: application/json' \
  --data '{"jsonrpc":"2.0","id":2,"method":"eth_blockNumber","params":[]}' ; echo


Gate: chainId 0x849 and blockNumber is a non-empty hex quantity.

4) PoSS Read Gates (Leader)

T2

4.1 snapshotCount()

curl -s -X POST http://127.0.0.1:8545 -H 'content-type: application/json' \
  --data '{"jsonrpc":"2.0","id":21,"method":"eth_call","params":[{"to":"0x560Ff18B151045561E9393bb8FaEA15C17BC02Fc","data":"0x098ab6a1"},"latest"]}' ; echo


Gate: returns a 32-byte hex word quantity.

4.2 latestSnapshotId()

curl -s -X POST http://127.0.0.1:8545 -H 'content-type: application/json' \
  --data '{"jsonrpc":"2.0","id":22,"method":"eth_call","params":[{"to":"0x560Ff18B151045561E9393bb8FaEA15C17BC02Fc","data":"0xe484cf32"},"latest"]}' ; echo


Gate: returns a 32-byte hex word quantity. For a fresh pilot it is typically equal to snapshotCount.

4.3 getSnapshot(1)

Replace 0x... with your ABI-encoded calldata for getSnapshot(uint256) if you already have it; otherwise reuse the known good payload from your validated test.

curl -s -X POST http://127.0.0.1:8545 -H 'content-type: application/json' \
  --data '{"jsonrpc":"2.0","id":23,"method":"eth_call","params":[{"to":"0x560Ff18B151045561E9393bb8FaEA15C17BC02Fc","data":"0x<GETSNAPSHOT_CALLDATA_FOR_1>"},"latest"]}' ; echo


Gate: non-empty ABI-encoded return payload for existing snapshot id, or a consistent empty/zeroed payload for non-existing id (depending on contract semantics).

5) PoSS Read Gates (Follower Parity)

Run the same three calls against follower RPC (8546) and compare results.

T2

# snapshotCount()
curl -s -X POST http://127.0.0.1:8546 -H 'content-type: application/json' \
  --data '{"jsonrpc":"2.0","id":31,"method":"eth_call","params":[{"to":"0x560Ff18B151045561E9393bb8FaEA15C17BC02Fc","data":"0x098ab6a1"},"latest"]}' ; echo

# latestSnapshotId()
curl -s -X POST http://127.0.0.1:8546 -H 'content-type: application/json' \
  --data '{"jsonrpc":"2.0","id":32,"method":"eth_call","params":[{"to":"0x560Ff18B151045561E9393bb8FaEA15C17BC02Fc","data":"0xe484cf32"},"latest"]}' ; echo

# getSnapshot(1)
curl -s -X POST http://127.0.0.1:8546 -H 'content-type: application/json' \
  --data '{"jsonrpc":"2.0","id":33,"method":"eth_call","params":[{"to":"0x560Ff18B151045561E9393bb8FaEA15C17BC02Fc","data":"0x<GETSNAPSHOT_CALLDATA_FOR_1>"},"latest"]}' ; echo


Gate: follower responses match leader responses byte-for-byte for the same block tag.

6) eth_getLogs Gate (Receipt Scan)

This gate validates that logs are returned for PoSS events using the minimal O(n) receipt scan.

6.1 Filter by address only

T2

curl -s -X POST http://127.0.0.1:8545 -H 'content-type: application/json' \
  --data '{"jsonrpc":"2.0","id":41,"method":"eth_getLogs","params":[{"address":"0x560Ff18B151045561E9393bb8FaEA15C17BC02Fc","fromBlock":"0x0","toBlock":"latest"}]}' ; echo


Gate: returns an array; for a chain with PoSS activity it must contain at least one log entry.

6.2 Filter by topic0 (event signature)

Use the event signature hash you validated (topic0). Example below uses the already observed topic0.
T2

curl -s -X POST http://127.0.0.1:8545 -H 'content-type: application/json' \
  --data '{"jsonrpc":"2.0","id":42,"method":"eth_getLogs","params":[{"address":"0x560Ff18B151045561E9393bb8FaEA15C17BC02Fc","fromBlock":"0x0","toBlock":"latest","topics":["0x22784526843c74fc6087008496473e93f622893f6eb447d27c0edca5f94fcfbd"]}]}' ; echo


Gate: returns an array; entries must show:

address equals PoSSRegistry

topics[0] equals the event signature hash

transactionHash, blockNumber, logIndex present

6.3 Follower parity for logs

T2

curl -s -X POST http://127.0.0.1:8546 -H 'content-type: application/json' \
  --data '{"jsonrpc":"2.0","id":43,"method":"eth_getLogs","params":[{"address":"0x560Ff18B151045561E9393bb8FaEA15C17BC02Fc","fromBlock":"0x0","toBlock":"latest","topics":["0x22784526843c74fc6087008496473e93f622893f6eb447d27c0edca5f94fcfbd"]}]}' ; echo


Gate: follower returns a compatible array (same logs for same range).

7) Restart Invariant Gate (Persistence)

This gate ensures that after a clean stop and restart, the following are preserved:

PoSS snapshots (contract storage)

receipts/logs for historical txs

stateRoot / receiptsRoot exposure via eth_getBlockByNumber (latest)

7.1 Stop both nodes

T1

pkill -INT noorcore; sleep 1; pgrep -a noorcore || echo "OK: noorcore stopped"


Gate: stopped.

7.2 Start leader then follower again

Repeat steps 2 and 3 exactly.

7.3 Re-run PoSS Read Gates + Logs Gate

Repeat steps 4, 5, 6.

Gate: results remain consistent across restart (snapshotCount/latestSnapshotId/getSnapshot + eth_getLogs still returns historical event(s)).

8) Block Metadata Gate (stateRoot / receiptsRoot)

T2

curl -s -X POST http://127.0.0.1:8545 -H 'content-type: application/json' \
  --data '{"jsonrpc":"2.0","id":99,"method":"eth_blockNumber","params":[]}' ; echo


Record the latest height H.

T2

curl -s -X POST http://127.0.0.1:8545 -H 'content-type: application/json' \
  --data '{"jsonrpc":"2.0","id":100,"method":"eth_getBlockByNumber","params":["latest",false]}' ; echo


Gate: response includes:

stateRoot is a non-zero 32-byte hash

receiptsRoot is present (valid 32-byte hash)

logsBloom is present (512 bytes hex / 0x + 1024 hex chars; may be zeroed in this milestone)

number matches eth_blockNumber

(Optional parity)
Run the same call on follower (8546) and verify parity.

9) Clean Stop (end of runbook)

T1

pkill -INT noorcore; sleep 1; pgrep -a noorcore || echo "OK: noorcore stopped"


T1

ss -ltnp | grep -E '(:30303|:30304|:8545|:8546|:8081|:8082)\b' || echo "OK: ports closed"

Expected Outcome (M13 “VALIDATED-FINAL”)

A run is considered PASS if:

Build succeeds.

Leader and follower start cleanly and serve RPC.

PoSS reads (snapshotCount/latestSnapshotId/getSnapshot) work on leader.

Follower returns parity for PoSS reads.

eth_getLogs returns PoSS event logs on leader (and parity-compatible on follower).

After restart, all above remain valid (persistence invariant).

eth_getBlockByNumber("latest") exposes non-zero stateRoot and stable roots format.