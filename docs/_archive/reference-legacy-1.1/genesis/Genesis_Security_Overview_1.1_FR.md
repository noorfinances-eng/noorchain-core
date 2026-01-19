NOORCHAIN — Synthèse Sécurité du Genesis

Version 1.1
Dernière mise à jour : 2025-XX-XX

1. Objet du document

Ce document définit toutes les garanties de sécurité, restrictions et invariants appliqués dans le bloc genesis de NOORCHAIN.

Il garantit :

sécurité juridique et économique

immuabilité des règles fondamentales

initialisation correcte de PoSS

protection contre une gouvernance malveillante

résilience long terme du mainnet

Il fait partie du Genesis Pack.

2. Principes de sécurité au Genesis

NOORCHAIN applique cinq principes fondamentaux :

Immutabilité de la supply

Aucun mint discrétionnaire

Séparation stricte des pouvoirs

Aucune garde de fonds utilisateurs

Conformité Legal Light CH

Ces principes sont irréversibles.

3. Garantie de supply fixe

Le genesis garantit :

Supply totale = 299 792 458 NUR

Aucune inflation sauf via la réserve PoSS

Réserve PoSS plafonnée à 80 %

Allocations 5/5/5/5/80 immuables

Aucune proposition governance ne peut modifier :

la supply totale

la distribution d’origine

la taille de la réserve PoSS

4. Sécurité de la gouvernance

Structure sécurisée :

Conseil de Fondation

Pouvoirs administratifs.

Comité Multi-sig (3/5)

Exécute les décisions mais ne peut pas créer de supply.

Aucun acteur ne peut :

minter

modifier PoSS structurellement

changer les allocations genesis

violer Legal Light

Toute proposition illégale est nulle automatiquement.

5. Sécurité PoSS

PoSS est désactivé au genesis par sécurité.

Garanties :

Aucun mint avant activation

Mint uniquement depuis la réserve PoSS

Ratio 70/30 immuable

Halving immuable

Limites journalières configurables

Aucun oracle externe

6. Sécurité EVM

Le genesis EVM inclut :

Base Fee désactivée

Aucune smart-contract préchargée

StateDB vide

Compatibilité complète Ethereum

Protection anti-replay via ChainID

7. Sécurité staking

Paramètres protégés :

Self-delegation minimale

Plafonds de commission

75 validateurs max

Déliaison 21 jours

8. Conformité juridique

NOORCHAIN respecte Legal Light CH :

aucun custody

aucune garantie de rendement

pas de fiat gateway interne

PoSS = récompense sociale, pas produit financier

9. Réduction de la surface d’attaque

Désactivés au lancement :

mint PoSS auto

smart-contracts pré-déployés

mécanismes de type trésor

gouvernance illimitée

expansion rapide du set de validateurs

10. Résumé

Le genesis de NOORCHAIN garantit :

supply fixe

économie immuable

PoSS sécurisé

staking solide

EVM sécurisé

gouvernance encadrée

conformité suisse

Version 1.1 adoptée pour le Genesis Pack.
