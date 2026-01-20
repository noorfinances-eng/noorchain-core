NOORCHAIN 2.1 — JSON-RPC Examples

Document ID: JSON_RPC_EXAMPLES_2.1
Version: v1.0
Date: January 2026
Status: Active
Scope: Canonical JSON-RPC request/response examples for NOORCHAIN 2.1, intended for tooling integration, debugging, and compatibility gates.

1. Purpose

This document provides deterministic JSON-RPC examples that:

demonstrate request and response formatting

serve as a compatibility gate reference

support operator evidence capture

reduce ambiguity for integrators

This is not a complete RPC specification. Normative semantics are defined in:

docs/RPC_SPEC_2.1.md

docs/rpc/RPC_ERROR_MODEL_2.1.md

Examples are written using curl against a localhost RPC endpoint.

2. Conventions

Replace 127.0.0.1:8545 with your RPC address.

Responses shown are illustrative; specific values (block numbers, hashes) will differ.

JSON-RPC envelope uses:

"jsonrpc":"2.0"

"id": <integer>

"method": "<method>"

"params": [...]

Hex quantities follow Ethereum-style rules (no leading zeros unless the value is zero).

3. Liveness and Identity
3.1 eth_chainId

Request:

curl -s http://127.0.0.1:8545 \
  -H 'content-type: application/json' \
  --data '{"jsonrpc":"2.0","id":1,"method":"eth_chainId","params":[]}'
echo


Response (example):

{"jsonrpc":"2.0","id":1,"result":"0x849"}

3.2 eth_blockNumber

Request:

curl -s http://127.0.0.1:8545 \
  -H 'content-type: application/json' \
  --data '{"jsonrpc":"2.0","id":2,"method":"eth_blockNumber","params":[]}'
echo


Response (example):

{"jsonrpc":"2.0","id":2,"result":"0x12"}

4. Block Reads
4.1 eth_getBlockByNumber (latest, minimal fields)

Request:

curl -s http://127.0.0.1:8545 \
  -H 'content-type: application/json' \
  --data '{"jsonrpc":"2.0","id":10,"method":"eth_getBlockByNumber","params":["latest",false]}'
echo


Response (example):

{
  "jsonrpc":"2.0",
  "id":10,
  "result":{
    "number":"0x12",
    "hash":"0xaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
    "parentHash":"0xbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb",
    "stateRoot":"0xcccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccc",
    "receiptsRoot":"0xdddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddd",
    "logsBloom":"0x" 
  }
}


Notes:

Field coverage and semantics are defined in docs/RPC_SPEC_2.1.md.

logsBloom may be empty or a fixed-length bloom depending on the current implementation scope.

4.2 eth_getBlockByNumber (specific height)

Request:

curl -s http://127.0.0.1:8545 \
  -H 'content-type: application/json' \
  --data '{"jsonrpc":"2.0","id":11,"method":"eth_getBlockByNumber","params":["0x1",false]}'
echo


Response (example):

{"jsonrpc":"2.0","id":11,"result":{"number":"0x1","hash":"0x...","stateRoot":"0x...","receiptsRoot":"0x...","logsBloom":"0x..."}}

4.3 eth_getBlockByNumber (> latest height returns null)

Request:

curl -s http://127.0.0.1:8545 \
  -H 'content-type: application/json' \
  --data '{"jsonrpc":"2.0","id":12,"method":"eth_getBlockByNumber","params":["0xffffffffffff",false]}'
echo


Response (example):

{"jsonrpc":"2.0","id":12,"result":null}

5. Account / World-State Reads
5.1 eth_getBalance

Request:

curl -s http://127.0.0.1:8545 \
  -H 'content-type: application/json' \
  --data '{"jsonrpc":"2.0","id":20,"method":"eth_getBalance","params":["0x4aA5DA75AFb6e81F433D4720cb7Cb2C6B1BA323c","latest"]}'
echo


Response (example):

{"jsonrpc":"2.0","id":20,"result":"0x56bc75e2d63100000"}


Notes:

Balance is a hex quantity (wei).

The authoritative semantics are defined in docs/RPC_SPEC_2.1.md.

5.2 eth_getTransactionCount (nonce)

Request:

curl -s http://127.0.0.1:8545 \
  -H 'content-type: application/json' \
  --data '{"jsonrpc":"2.0","id":21,"method":"eth_getTransactionCount","params":["0x4aA5DA75AFb6e81F433D4720cb7Cb2C6B1BA323c","latest"]}'
echo


Response (example):

{"jsonrpc":"2.0","id":21,"result":"0x2"}

5.3 eth_getCode

Request:

curl -s http://127.0.0.1:8545 \
  -H 'content-type: application/json' \
  --data '{"jsonrpc":"2.0","id":22,"method":"eth_getCode","params":["0x1000000000000000000000000000000000000001","latest"]}'
