# NOORCHAIN — EVM & MetaMask Setup (Testnet)

Ce document explique comment, **le jour où le testnet sera en ligne**,
ajouter NOORCHAIN (EVM) dans MetaMask et tester des smart contracts.

Pour l’instant, c’est un guide théorique / prêt à l’emploi.

---

## 1. Paramètres généraux à prévoir pour MetaMask

Quand le RPC EVM sera disponible, il faudra créer un réseau personnalisé
dans MetaMask avec des champs du type :

- **Network name** : `NOORCHAIN Testnet`
- **New RPC URL** : `https://rpc-testnet.noorchain.org` *(exemple, à définir)*
- **Chain ID** : `TBD_EVM_CHAIN_ID`
- **Currency symbol** : `NUR`
- **Block explorer URL** : `https://explorer-testnet.noorchain.org` *(optionnel au début)*

> 📌 Remarque  
> Le `Chain ID` EVM sera défini dans le code (ex: via une constante `EvmChainID`)
> et devra être cohérent avec la config du module `evm` dans le genesis.

---

## 2. Lien entre Cosmos `ChainID` et EVM `chainId`

NOORCHAIN a deux notions :

- **ChainID Cosmos** : ex. `noorchain-testnet-1`
- **chainId EVM** : ex. `1001`, `2025`, etc. (entier utilisé par MetaMask)

Les règles :

- `ChainID` Cosmos est utilisé par le node (`noord`) et les modules SDK
- `chainId` EVM est utilisé par :
  - MetaMask
  - les transactions Ethereum (EIP-155)
  - les librairies Web3 / ethers.js

Ils doivent être définis de manière cohérente dans :

- la config EVM (genesis → `evm.params.chain_config`)
- la doc MetaMask (ce fichier)
- les dApps qui se connecteront à NOORCHAIN

---

## 3. Exemple de configuration MetaMask (future)

Quand le testnet sera prêt, une config possible serait :

- **Network name** : `NOORCHAIN Testnet`
- **New RPC URL** : `https://rpc-testnet.noorchain.org`
- **Chain ID** : `1025` *(exemple, à figer plus tard)*
- **Currency symbol** : `NUR`
- **Block explorer URL** : *(à ajouter plus tard, si un explorer est déployé)*

Pour l’instant, ces valeurs sont **PLACEHOLDER**.  
Elles seront fixées au moment :

1. où le genesis EVM sera finalisé,
2. où le node RPC public sera disponible,
3. où le `chainId` sera figé officiellement.

---

## 4. Flux de test EVM typique (jour J)

Une fois le testnet opérationnel et MetaMask configuré :

1. **Ajouter NOORCHAIN Testnet** dans MetaMask avec les paramètres officiels.
2. **Récupérer des NUR de test** (via faucet interne ou crédits initiaux).
3. **Déployer un premier contrat simple** (ex. via Remix) :

   ```solidity
   // Exemple ultra-simple pour tester NOORCHAIN EVM
   // (à déployer sur le testnet NOORCHAIN)
   contract HelloNoor {
       function hi() public pure returns (string memory) {
           return "Hello from NOORCHAIN EVM";
       }
   }
