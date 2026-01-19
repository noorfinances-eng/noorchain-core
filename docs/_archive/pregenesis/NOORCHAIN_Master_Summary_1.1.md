NOORCHAIN — MASTER SUMMARY 2025
Synthèse Officielle — Version 1.1
Last Updated: 2025-12-XX
1. Vision

NOORCHAIN est une blockchain suisse éthique, construite en combinant :

Cosmos SDK 0.46.11

CometBFT 0.34.27 (regen-network replace)

Ethermint 0.22.0 (compatibilité EVM)

et un mécanisme unique :

⭐ Proof of Signal Social (PoSS)

Un système de validation sociale distribuant des récompenses non financières pour :

micro-dons,

participations vérifiées,

contenus éducatifs,

contributions CCN.

Objectif : favoriser les comportements positifs et renforcer le rôle social des communautés, écoles, ONG et associations.

NOORCHAIN se veut :

non spéculative,

utile,

transparente,

éthique,

conforme au Swiss Legal Light CH.

2. Éléments fondamentaux
2.1 Supply Fixe — “Vitesse de la lumière”

Supply totale : 299 792 458 NUR
Aucune inflation. Aucune émission supplémentaire.

2.2 Halving

Réduction automatique de l’émission PoSS tous les 8 ans.

2.3 Répartition économique (modèle suisse officiel)
Allocation	Pourcentage	Description
Fondation	5%	Non-profit, gouvernance, audits, transparence
Dev Sàrl	5%	Allocation fonctionnelle du fondateur (vesting)
PoSS Stimulus	5%	Soutien aux premières ONG/curators
Pré-vente optionnelle	5%	Investisseurs suisses uniquement, vesting obligatoire
Réserve PoSS	80%	Récompenses PoSS sur plusieurs décennies

→ Modèle inspiré de Tezos et Nym, adapté au cadre légal suisse.

3. Architecture Technique
3.1 Stack finale (Phase 4 validée)

Cosmos SDK 0.46.11

Ethermint 0.22.0 (EVM complet)

CometBFT 0.34.27

Modules natifs : Auth, Bank, Params, Staking, Gov

Modules Noorchain :

x/noorsignal (PoSS v4 complet)

x/feemarket (EVM gas handling)

3.2 Structure de l’application

NoorchainApp avec :

BaseApp

InterfaceRegistry

Codec

Keepers (Bank, Auth, Staking, EVM, FeeMarket, NoorSignal)

ModuleManager

Simulation Manager (optionnel Phase 7)

3.3 Stores

KV Stores pour tous les modules

State model défini dans Phase4_07

BeginBlock/EndBlock intégrés

4. PoSS — Proof of Signal Social

(Version complète et officielle)

4.1 Types de signaux

Micro-dons

Participations

Contenus

CCN (curation communautaire)

4.2 Règles immuables

70% Participant

30% Curateur

Poids des signaux définis par paramètres

Hard-cap 299 millions → pas d’inflation

Halving tous les 8 ans

Limites journalières par participant et par curateur

4.3 Module PoSS

Implémente :

Params PoSS (activables/désactivables)

Keeper complet

Compteurs journaliers

Halving

Génération reward théorique

MsgCreateSignal

Queries

Événements PoSS

État actuel :

✔ Module complet
✔ Paramètres intégrés
✔ BeginBlock OK
✔ Génesis State OK
✔ Testnet OK
❗ Mint désactivé tant que PoSSEnabled = false (Legal Light CH)

5. Gouvernance & Fondation
5.1 Fondation NOORCHAIN (Association suisse)

Forme : association non-profit (CC 60-79)

Rôle : stewardship, supervision, conformité

Pas de PSP, pas de custody, pas de rendements

5% allocation gérée avec multi-sig 3/5

5.2 Multi-sig Committee (3/5)

Composé de :

fondateur

représentant ONG

expert technique

observateur légal

représentant éducatif

Rôle : exécuter les décisions approuvées.

5.3 Curators

3 niveaux : Bronze, Argent, Or
Rôle : valider les signaux PoSS
Récompense : 30% du reward PoSS

6. Legal Light CH — Résumé

NOORCHAIN est :

non-custodial

non-spéculatif

utility token

non-investment

non-PSP

conforme FINMA Light

Aucune fonctionnalité ne peut violer :

“no custody”

“no yield”

“no investment offering”

7. Genesis (Mainnet)
7.1 Les 5 adresses officielles (générées en Phase 7)

Foundation (5%)

Dev Sàrl (5%)

PoSS Stimulus (5%)

Presale (5%)

PoSS Reserve (80%)

7.2 Génesis.json

Contient :

allocations fixes

PoSS params

halving

modules init

version app

chain-id

addresses (bech32)

8. Roadmap Officielle (Phases 1 → 10)

Cadrage ✔

Skeleton Cosmos/EVM ✔

Docs officielles ✔

Implémentation complète ✔

Gouvernance & Legal ✔

Genesis Pack (prochaine étape)

Mainnet 1.0

dApps & Ecosystème

Partenariats & Audits

Interopérabilité & Liquidité (post-mainnet)

9. Sécurité

Cold wallets recommandés

Réserve PoSS protégée

Pas de custody interne

Pas de PSP interne

Audits réguliers

Multi-sig obligatoire pour fonds sensibles

10. État Actuel du Projet (Décembre 2025)

Noyau Cosmos/EVM opérationnel

Module PoSS complet

Testnet fonctionnel

Phase 5 terminée

Ready → Phase 6 (Genesis Pack & Site officiel)