echo


Response (example):

{"jsonrpc":"2.0","id":22,"result":"0x6000..."}


If no code exists at the address:

{"jsonrpc":"2.0","id":22,"result":"0x"}

5.4 eth_getStorageAt

Request:

curl -s http://127.0.0.1:8545 \
  -H 'content-type: application/json' \
  --data '{"jsonrpc":"2.0","id":23,"method":"eth_getStorageAt","params":["0x1000000000000000000000000000000000000001","0x0","latest"]}'
echo


Response (example):

{"jsonrpc":"2.0","id":23,"result":"0x0000000000000000000000000000000000000000000000000000000000000000"}

6. Transaction Path Examples (When Enabled)

This section provides example call shapes. Actual tx hashes and receipts depend on a funded account, signing, and the node’s transaction inclusion behavior.

6.1 eth_sendRawTransaction

Request (example shape):

curl -s http://127.0.0.1:8545 \
  -H 'content-type: application/json' \
  --data '{"jsonrpc":"2.0","id":30,"method":"eth_sendRawTransaction","params":["0x02f8..."]}'
echo


Response (example):

{"jsonrpc":"2.0","id":30,"result":"0xaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}

6.2 eth_getTransactionByHash

Request:

curl -s http://127.0.0.1:8545 \
  -H 'content-type: application/json' \
  --data '{"jsonrpc":"2.0","id":31,"method":"eth_getTransactionByHash","params":["0xaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"]}'
echo


Response (example):

{
  "jsonrpc":"2.0",
  "id":31,
  "result":{
    "hash":"0xaaaaaaaa...",
    "from":"0x4aA5DA75AFb6e81F433D4720cb7Cb2C6B1BA323c",
    "to":"0x1000000000000000000000000000000000000001",
    "nonce":"0x2",
    "input":"0x...",
    "value":"0x0"
  }
}

6.3 eth_getTransactionReceipt

Request:

curl -s http://127.0.0.1:8545 \
  -H 'content-type: application/json' \
  --data '{"jsonrpc":"2.0","id":32,"method":"eth_getTransactionReceipt","params":["0xaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"]}'
echo


Response (example):

{
  "jsonrpc":"2.0",
  "id":32,
  "result":{
    "transactionHash":"0xaaaaaaaa...",
    "blockNumber":"0x12",
    "status":"0x1",
    "contractAddress":null,
    "logs":[]
  }
}


Notes:

Receipt fields and expectations are normative in docs/RPC_SPEC_2.1.md.

Logs presence depends on contract execution and event emission.

7. PoSS Registry Read Examples (If Deployed)

This section provides generic eth_call shapes for contract reads. Concrete selectors and addresses are environment-specific. Use the deployed PoSSRegistry address for your environment.

7.1 eth_call — snapshotCount()

Request (example):

curl -s http://127.0.0.1:8545 \
  -H 'content-type: application/json' \
  --data '{
    "jsonrpc":"2.0",
    "id":40,
    "method":"eth_call",
    "params":[
      {"to":"0x560Ff18B151045561E9393bb8FaEA15C17BC02Fc","data":"0x098ab6a1"},
      "latest"
    ]
  }'
echo


Response (example):

{"jsonrpc":"2.0","id":40,"result":"0x0000000000000000000000000000000000000000000000000000000000000001"}

7.2 eth_call — latestSnapshotId()

Request (example):

curl -s http://127.0.0.1:8545 \
  -H 'content-type: application/json' \
  --data '{
    "jsonrpc":"2.0",
    "id":41,
    "method":"eth_call",
    "params":[
      {"to":"0x560Ff18B151045561E9393bb8FaEA15C17BC02Fc","data":"0xe484cf32"},
      "latest"
    ]
  }'
echo


Response (example):

{"jsonrpc":"2.0","id":41,"result":"0x0000000000000000000000000000000000000000000000000000000000000001"}

8. Error Examples

Error formats are normative in:

docs/rpc/RPC_ERROR_MODEL_2.1.md

8.1 Method Not Found

Response (example):

{
  "jsonrpc":"2.0",
  "id":90,
  "error":{"code":-32601,"message":"Method not found"}
}

8.2 Invalid Params

Response (example):

{
  "jsonrpc":"2.0",
  "id":91,
  "error":{"code":-32602,"message":"Invalid params"}
}

8.3 Internal Error

Response (example):

{
  "jsonrpc":"2.0",
  "id":92,
  "error":{"code":-32603,"message":"Internal error"}
}

9. References

docs/RPC_SPEC_2.1.md

docs/rpc/RPC_ERROR_MODEL_2.1.md

docs/rpc/RPC_COMPAT_MATRIX_2.1.md

docs/API_STABILITY_POLICY_2.1.md

docs/STATE_MODEL_2.1.md

docs/dev/TOOLING_HARDHAT_VIEM_2.1.md