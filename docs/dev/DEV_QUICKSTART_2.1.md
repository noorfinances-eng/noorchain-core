NOORCHAIN 2.1 — Developer Quickstart

Document ID: DEV_QUICKSTART_2.1
Version: v1.0
Date: January 2026
Status: Active
Scope: Local developer workflow for building and running NOORCHAIN 2.1, validating JSON-RPC liveness, and executing basic tooling flows in controlled environments.

1. Purpose

This quickstart provides a deterministic baseline for developers to:

build the node binary

start a local node (or leader/follower pack)

validate JSON-RPC liveness and basic semantics

run basic contract/tooling workflows (where supported)

This document is not a full spec. Normative behavior is defined in:

docs/RPC_SPEC_2.1.md

docs/STATE_MODEL_2.1.md

2. Prerequisites
2.1 Required Tooling

Go toolchain (version pinned by repository requirements)

Standard shell tools (curl, ss, pgrep)

Node.js tooling (only if using Hardhat/viem scripts)

2.2 Repository Layout Assumptions

This guide assumes:

you are working from repository root

node binary is built from ./core

data directories are under ./data/* (local) unless otherwise specified

logs may be written to ./logs/*

If your environment differs, adapt paths consistently.

3. Baseline Discipline
3.1 One Change at a Time

For reliability:

verify state

make one change

verify again

Avoid mixing node lifecycle and tooling in the same terminal session.

3.2 Separate Terminals

T1: node process start/stop + node logs

T2: build, curl calls, deploy scripts, inspections

4. Build

From repository root:

go build -o noorcore ./core


Validation:

binary ./noorcore exists

no build errors

do not assume go build ./... updated the node binary; build the ./core target explicitly

5. Start a Single Local Node (Leader)
5.1 Prepare Directories
mkdir -p ./logs
mkdir -p ./data/node1

5.2 Start Node (Example)

Use explicit binds and log redirection:

./noorcore \
  -chain-id noorchain-2-1-local \
  -data-dir ./data/node1 \
  -p2p-addr 127.0.0.1:30303 \
  -rpc-addr 127.0.0.1:8545 \
  -health-addr 127.0.0.1:8081 \
  >> ./logs/node1.log 2>&1 &
echo PID=$!


This is a local example. Exact flags and validated packs may differ by milestone/runbook.

6. Liveness Gates (Single Node)
6.1 Process Gate
pgrep -a noorcore

6.2 Ports Gate
ss -ltnp | egrep 'noorcore|:30303|:8545|:8081'

6.3 RPC Liveness
curl -s http://127.0.0.1:8545 \
  -H 'content-type: application/json' \
  --data '{"jsonrpc":"2.0","id":1,"method":"eth_chainId","params":[]}'
echo

curl -s http://127.0.0.1:8545 \
  -H 'content-type: application/json' \
  --data '{"jsonrpc":"2.0","id":2,"method":"eth_blockNumber","params":[]}'
echo


Expected:

eth_chainId returns a hex quantity (chain-specific)

eth_blockNumber returns a hex quantity

Normative response formats are defined in docs/RPC_SPEC_2.1.md.

6.4 Health (If Enabled)
curl -s http://127.0.0.1:8081/health || true
echo

7. Stop Node (Clean)
pkill -INT noorcore || true
sleep 1
pgrep -a noorcore || true


Confirm ports are no longer bound:

ss -ltnp | egrep 'noorcore|:30303|:8545|:8081' || true

8. Start a Local 2-Node Pack (Leader + Follower)

Use the validated multi-node runbook as the primary source of truth:

docs/RUNBOOK-M10-MAINNETLIKE.md (or the current runbook for your tag)

8.1 General Pattern

leader and follower must use distinct:

data dirs

P2P ports

RPC ports

health ports

follower should be configured with -follow-rpc <leaderURL> when required by the deployment model.

8.2 Minimal Parity Gates

After both nodes are running:

eth_chainId equal on both

eth_blockNumber non-decreasing

eth_getBlockByNumber (latest) returns consistent metadata

parity-gated world-state reads return the same results (as specified)

The definitive parity scope is in:

docs/rpc/RPC_COMPAT_MATRIX_2.1.md

9. Tooling (Hardhat / viem) — Baseline Guidance

Tooling workflows depend on the repo’s configured networks and scripts.

9.1 Avoid Common Failure Modes

If tooling fails with connection errors:

confirm the leader RPC is listening (ss -ltnp)

confirm RPC liveness (eth_chainId)

confirm you are using the correct network name and config file for the repo

9.2 Secrets Handling

Never commit private keys.

Provide private keys at runtime via environment variables or interactive shell input.

10. Troubleshooting
10.1 RPC Connection Refused

check process exists (pgrep -a noorcore)

check port binding (ss -ltnp)

inspect logs (tail -n 120 ./logs/node1.log)

10.2 LevelDB Lock Error

ensure no second node is using the same -data-dir

stop conflicting processes

do not delete lock files unless you have proven no process holds the DB

10.3 “Address already in use”

identify owning PID via ss -ltnp

reconfigure ports (especially health port) to avoid collisions

11. References

docs/NOORCHAIN_Index_2.1.md

docs/OPERATIONS_PLAYBOOK_2.1.md

docs/ARCHITECTURE_2.1.md

docs/STATE_MODEL_2.1.md

docs/RPC_SPEC_2.1.md

docs/rpc/RPC_ERROR_MODEL_2.1.md

docs/rpc/RPC_COMPAT_MATRIX_2.1.md

docs/RELEASE_PROCESS_2.1.md

Runbooks under docs/RUNBOOK-*.md