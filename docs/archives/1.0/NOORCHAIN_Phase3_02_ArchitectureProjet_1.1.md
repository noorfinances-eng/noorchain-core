> ⚠️ **Status: TECH MIGRATION**  
> This document is being realigned to the **NOORCHAIN 2.0 Technical Baseline**.  
> Reference: `NOORCHAIN_Tech_Baseline_2.0_1.1.md`  
> Branch: `main-3` — Date: 2025-12-18

NOORCHAIN — Phase 3
File 20 — Official Project Architecture (v1.1)
Updated: 2025-XX-XX
Language: English
🎯 Document Purpose

This document defines the official, stable, and complete architecture
of the noorchain-core repository for Phases 3 → 9
(Testnet → Mainnet).

It acts as the single source of truth for:

project structure,

file organisation,

module boundaries,

future implementation planning (PoSS, EVM, genesis, CLI, RPC).

Its mission is to prevent architectural drift and maintain long-term structural coherence.

1. Repository Root Structure

The noorchain-core project contains six main directories:

noorchain-core/
│
├── app/          → Core Cosmos application (BaseApp, keepers, modules)
├── cmd/          → CLI binaries (noord)
├── x/            → Custom modules (including PoSS)
├── proto/        → Protobuf definitions
├── testnet/      → Genesis files and testnet configuration
└── docs/         → Official documentation


This structure will remain stable across Testnet and Mainnet.

2. Directory Details
2.1 app/ — Core Cosmos Application

This directory contains the heart of NOORCHAIN.
The files below will be progressively implemented in Phase 4:

app.go — Main application definition

app_builder.go — Extension-ready app builder

config.go — NOORCHAIN configuration (Bech32 prefixes, denom, version)

encoding.go — Codec (Protobuf + Amino)

keepers.go — Global keepers declaration

module_manager.go — Module registration and execution order

params.go — Global parameter handling

types.go — Internal app types

Important Note:
In Phase 3, we define structure only. No Cosmos logic is implemented yet.

2.2 cmd/noord/ — Node Binary

This directory contains the executable entry point of the blockchain:

main.go — The NOORCHAIN node main function

Future additions (Phase 5+):

initialization commands

genesis tools

unsafe reset

key utilities

node start commands

2.3 x/ — Custom Modules

This is the official location for NOORCHAIN modules.

Current + future content:

noorsignal/ — PoSS (Proof of Signal Social) module

keeper

types

module logic

genesis handlers

message definitions

events

During Phase 4, we will implement a minimal empty skeleton,
then gradually add PoSS logic.

2.4 proto/ — Protobuf Definitions

All protocol-level .proto files are stored here.

Expected future structure:

proto/noorsignal/tx.proto

proto/noorsignal/query.proto

proto/noorsignal/types.proto

proto/noorsignal/genesis.proto

Proto generation is deferred until Phase 4.

2.5 testnet/ — Testnet & Genesis Configuration

This directory will be populated in Phase 6 with:

genesis.json

genesis_distribution.json

config.toml

addrbook.json

chain-id definition

persistent peer lists

The directory remains empty during Phase 3.

2.6 docs/ — Official Documentation

The /docs directory contains all documentation files,
versioned 1.1, including:

technical architecture

governance

legal framework

PoSS specifications

genesis constraints

upgrade plans

This architecture document lives in:

docs/NOORCHAIN_Phase3_02_ArchitectureProjet_1.1.md

3. Official Architectural Rules
3.1 No Unnecessary Cosmos Files

Only add Cosmos-related files in Phase 4.
Phase 3 remains strictly conceptual.

3.2 PoSS Module Must Be Fully Isolated

PoSS logic must only appear in x/noorsignal/.
It must not contaminate:

app/

cmd/

runtime/

shared stores

3.3 Phase 3 Files Must Stay Simple

No logic, handlers, or keepers are written yet.
Only structure and future placeholders.

3.4 Documentation is the Single Source of Truth

If a file or directory is not described in /docs,
it must not be created.

3.5 No Code Generation Without Explicit Approval

No proto generation, buf setup, or Makefile expansion occurs during Phase 3.

4. Executive Summary

The noorchain-core repository is composed of six stable directories:

app/ — Cosmos application core

cmd/noord/ — Node binary

x/ — Custom modules (including PoSS)

proto/ — Protobuf definitions

testnet/ — Network configuration

docs/ — Full project documentation

This architecture is final, approved, and versioned.
It guarantees order, modularity, and future-proof development.

5. Status

Validated:
This architecture is officially adopted as Phase 3 baseline
and will guide all future implementation phases.
