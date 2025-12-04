NOORCHAIN — Paramètres du Genesis

Version 1.1
Dernière mise à jour : 2025-XX-XX

1. Objet du document

Ce document définit l’ensemble des paramètres du protocole intégrés dans le
Genesis du Mainnet.

Il garantit l’alignement entre :

la configuration du protocole

les paramètres initiaux PoSS

les règles de gouvernance

la conformité Legal Light CH

le modèle économique 5/5/5/5/80

Ces paramètres font partie du Genesis Pack.

2. Paramètres centraux de la chaîne
2.1 Identité

Nom de chaîne : noorchain
Chain ID : noorchain-1

2.2 Unité monétaire

Dénomination de base : unur
Dénomination affichée : NUR
Décimales : 18

2.3 Temps de bloc

Temps cible : 5 secondes
Gaz maximal par bloc : 40 000 000
Ajustement gaz EVM : activé

2.4 Gouvernance

Dépôt minimum : 10 000 NUR
Période de vote : 5 jours
Quorum : 33 %
Seuil d’adoption : 50 % + 1
Seuil de veto : 33 %

3. Paramètres PoSS (Genesis)

Ces paramètres initialisent le module PoSS.
Ils sont modifiables par gouvernance sauf les règles structurelles.

Paramètre	Valeur
PoSSEnabled	false
BaseReward	1 unur
WeightMicroDonation	2
WeightParticipation	1
WeightContent	3
WeightCCN	4
MaxSignalsPerDay	10
MaxSignalsPerCuratorPerDay	20
MaxRewardPerDay	100 unur
HalvingPeriodYears	8
PoSSReserveDenom	unur

Invariants structurels (immuables) :
– Répartition 70 % participant / 30 % curateur
– Aucune inflation au-delà de la réserve PoSS
– Supply fixe pour toujours

4. Paramètres économiques

Supply totale : 299 792 458 NUR

Selon le modèle 5 / 5 / 5 / 5 / 80 :

Catégorie	Pourcentage	Notes
Fondation	5%	Multi-sig 3/5
Noor Dev Sàrl	5%	Allocation fonctionnelle
Stimulus PoSS	5%	Premiers partenaires
Pré-vente	5%	Vesting obligatoire
Réserve PoSS	80%	Source unique de minting PoSS

Ces valeurs sont intégrées dans le genesis.json.

5. Limites de gouvernance (verrouillées)

Ne peuvent jamais être modifiées :

Supply totale

Ratio PoSS 70/30

Halving 8 ans

Pourcentages 5/5/5/5/80

Taille de la réserve

Contraintes Legal Light

Aucun mint discrétionnaire

6. Paramètres EVM au Genesis

Base Fee : 0
Prix du gaz minimum : 0
Config EVM : compatible Shanghai (Ethermint)
StateDB : initialisée vide

7. Paramètres de staking

Temps de déliaison : 21 jours
Max validateurs : 75
Commission max : 100 %
Variation max : 1 % / jour
Self-delegation minimale : 1 NUR

8. Résumé

Ces paramètres définissent la forme initiale de NOORCHAIN et garantissent :

un comportement déterministe

une conformité Legal Light

une compatibilité EVM stable

une chaîne prête pour PoSS

une gouvernance claire et encadrée

Version 1.1, valable pour le lancement du mainnet.
