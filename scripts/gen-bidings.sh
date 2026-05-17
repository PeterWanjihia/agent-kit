#!/usr/bin/env bash

set -e

echo "Generating ERC-8004 Go bindings..."

# Create output dirs
mkdir -p pkg/contract/registry
mkdir -p pkg/contract/reputation

# -----------------------------
# Identity / Registry contract
# -----------------------------
echo "Generating AgentRegistry bindings..."
abigen \
  --abi=pkg/contract/abi/identityRegistry.json \
  --pkg=registry \
  --out=pkg/contract/registry/agent_registry.go

# -----------------------------
# Reputation contract
# -----------------------------
echo " Generating AgentReputation bindings..."
abigen \
  --abi=pkg/contract/abi/reputationRegistry.json \
  --pkg=reputation \
  --out=pkg/contract/reputation/agent_reputation.go

echo " Done generating ERC-8004 bindings"