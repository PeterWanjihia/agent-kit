package chain

import (
	"context"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ronexlemon/agent-kit/pkg/chain"
	"github.com/ronexlemon/agent-kit/pkg/contract"
	"github.com/ronexlemon/agent-kit/pkg/erc8004"
)

const (
	NetworkMainnet = "mainnet"
	NetworkTestnet = "testnet"
)

type NetworkClientConfig struct {
	Chain   string
	Network string
}

type ResolvedNetworkClient struct {
	Chain   string
	Network string
	RPCURL  string
	ChainID int64

	IdentityRegistryAddress   string
	ReputationRegistryAddress string

	Eth    *ethclient.Client
	Client *erc8004.ERC8004Client
}

func ResolveNetworkClient(ctx context.Context, cfg NetworkClientConfig) (*ResolvedNetworkClient, error) {
	chainKey, err := normalizeChainKey(cfg.Chain)
	if err != nil {
		return nil, err
	}

	network := normalizeNetwork(cfg.Network)

	rpcURL, err := resolveRPCURL(chainKey, network)
	if err != nil {
		return nil, err
	}

	identity, reputation, err := resolveContractAddresses(chainKey, network)
	if err != nil {
		return nil, err
	}

	eth, err := chain.NewCreateClient(rpcURL)
	if err != nil {
		return nil, err
	}

	chainID, err := eth.ChainID(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to resolve chain id: %w", err)
	}

	client, err := erc8004.NewERC8004Client(eth, identity, reputation)
	if err != nil {
		return nil, fmt.Errorf("failed to create erc8004 client: %w", err)
	}

	return &ResolvedNetworkClient{
		Chain:                     chainKey,
		Network:                   network,
		RPCURL:                    rpcURL,
		ChainID:                   chainID.Int64(),
		IdentityRegistryAddress:   identity,
		ReputationRegistryAddress: reputation,
		Eth:                       eth,
		Client:                    client,
	}, nil
}

func normalizeNetwork(network string) string {
	network = strings.TrimSpace(strings.ToLower(network))
	if network == "" {
		return NetworkTestnet
	}

	return network
}

func normalizeChainKey(chainName string) (string, error) {
	chainName = strings.TrimSpace(strings.ToLower(chainName))
	if chainName == "" {
		return "", fmt.Errorf("chain is required")
	}

	switch chainName {
	case "ethereum", "eth":
		return "Ethereum", nil
	case "base":
		return "Base", nil
	case "arbitrum", "arb":
		return "Arbitrum", nil
	case "avalanche", "avax":
		return "Avalanche", nil
	case "celo":
		return "Celo", nil
	case "monad":
		return "Monad", nil
	default:
		return "", fmt.Errorf("unsupported chain %q", chainName)
	}
}

func resolveRPCURL(chainKey string, network string) (string, error) {
	var rpcConfig string

	switch network {
	case NetworkMainnet:
		cfg, ok := chain.Mainnet_RPCS[chainKey]
		if !ok {
			return "", fmt.Errorf("rpc url not configured for %s mainnet", chainKey)
		}
		rpcConfig = cfg.RPC_URL

	case NetworkTestnet:
		cfg, ok := chain.Testnet_RPCS[chainKey]
		if !ok {
			return "", fmt.Errorf("rpc url not configured for %s testnet", chainKey)
		}
		rpcConfig = cfg.RPC_URL

	default:
		return "", fmt.Errorf("unsupported network %q", network)
	}

	rpcConfig = strings.TrimSpace(rpcConfig)
	if rpcConfig == "" {
		return "", fmt.Errorf("rpc url is empty for %s %s", chainKey, network)
	}

	return rpcConfig, nil
}

func resolveContractAddresses(chainKey string, network string) (string, string, error) {
	var identity string
	var reputation string

	switch network {
	case NetworkMainnet:
		cfg, ok := contract.Mainnet[chainKey]
		if !ok {
			return "", "", fmt.Errorf("contracts not configured for %s mainnet", chainKey)
		}
		identity = cfg.ContractAddress.Identity
		reputation = cfg.ContractAddress.Reputation

	case NetworkTestnet:
		cfg, ok := contract.Testnets[chainKey]
		if !ok {
			return "", "", fmt.Errorf("contracts not configured for %s testnet", chainKey)
		}
		identity = cfg.ContractAddress.Identity
		reputation = cfg.ContractAddress.Reputation

	default:
		return "", "", fmt.Errorf("unsupported network %q", network)
	}

	identity = strings.TrimSpace(identity)
	reputation = strings.TrimSpace(reputation)

	if identity == "" {
		return "", "", fmt.Errorf("identity registry address is empty for %s %s", chainKey, network)
	}

	if reputation == "" {
		return "", "", fmt.Errorf("reputation registry address is empty for %s %s", chainKey, network)
	}

	return identity, reputation, nil
}
