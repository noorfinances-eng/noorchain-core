# NOORCHAIN — Testnet 1.0 Start Flow (InitChain)

Ce document décrit ce qui se passe lors du tout premier `noord start`
sur le testnet NOORCHAIN (noorchain-testnet-1).

---

## 1. Lecture du genesis

Le node lit `testnet/genesis.json` :

- `chain_id = "noorchain-testnet-1"`
- `genesis_time = "2025-01-01T00:00:00Z"`
- `app_state` contient au minimum :
  - `noorsignal` (config PoSS + curators vides)
  - `bank` (économie NUR, ajoutée par ApplyEconomicGenesis)

---

## 2. Construction de l'application NOORCHAIN

`NewNoorchainAppWithCosmos(...)` fait :

- crée un `AppBuilder`
- construit :
  - `BaseApp`
  - `AppKeepers` (Account, Bank, Params, NoorSignal)
  - `AppModules` (Manager + Configurator)
- enregistre :
  - `BeginBlocker`
  - `EndBlocker`
  - `InitChainer`

---

## 3. InitChainer (genesis économique + modules)

`InitChainer` exécute :

1. Décodage du genesis :
   ```go
   genesisState map[string]json.RawMessage
