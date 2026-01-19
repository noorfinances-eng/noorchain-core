FONDATION NOORCHAIN
Gouvernance du Genesis
Phase 5 — C3
Version 1.1
Dernière mise à jour : 2025-XX-XX
1. Objet de la gouvernance du Genesis

Ce document définit la configuration de gouvernance initiale incluse dans le
bloc Genesis du Mainnet (“genesis.json”).

Il garantit :

une allocation transparente

une répartition économique claire

une conformité totale au cadre Swiss Legal Light

une séparation nette entre parties publiques, privées et protocolaires

une initialisation correcte des paramètres PoSS

l’assignation cohérente des adresses institutionnelles

l’intangibilité des principes fondamentaux

Cette gouvernance du Genesis constitue la base de la crédibilité et de la sécurité juridique de NOORCHAIN.

2. Principes immuables du Genesis

Les principes suivants sont intégrés dans le Genesis et ne peuvent jamais être modifiés :

Offre fixe
Cap total = 299 792 458 NUR.

Halving tous les 8 ans
Cycle de réduction fixe et non modifiable.

Aucune inflation
Aucun mint discrétionnaire n’est autorisé.

Modèle économique 5 / 5 / 5 / 5 / 80

5 % Fondation

5 % Noor Dev Sàrl

5 % Stimulus PoSS

5 % Pré-vente optionnelle

80 % Réserve PoSS mintable

Répartition PoSS immuable

70 % pour le participant

30 % pour le curateur

Aucune garde de fonds utilisateurs
La Fondation n’a pas le droit de détenir des actifs d’utilisateurs.

Non-investment token
Aucune promesse de rendement ; aucun produit financier.

Conformité Legal Light CH
Entièrement non-custodial, non-financier, non-spéculatif.

3. Adresses du Genesis (à remplacer en Phase 7)

Les adresses suivantes sont des placeholders :

3.1 Adresse de la Fondation (5 %)

noor1foundationxxxxxxxxxxxxxxxxxxxxx

3.2 Adresse Noor Dev Sàrl (5 %)

noor1devsarlxxxxxxxxxxxxxxxxxxxxxxxxxxx

3.3 Réserve Stimulus PoSS (5 %)

noor1stimulusxxxxxxxxxxxxxxxxxxxxxxxx

3.4 Réserve Pré-vente (5 %)

noor1presalexxxxxxxxxxxxxxxxxxxxxxxxxx

3.5 Réserve PoSS Mintable (80 %)

noor1possreservexxxxxxxxxxxxxxxxxxxxxx

Elles seront remplacées par les 5 vraies adresses bech32 lors de la Phase 7.

4. Allocations du Genesis (intégrées dans genesis.json)
Catégorie	Pourcentage	Montant (NUR)
Fondation	5%	14 989 623
Noor Dev Sàrl	5%	14 989 623
Stimulus PoSS	5%	14 989 623
Pré-vente	5%	14 989 623
Réserve PoSS	80%	239 833 984
Total	100%	299 792 458

Ces allocations sont définitives et non modifiables.

5. Pouvoirs de gouvernance au Genesis
5.1 Fondation

Supervise les missions, documents, transparence

Gère les Curateurs

Publie les rapports publics

5.2 Comité Multi-sig (3/5)

Exécute les décisions de la Fondation

Contrôle l’adresse Fondation

Peut modifier les paramètres PoSS non structurels

Peut suspendre temporairement PoSS en cas d’urgence

5.3 Noor Dev Sàrl

Pas de pouvoir de gouvernance au genesis

Détient l’allocation fonctionnelle (5 %)

Peut proposer des améliorations

5.4 Curateurs

Aucun pouvoir de gouvernance formel

Rôle purement consultatif

6. Paramètres du Genesis (PoSS et Protocole)
6.1 Paramètres PoSS initiaux

PoSSEnabled : false
BaseReward : 1 unur
WeightMicroDonation : 2
WeightParticipation : 1
WeightContent : 3
WeightCCN : 4
MaxSignalsPerDay : 10
MaxSignalsPerCuratorPerDay : 20
MaxRewardPerDay : 100 unur
HalvingPeriodYears : 8

6.2 Contraintes économiques

la réserve PoSS (80%) est la seule source de minting

aucune extension de pool n’est autorisée

aucune inflation possible

7. Limites de gouvernance (verrouillées)

Interdictions absolues :

modifier l’offre totale

modifier le split PoSS 70/30

créer des tokens supplémentaires

changer les pourcentages 5/5/5/5/80

supprimer ou réduire les allocations Fondation ou Sàrl

contourner le cadre Legal Light

8. Chemin d’évolution post-genesis

Les évolutions possibles incluent :

activation officielle de PoSS

ajout de Curateurs

mise à jour des paramètres PoSS

ajout de modules (Hub, Pay, Studio)

extension du multi-sig

votes consultatifs des Curateurs

9. Documents requis au lancement

À publier ou intégrer dans le Genesis Pack :

Charte de gouvernance

Charte Multi-sig

Legal Light PDF

Allocation Genesis PDF

Statuts de la Fondation

10. Adoption

Ce document est adopté par le
Conseil de la Fondation NOORCHAIN
et constitue un élément du Genesis Pack 1.1.

Signatures :

Version 1.1
Préparé pour la Fondation NOORCHAIN — Phase Gouvernance
