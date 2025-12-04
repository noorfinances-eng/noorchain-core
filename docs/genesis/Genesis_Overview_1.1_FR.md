NOORCHAIN — Vue d’ensemble du Genesis

Version 1.1 — Document Officiel
Dernière mise à jour : 2025-XX-XX
Langue : FR

1. Résumé

Ce document présente la vue d’ensemble officielle du Genesis de NOORCHAIN 1.0.
Il définit l’offre totale immuable, le modèle économique, les règles de gouvernance, les fondations du PoSS, ainsi que le rôle des cinq portefeuilles fondateurs.

Il fait partie du Genesis Pack 1.1, destiné aux développeurs, auditeurs, partenaires, institutions et investisseurs.

2. Philosophie du Genesis

Le Genesis de NOORCHAIN suit trois principes fondamentaux :

1. Offre immuable
L’offre du token est fixée pour toujours à 299 792 458 NUR (référence symbolique à la vitesse de la lumière).

2. Modèle éthique suisse
La chaîne fonctionne sous le cadre Legal Light suisse : transparence, non-spéculation, absence de garde de fonds, absence de rendement garanti.

3. Consensus centré sur l’humain
Le mécanisme PoSS (Proof of Signal Social) utilise une réserve pré-allouée pour récompenser des actions sociales validées.

3. Offre Totale (Immuable)

Offre totale : 299 792 458 NUR
Décimales : 18
Inflation : 0 % de manière permanente
Futurs mint : interdits
Modification via gouvernance : impossible

L’offre totale ne peut jamais être modifiée durant la vie de la blockchain.

4. Allocation Économique (5 / 5 / 5 / 5 / 80)

5 % — Fondation NOOR (gouvernance publique, transparence, opérations)

5 % — Noor Dev Sàrl (entité de développement : 5 % liquides + 10 % vesting hors-genesis)

5 % — Réserve PoSS Stimulus (incitations au lancement)

5 % — Pré-vente optionnelle (avec vesting, conforme aux règles suisses)

80 % — Réserve PoSS (minage social, halving tous les 8 ans)

Ce modèle est fixe et ne peut pas être modifié après le Genesis.

5. Règles Structurelles PoSS Dans le Genesis

Les paramètres suivants sont structurels et immutables :

Distribution des récompenses : 70 % participant / 30 % curateur

Cycle de halving : tous les 8 ans

Source des récompenses : uniquement la Réserve PoSS

Inflation : interdite

Limites journalières actives

Tiers de curateurs (Bronze / Argent / Or)

Poids des signaux (Micro-donation, Participation, Contenu, CCN)

Ces règles sont permanentes et ne peuvent pas être modifiées par la gouvernance.

6. Fondations de Gouvernance (Immuables)

Le Genesis impose les contraintes suivantes :

Multi-sig Fondation : seuil 3/5

Aucune garde de fonds par la Fondation ou la Sàrl

Aucun PSP interne à la blockchain

Offre et allocations : impossibles à modifier

Règles structurelles PoSS : immuables

Seuls les paramètres PoSS (pas la structure) peuvent évoluer via gouvernance

Ce cadre garantit conformité et stabilité à long terme.

7. Portefeuilles du Genesis (Placeholders Jusqu’à la Phase 7)

Cinq portefeuilles seront intégrés durant la Phase 7 (Mainnet) :

Portefeuille Fondation NOOR

Portefeuille Noor Dev Sàrl

Portefeuille PoSS Stimulus

Portefeuille Pré-vente optionnelle

Portefeuille Réserve PoSS

Ils seront insérés dans :

genesis.json (testnet et mainnet)

x/noorsignal/types/addresses.go

Documents de gouvernance et légaux

Genesis Pack PDF

8. Modèle de Sécurité du Genesis

Le Genesis garantit :

la création déterministe du set de validateurs

un state root déterministe

PoSS désactivé initialement sur mainnet (sécurité prioritaire)

aucun risque d’inflation

aucune dépendance externe pour l’offre ou les règles économiques

Il est conçu pour être conforme aux exigences Legal Light et durable.

9. Rôle de Ce Document

Cette vue d’ensemble sert de :

référence générale du Genesis

fondation économique et de gouvernance

lien entre documentation légale et genesis.json

base du Genesis Pack PDF (Phase 6)

10. Versioning

Genesis Overview — Version 1.1
Modifications soumises à gouvernance.
Éléments structurels : non modifiables.
