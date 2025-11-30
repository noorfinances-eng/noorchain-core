# NOORCHAIN — Internal Testnet 1 (squelette)

Ce dossier contient les fichiers de base pour le **Testnet interne 1** de NOORCHAIN.
Pour l’instant, il s’agit d’un **squelette** : aucun état applicatif réel, aucun validateur configuré.

---

## 🎯 Objectifs de Testnet 1

- Définir un **`chain_id` propre** pour le testnet :  
  `noorchain-testnet-1`
- Créer un **modèle de fichier `genesis`** qui servira de base aux prochaines étapes :
  - Ajout des comptes de test
  - Ajout des validateurs
  - Ajout des modules PoSS (x/noorsignal), EVM, etc.

---

## 📁 Fichiers

- `genesis.noorchain-testnet-1.template.json`  
  → Modèle de genesis, **non utilisable en production** tel quel.  
  L’objectif est d’avoir un point de départ homogène pour les futures étapes Testnet 2 / Testnet 3.

---

## 🚧 État actuel

- Aucun nœud NOORCHAIN n’est configuré pour utiliser ce genesis.
- Les modules `auth`, `bank`, `staking`, `evm`, `feemarket`, `noorsignal` n’ont pas encore de configuration de genesis spécifique.
- Ce dossier sert uniquement de **base documentaire et technique** pour les prochaines étapes.

Les prochaines étapes (Testnet 2, Testnet 3, …) ajouteront :

- Des comptes de test (fondation, fondateur, réserve PoSS…)
- Un validateur de test
- Les params de chaines (staking, gov, EVM, FeeMarket, PoSS)
- Les scripts pour lancer un réseau local (un seul nœud, puis multi-nœuds).

---
