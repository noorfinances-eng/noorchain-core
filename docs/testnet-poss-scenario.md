# NOORCHAIN Testnet 1.0 — Scénario de test PoSS

Ce document décrit un scénario de test simple pour valider que le module
PoSS (x/noorsignal) fonctionne correctement sur **noorchain-testnet-1**.

Objectif : vérifier que :

- un participant peut soumettre un signal,
- un curator peut valider ce signal,
- les récompenses NUR (unur) sont bien distribuées 70/30,
- les events PoSS sont émis,
- les queries gRPC retournent les bonnes informations.

---

## 1. Prérequis

- Chaîne démarrée avec le chain_id :

  ```text
  noorchain-testnet-1
