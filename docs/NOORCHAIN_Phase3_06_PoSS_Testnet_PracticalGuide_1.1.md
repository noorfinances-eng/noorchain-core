# NOORCHAIN 1.0 — PoSS Testnet Practical Guide (v1.1)

**Scope**  
This document explains how to *practically* test the PoSS module (`x/noorsignal`) on a NOORCHAIN testnet, based on the current code state:

- PoSSEnabled = `false` by default (no real minting).
- Signals can already be created and counted.
- Rewards are computed theoretically but always `0/0` in live mode.

It complements:

- `NOORCHAIN_Phase3_03_PoSS_Specs_1.1.md`
- `NOORCHAIN_Phase3_05_PoSS_Status_and_Testnet_1.1.md`

---

## 1. Testnet Setup (minimal)

### 1.1. Init a local testnet

```bash
noord init noorchain-testnet --chain-id noorchain-1
