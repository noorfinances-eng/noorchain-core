# NOORCHAIN Testnet 1.0 — Plan pour les adresses réelles

Ce document explique comment, le moment venu, générer les **vraies adresses**
NOORCHAIN (bech32 `noor1...`) et les synchroniser dans tous les fichiers
nécessaires au testnet.

⚠️ **À faire plus tard**, quand un environnement local Go + `noord` sera prêt.

---

## 1. Rôles et adresses nécessaires

Nous avons besoin de **5 adresses** principales pour le genesis économique :

1. **NOOR Foundation (association light)**  
   - Rôle : vitrine publique, transparence, gouvernance "light"  
   - Pourcentage : **5 %** du supply

2. **Founder Wallet (Walid perso)**  
   - Rôle : 5 % fondateur (temporaire, avant création de la Sàrl)  
   - Pourcentage : **5 %**

3. **PoSS Stimulus Pool**  
   - Rôle : stimuler les premiers signaux / partenaires / curateurs  
   - Pourcentage : **5 %**

4. **Optional Pre-Sale Pool**  
   - Rôle : éventuelle pré-vente contrôlée (investisseurs suisses, vesting)  
   - Pourcentage : **5 %**

5. **PoSS Reserve (80 %)**  
   - Rôle : réserve principale pour le minage social PoSS  
   - Pourcentage : **80 %**

Toutes ces adresses seront au format bech32 :

```text
noor1xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
