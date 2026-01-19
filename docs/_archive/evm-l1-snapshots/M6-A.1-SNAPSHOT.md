# NOORCHAIN 2.1 — M6-A.1 Snapshot

## Purpose
Technical baseline freeze after M5 (persistent P2P network).

## Date
2025-12-25

## Branch
evm-l1

## Reference commit
68584a5

## Tag
M5-STABLE

## Environment
- Go: 1.25.5
- OS: Linux (Codespaces)

## Frozen scope
- Sovereign EVM L1 core
- Configurable boot peers
- Persistent P2P sessions
- Automatic network startup

## Off-Git archive
- File: M5.tar.gz
- SHA-256: fe63cce5cf3170f41cdcc5233003615b95387eded55daacc1f7de04e4b24b97a

## Validation
- Two-node network stable
- Reconnection after restart validated
- No CLOSE-WAIT accumulation observed
