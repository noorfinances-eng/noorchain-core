# NOORCHAIN — Mainnet 1.0 Checklist (Officielle)

Cette checklist doit être suivie **dans l’ordre exact**.

---

# 1. Génération des 5 adresses officielles

- [ ] Fondation NOOR
- [ ] Dev Wallet Fondateur (5 %)
- [ ] PoSS Stimulus
- [ ] Pré-vente optionnelle
- [ ] PoSS Reserve (80 %)

Toutes doivent être dans le format :


---

# 2. Mise à jour du code

- [ ] Modifier `x/noorsignal/types/addresses.go`
- [ ] Recompiler l’application
- [ ] Vérifier que `ApplyEconomicGenesis` corrige bien les balances

---

# 3. Génération du genesis mainnet

- [ ] Copier `mainnet/genesis.template.json` → `mainnet/genesis.json`
- [ ] Remplacer les adresses par les vraies
- [ ] Vérifier balances et supply
- [ ] Exécuter :


Résultat attendu : **OK**

---

# 4. Sécurité et validation

- [ ] Vérification manuelle ligne par ligne
- [ ] Vérification supply NUR = 299 792 458 × 10⁶ unur
- [ ] Vérification 5/5/5/5/80
- [ ] Signature numérique (optionnel)
- [ ] Stockage du genesis sur plusieurs devices

---

# 5. Publication officielle

- [ ] Publier `genesis.json` sur GitHub
- [ ] Publier la version hash SHA256
- [ ] Préparer la commande de démarrage :


---

# 6. Lancement du Mainnet

- [ ] Premier bootstrap du réseau
- [ ] Vérification RPC / gRPC
- [ ] Vérification Query :
- [ ] Vérification Bank :


---

# 7. Ouverture publique

- [ ] Publication page officielle
- [ ] Publication du hash du genesis
- [ ] Publication des adresses officielles
- [ ] Activation progressive du PoSS réel (BankKeeper)

---

Cette checklist doit être suivie strictement pour éviter toute erreur économique.

