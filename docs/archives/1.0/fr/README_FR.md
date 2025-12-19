NOORCHAIN Core — Version Française

Ce document présente la version française du fichier README principal du dépôt NOORCHAIN Core.

La version officielle destinée aux développeurs et auditeurs externes est en anglais et se trouve à la racine du dépôt, dans le fichier nommé README.md.

Présentation

NOORCHAIN Core est l’implémentation de référence du réseau NOORCHAIN.

Le projet repose sur Cosmos SDK et Ethermint, et fournit les composants essentiels d’un nœud blockchain : gestion des comptes, soldes, staking, gouvernance, exécution EVM, gestion du marché des frais (EIP-1559) et module PoSS (Proof of Signal Social).

Compilation

La compilation du nœud s’effectue en utilisant Go.
L’entrée du programme se trouve dans le répertoire cmd/noord.
Go version 1.22 ou supérieure est nécessaire.

Exécution d’un nœud local

Une fois le binaire compilé, il peut être lancé directement.
Le nœud utilise par défaut un répertoire personnel pour stocker son état et ses données locales.

Structure du dépôt

app
Initialisation de l’application, gestion des keepers et des modules.

cmd/noord
Entrée principale du nœud NOORCHAIN.

config
Fichiers d’aide à la configuration.

scripts
Scripts opérationnels utilisés pour les tests réseau et diverses tâches de maintenance.

x/auth
Module d’authentification et gestion des identités.

x/bank
Module de gestion des soldes et des transferts.

x/staking
Module de staking et fonctionnalités associées.

x/gov
Module de gouvernance et gestion des paramètres en chaîne.

x/evm
Couche de compatibilité EVM fournie par Ethermint.

x/feemarket
Logique de marché des frais basée sur EIP-1559.

x/noorsignal
Module PoSS (Proof of Signal Social).

docs
Documentation technique comprenant les spécifications du protocole, le genesis, la gouvernance et le PoSS.

Documentation

La documentation technique complète est accessible dans le dossier docs.
Les versions françaises sont regroupées dans docs/fr.

Licence

À définir.
