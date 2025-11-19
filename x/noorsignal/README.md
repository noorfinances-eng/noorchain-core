# NOORCHAIN — PoSS Module (`x/noorsignal`)

This module implements the core of **Proof of Signal Social (PoSS)** for NOORCHAIN.

> Status: early technical skeleton — types, keeper and config are present,
> but no reward or halving logic is implemented yet.

---

## 1. Purpose

The `noorsignal` module is responsible for:

- recording **social signals** emitted by participants
- tracking and managing **curators** (social validators)
- storing and applying the **global PoSS configuration**:
  - base reward
  - 70% / 30% split (participant / curator)
  - daily limits
  - enable/disable flag

The actual **minting logic, reward distribution and halving** will be implemented
in later phases and wired to the global NUR supply / PoSS rules.

---

## 2. Current Code Structure

### 2.1 Types

Located in: `x/noorsignal/types/types.go`

- `Signal`  
  Represents a social signal emitted by a participant:
  - `Id uint64`
  - `Participant sdk.AccAddress`
  - `Curator sdk.AccAddress`
  - `Weight uint32` (1x, 2x, 5x… encoded as integers)
  - `Time time.Time`
  - `Metadata string` (content hash, external ID, etc.)

- `Curator`  
  Represents a social validator:
  - `Address sdk.AccAddress`
  - `Level string` (e.g. “bronze”, “silver”, “gold”)
  - `TotalSignalsValidated uint64`
  - `Active bool`

- `PossConfig`  
  Global PoSS configuration:
  - `BaseReward uint64`
  - `ParticipantShare uint32` (e.g. 70)
  - `CuratorShare uint32` (e.g. 30)
  - `MaxSignalsPerDay uint32`
  - `Enabled bool`

- `DefaultPossConfig()`  
  Returns a default configuration coherent with NOORCHAIN’s model:
  - 70% participant / 30% curator
  - module enabled
  - symbolic `BaseReward` and daily limits.

> Note: these Go structs are the conceptual model.  
> Final protobuf definitions (`.proto`) will be added later.

---

### 2.2 Keys and Store Layout

Located in: `x/noorsignal/types/keys.go`

Defines:

- `ModuleName = "noorsignal"`
- `StoreKey`, `RouterKey`, `QuerierRoute` (standard Cosmos conventions)

Store prefixes:

- `KeyPrefixSignals = []byte{0x01}`
- `KeyPrefixCurators = []byte{0x02}`
- `KeyPrefixConfig = []byte{0x03}`

Helper functions:

- `GetSignalStore(parent prefix.Store) prefix.Store`
- `GetCuratorStore(parent prefix.Store) prefix.Store`
- `GetConfigStore(parent prefix.Store) prefix.Store`

These functions are used by the keeper to obtain scoped KVStores.

---

### 2.3 Keeper

Located in: `x/noorsignal/keeper/keeper.go`

The `Keeper` struct currently holds:

- `storeKey storetypes.StoreKey`
- `cdc codec.Codec`

Constructor:

- `NewKeeper(cdc codec.Codec, storeKey storetypes.StoreKey)`

Internal helpers:

- `getStore(ctx sdk.Context) sdk.KVStore`
- `signalStore(ctx sdk.Context) prefix.Store`
- `curatorStore(ctx sdk.Context) prefix.Store`
- `configStore(ctx sdk.Context) prefix.Store`

Config management:

- `SetConfig(ctx sdk.Context, cfg PossConfig)`  
  Encodes and stores the global PoSS configuration.

- `GetConfig(ctx sdk.Context) (PossConfig, bool)`  
  Reads and decodes the configuration; returns `(config, found)`.

- `InitDefaultConfig(ctx sdk.Context)`  
  If no configuration is present, stores `DefaultPossConfig()`.  
  If a configuration already exists, it is preserved.

No signal or curator logic is implemented yet (no writes/reads for them).

---

## 3. Planned Extensions

The following elements are planned but not yet implemented:

1. **Protobuf Definitions (`.proto`)**
   - `noorsignal/signal.proto`
   - `noorsignal/curator.proto`
   - `noorsignal/config.proto`
   - `noorsignal/tx.proto` for Msg definitions

2. **Message Types (Tx)**
   Examples:
   - `MsgSubmitSignal`
   - `MsgValidateSignal`
   - `MsgUpdatePossConfig` (governance or curator-limited)

3. **Genesis Handling**
   - `InitGenesis` and `ExportGenesis` functions
   - default PoSS config included in genesis
   - optional pre-registered curators

4. **Reward Logic**
   - calculation of PoSS rewards per signal
   - enforcement of halving rules (linked to global NUR supply)
   - daily limits per participant
   - enforcement of 70% / 30% split

5. **Integration with Other Modules**
   - `BankKeeper` for NUR transfers
   - potentially `StakingKeeper` or `GovKeeper` for curator governance
   - interaction with global supply management for PoSS

---

## 4. Development Notes

- The current implementation is intentionally minimal and safe:
  - no minting
  - no reward distribution
  - no halving logic
- This skeleton allows:
  - early review of the design
  - step-by-step implementation of PoSS features
  - future audit of the module in isolation

This README should be updated as soon as:

- protobuf files are added,
- messages are implemented,
- genesis and reward logic are wired.
