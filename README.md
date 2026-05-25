# Agent-Kit
A developer CLI for building, publishing, and registering ERC-8004 agents on-chain using IPFS (Pinata) and EVM networks like Celo.

- [Overview](#overview)
- [Architecture](#architecture)
- [Installation](#installation)
- [Commands](#commands)
- [Supported Networks](#supported-networks)
- [Key-Concepts](#key-concepts)
- [Security-Notes](#security-notes)
# Overview
**agent-kit** provides a full lifecycle tool for decentralized agents:
1. Create agent metadata (ERC-8004 format)
2. Publish metadata to IPFS (Pinata)
3. Register agent on-chain (Celo / EVM chains)
4. Give agent Reputation
5. Get Agent Metadata


# Architecture
```
agent-kit CLI
   ├── create   → Generate metadata JSON
   ├── publish  → Upload to IPFS (Pinata)
   └── register → Write to blockchain registry
```
# Installation
Option 1 — Run with Go (development)
```
go run ./cmd/agent-cli --help
```
Run:
```
./agent-kit --help
```
Option 2 — Docker (recommended)
```
docker build -t agent-kit .
```
# Commands
1. CREATE AGENT
Go:
```
go run ./cmd/agent-cli create \
  --name "Weather Agent" \
  --desc "Provides weather forecasting" \
  --image "https://example.com/image.png" \
  --service "api,https://api.example.com,v1" \
  --service "mcp,https://mcp.example.com,v2" \
  --out agent.json
```
Docker:
```
docker run --rm \
  -v $(pwd):/app \
  agent-kit create \
  --name "Weather Agent" \
  --desc "Provides weather forecasting" \
  --image "https://example.com/image.png" \
  --service "api,https://api.example.com,v1" \
  --service "mcp,https://mcp.example.com,v2" \
  --out /app/agent.json
```
Output file:
```
{
  "type": "https://eips.ethereum.org/EIPS/eip-8004#registration-v1",
  "name": "Weather Agent",
  "description": "Provides weather forecasting",
  "image": "https://example.com/image.png",
  "services": [
    {
      "name": "api",
      "endpoint": "https://api.example.com",
      "version": "v1"
    }
  ]
}
```
2. PUBLISH TO IPFS (PINATA)
Upload metadata to IPFS and get a CID + URI.
Go:
```
go run ./cmd/agent-cli publish \
 --file agent.json \  
 --jwt "cd5620032863163c44fa" \
 --name "weather-agent" \
 --json

```
Docker:
```
docker run --rm \
  -v $(pwd):/app \
  agent-kit publish \
  --file /app/agent.json \
  --jwt "cd5620032863163c44fa" \
  --name "weather-agent"
```
Output:
```
{
 "cid": "QmV3dY5GkAiddzXN7k8mYjukWbvv6vyPqcTBdk2mvPYJjr",
 "uri": "ipfs://QmV3dY5GkAiddzXN7k8mYjukWbvv6vyPqcTBdk2mvPYJjr",
 "size": 332,
 "timestamp": "2026-05-19T18:00:34.899Z"
}

```
3. REGISTER ON-CHAIN
Register your IPFS-hosted agent on a blockchain registry.
Go:
```
go run ./cmd/agent-cli register \
 --agent-uri "ipfs://url" \
 --chain "Celo" \
 --network "testnet" \
 --private-key "Your Private key" \
 --json

```
Docker:
```
docker run --rm \
  agent-kit register \
  --agent-uri "ipfs://QmV3dY5GkA..." \
  --chain "Celo" \
  --private-key "Your Private key" \
  --network "testnet"
```
Output:
```
{
  "txHash": "0x1929afbdfb00e8282d7635eec326190f8524ca0c2851f64c282bd1d5ed311116",
  "uri": "ipfs://QmV3dY5GkA...",
  "chain": "Celo",
  "network": "testnet",
  "chainId": 11142220,
  "registry": "0x8004A818BFB912233c491871b3d84c89A494BD9e"
}
```
# Supported Networks
- Celo
- Base
- Monad
- Avax
- Arbitrum
- Ethereum

# Key concepts
**ERC-8004 Agent Metadata**
Each agent contains:
- Name
- Description
- Image
- Services (API, MCP, etc.)
- Type: ERC-8004 registration schema
# Security Notes
- Never commit PRIVATE_KEY
- Never commit PINATA_JWT
# License
MIT
