
All module interactions must follow a strict order to ensure:

- determinism  
- consensus correctness  
- EVM state stability  
- PoSS reward correctness  
- governance consistency  

---

# 3. BeginBlock Lifecycle (Authoritative Order)

### **BeginBlock MUST execute in this exact sequence:**

1. **FeeMarket**  
2. **EVM**  
3. **Staking**  
4. **PoSS (noorsignal)**  
5. **Governance**

This order is consensus-critical.

---

## 3.1 FeeMarket BeginBlock

The FeeMarket module:

- updates the EIP-1559 base fee  
- applies gas dynamics for the upcoming block  
- ensures proper pricing for EVM execution  

---

## 3.2 EVM BeginBlock

The EVM module:

- prepares the EVM block context  
- resets EVM logs  
- initializes block-level variables:  
  - base fee  
  - coinbase  
  - block time  
  - block height  

This must *always* occur before processing any Ethereum tx.

---

## 3.3 Staking BeginBlock

The Staking module:

- processes validator updates  
- recomputes voting power  
- handles unjailing events  
- handles slashing-related state  

Staking must always run before PoSS, because PoSS depends on:

- validator power  
- accurate chain timestamp  

---

## 3.4 PoSS BeginBlock

The PoSS module (x/noorsignal):

- reads PoSS parameters from ParamsKeeper  
- checks whether PoSS is enabled  
- evaluates halving epoch  
- updates daily counters  
- prepares state for reward computation  
- (future) processes pending signals  
- (future) applies the 70/30 reward distribution  

PoSS may not modify validator sets or gas rules.

---

## 3.5 Governance BeginBlock

The Governance module:

- checks active proposals  
- advances voting periods  
- closes proposals when conditions are met  

Governance must always run **after** PoSS.

---

# 4. DeliverTx Lifecycle

Every transaction is processed through the following deterministic pipeline:

---

## 4.1 AnteHandler

The AnteHandler performs:

- signature verification  
- account sequence checks  
- account number validation  
- gas payer balance checks  
- EVM-specific validation (for `MsgEthereumTx`)  

If the AnteHandler fails → the transaction fails before message execution.

---

## 4.2 Message Routing

The Router dispatches each message to the correct module:

- Cosmos SDK messages  
- Staking & Governance messages  
- PoSS messages (`MsgCreateSignal`, etc.)  
- Ethereum transactions (`MsgEthereumTx`)  

---

## 4.3 State Transitions

Modules perform:

- KVStore writes  
- event emission  
- reward distribution (when PoSS is enabled)  
- EVM bytecode execution  
- logs and receipts generation  

Gas is fully tracked throughout this stage.

---

## 4.4 EVM-Specific Execution

For `MsgEthereumTx`:

- execution occurs inside the EVM stateDB  
- state only commits if the tx is successful  
- errors trigger a full EVM revert  
- logs, receipts, and bloom filters are produced  

EVM execution must be 100% deterministic across all nodes.

---

# 5. EndBlock Lifecycle

EndBlock must run in the following strict order:

1. **Staking**  
2. **Governance**

---

## 5.1 Staking EndBlock

Staking:

- finalizes validator updates  
- outputs a new validator set diff  
- provides updates to CometBFT  

This is the only place where validator updates are committed.

---

## 5.2 Governance EndBlock

Governance:

- finalizes proposals whose voting has ended  
- applies accepted changes  
- rejects or marks failed proposals  

Governance must run after staking to ensure valid voting power calculations.

---

# 6. Commit Phase

The Commit phase finalizes the block.

---

## 6.1 State Commitment

The multistore:

- commits all module states via IAVL  
- produces the new app hash  
- anchors this hash for the next block  

This is the consensus root of the chain.

---

## 6.2 EVM State Commit

EVM performs:

- stateDB flush  
- log indexing  
- bloom filter generation  
- transaction receipt persistence  

This must occur **before** PoSS commit.

---

## 6.3 PoSS Commit

PoSS:

- commits daily counters  
- commits halving epoch index  
- commits reward metadata  
- prepares internal caches for next block  

PoSS commit must occur **after** EVM commit.

---

# 7. Lifecycle Diagram
            ▼
    ┌──────────────────┐
    │   BeginBlock     │
    └──────────────────┘
feemarket → evm → staking → poss → gov
▼
┌──────────────────┐
│ DeliverTx │
└──────────────────┘
ante → routing → state writes
▼
┌──────────────────┐
│ EndBlock │
└──────────────────┘
staking → gov
▼
┌──────────────────┐
│ Commit │
└──────────────────┘
state commit → evm commit → poss commit

---

# 8. Determinism Constraints

These rules are **mandatory** for consensus safety:

- FeeMarket must run **before** EVM  
- Staking must run **before** PoSS  
- PoSS must run **before** Governance  
- EVM must commit **before** PoSS commit  
- PoSS must **never** run in EndBlock  
- EndBlock must contain only staking → governance  
- No module may modify validator sets except staking  

Breaking these rules leads to:

- reward inconsistencies  
- EVM mismatches  
- invalid validator power  
- non-deterministic state transitions  
- governance state corruption  

---

# 9. Summary

NOORCHAIN’s block lifecycle is:

| Phase       | Order |
|-------------|-------|
| **BeginBlock** | feemarket → evm → staking → poss → gov |
| **DeliverTx**  | ante → routing → state writes |
| **EndBlock**   | staking → gov |
| **Commit**     | state commit → evm commit → poss commit |

This document is the **canonical reference** for block execution.  
All coding in Phase 4C and all audits in Phase 5–7 must follow it.


