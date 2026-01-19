#!/usr/bin/env bash

set -e

echo "🔧 Initializing NOORCHAIN local testnet (filesystem only)..."

CHAIN_DIR="./data-testnet"

# On repart propre
rm -rf "$CHAIN_DIR"

# On crée la structure minimale attendue par le node
mkdir -p "$CHAIN_DIR/config"

# On copie notre genesis préparé
cp testnet/genesis.json "$CHAIN_DIR/config/genesis.json"

echo "✅ Testnet directory initialized in $CHAIN_DIR"
echo "  - genesis.json copied to $CHAIN_DIR/config/genesis.json"
echo
echo "👉 Prochaine étape (plus tard) :"
echo "   ./noord start --home $CHAIN_DIR"
echo "   (quand la commande start sera complètement câblée côté CLI/serveur)"
