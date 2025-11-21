# NOORCHAIN Testnet 1.0 — PoSS Dry Run (simulation “papier”)

Ce document décrit, étape par étape, ce qui se passe dans NOORCHAIN
lorsqu’on exécute un cycle PoSS minimal :

1. Un participant soumet un signal social (MsgSubmitSignal)
2. Un curator valide ce signal (MsgValidateSignal)
3. Les récompenses NUR sont calculées et distribuées (70/30)
4. Les events et queries reflètent l’état final

Ce scénario est une **simulation papier**, utile pour :

- vérifier la cohérence du design PoSS,
- préparer les tests réels du testnet,
- expliquer le fonctionnement à des auditeurs / partenaires.

---

## 1. Préconditions (avant le signal)

On suppose que :

- Le testnet est lancé avec le genesis décrit dans :
  - `testnet/genesis.json`
  - `testnet/genesis_distribution.json`
- Les adresses bech32 ont été remplacées correctement dans :
  - `testnet/genesis.json`
  - `docs/testnet-addresses-master.json`
  - `x/noorsignal/types/addresses.go`
- La réserve PoSS (80 %) possède **239 833 966.4 NUR**
  - soit `239833966400000` en `unur`.
- La configuration PoSS par défaut est :

```json
{
  "base_reward": 100,
  "participant_share": 70,
  "curator_share": 30,
  "max_signals_per_day": 50,
  "enabled": true,
  "era_index": 0
}
