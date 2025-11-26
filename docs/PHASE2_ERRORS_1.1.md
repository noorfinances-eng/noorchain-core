# NOORCHAIN — Phase 2 Errors (1.1)

> Fichier dédié au suivi des erreurs `go build ./...` et `go mod tidy`  
> pendant la Phase 2 (Infrastructure & Outillage Testnet).

---

## 1. Commandes de test

- `go mod tidy`
- `go build ./...`

---

## 2. Journal des erreurs

> Format recommandé :
> - Date
> - Commande
> - Extrait de l’erreur
> - Fichier / package concerné
> - Commentaire éventuel

---

### 2.1 Exemple d’entrée

- **Date :** 2025-11-26  
- **Commande :** `go build ./...`  
- **Erreur :**
  - `undefined: something`  
- **Fichier :** `app/xxx.go`  
- **Commentaire :** à corriger dans la prochaine étape.

---

## 3. Rappel

- Ce fichier ne contient **aucun code**.
- Il sert uniquement à documenter les erreurs Phase 2.
- PoSS et la Phase 4 ne doivent pas apparaître ici.
