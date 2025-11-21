# NOORCHAIN — Mainnet Preparation Procedure

Cette procédure décrit le passage précis Testnet → Mainnet.

---

# 1. Prérequis

- Code final compilable
- Adresses officielles générées
- Génération du dossier mainnet/
- Sécurité : cold wallet pour la réserve PoSS

---

# 2. Mise à jour du code

Modifier le fichier :
x/noorsignal/types/addresses.go

En remplaçant les placeholders par les 5 vraies adresses :

- Foundation
- Dev Wallet
- Stimulus
- Presale
- PoSS Reserve

---

# 3. Reconstruction de l’application
make build

Vérifier que la chaîne compile **sans aucun warning critique**.

---

# 4. Génération du Genesis Mainnet

Copier :
mainnet/genesis.template.json → mainnet/genesis.json

Remplacer chaque adresse `<...>` par l’adresse réelle.

---

# 5. Validation du Genesis

noord validate-genesis

Doit renvoyer **OK**.

Si erreur → corriger immédiatement.

---

# 6. Publication officielle

- Commit + tag de la version “Mainnet 1.0”
- Upload du genesis sur GitHub
- Publication du hash SHA256
- Communication officielle auprès des partenaires (ONG, écoles, Curateurs)

---

# 7. Lancement du réseau

Commande :

noord start --home ~/.noorchain

Contrôles post-lancement :

- RPC OK
- gRPC OK
- Query PoSS OK
- Supply OK
- Balances OK

---

# 8. Activation progressive du PoSS réel

Une fois le réseau stable :

1. Activer BankKeeper dans ValidateSignal  
2. Lancer PoSS Reward Distribution en conditions réelles  
3. Monitorer la réserve PoSS

---

# 9. Pérennité du réseau

- backups  
- sécurité cold wallet  
- rotation des clés  
- observabilité RPC/gRPC  




