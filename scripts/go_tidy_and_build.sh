#!/usr/bin/env bash

# NOORCHAIN — Phase 2 helper
# Lance go mod tidy puis go build ./... à la racine du projet.

echo ">>> NOORCHAIN — go mod tidy"
go mod tidy

echo
echo ">>> NOORCHAIN — go build ./..."
go build ./...

echo
echo ">>> Terminé."
