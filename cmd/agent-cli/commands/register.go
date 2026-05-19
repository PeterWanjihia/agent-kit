package commands

import (
	"fmt"
	"os"
	"strings"
	"github.com/ronexlemon/agent-kit/pkg/chain"
	network "github.com/ronexlemon/agent-kit/pkg/network"
	"github.com/ronexlemon/agent-kit/pkg/types"
	"github.com/ronexlemon/agent-kit/utils/agent"
	"github.com/spf13/cobra"
)





func NewRegisterCommand() *cobra.Command {
	opts := &types.RegisterOptions{}

	cmd := &cobra.Command{
		Use:          "register",
		Short:        "Register an agent URI on-chain",
		Long:         "Register an already-published ERC-8004 agent metadata URI on the Identity Registry.",
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			result, err := runRegister(cmd, opts)
			if err != nil {
				return err
			}

			return agent.PrintRegisterResult(cmd, result, opts.JsonOutput)
		},
	}

	flags := cmd.Flags()

	flags.StringVar(&opts.AgentURI, "agent-uri", "", "Published agent metadata URI, for example ipfs://...")
	flags.StringVar(&opts.ChainName, "chain", "", "Chain name, for example Celo")
	flags.StringVar(&opts.Network, "network", network.NetworkTestnet, "Network name. Defaults to testnet")
	flags.StringVar(&opts.PrivateKey, "private-key", "", "Private key. Defaults to PRIVATE_KEY environment variable")
	flags.BoolVar(&opts.JsonOutput, "json", false, "Print register result as JSON")

	return cmd
}

func runRegister(cmd *cobra.Command, opts *types.RegisterOptions) (*types.RegisterResult, error) {
	if err := normalizeAndValidate(opts); err != nil {
		return nil, err
	}

	if err := agent.ValidatePrivateKey(opts.PrivateKey); err != nil {
		return nil, err
	}

	resolved, err := network.ResolveNetworkClient(cmd.Context(), network.NetworkClientConfig{
		Chain:   opts.ChainName,
		Network: opts.Network,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to resolve network client: %w", err)
	}

	auth, err := chain.MakeAuth(opts.PrivateKey, resolved.ChainID)
	if err != nil {
		return nil, fmt.Errorf("failed to create transaction auth: %w", err)
	}

	tx, err := resolved.Client.RegisterAgent(auth, opts.AgentURI)
	if err != nil {
		return nil, fmt.Errorf("failed to register agent: %w", err)
	}

	return &types.RegisterResult{
		TxHash:   tx.Hash().Hex(),
		URI:      opts.AgentURI,
		Chain:    resolved.Chain,
		Network:  resolved.Network,
		ChainID:  resolved.ChainID,
		Registry: resolved.IdentityRegistryAddress,
	}, nil
}



func  normalizeAndValidate(opts *types.RegisterOptions) error {
	opts.AgentURI = strings.TrimSpace(opts.AgentURI)
	opts.ChainName = strings.TrimSpace(opts.ChainName)
	opts.Network = strings.TrimSpace(opts.Network)
	opts.PrivateKey = strings.TrimSpace(opts.PrivateKey)

	if opts.PrivateKey == "" {
		opts.PrivateKey = strings.TrimSpace(os.Getenv("PRIVATE_KEY"))
	}

	if opts.AgentURI == "" {
		return fmt.Errorf("agent uri is required")
	}

	if opts.ChainName == "" {
		return fmt.Errorf("chain is required")
	}

	if opts.Network == "" {
		opts.Network = network.NetworkTestnet
	}

	if opts.PrivateKey == "" {
		return fmt.Errorf("private key is required")
	}

	return nil
}


