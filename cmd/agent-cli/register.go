package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ronexlemon/agent-kit/pkg/chain"
	network "github.com/ronexlemon/agent-kit/pkg/network"
	"github.com/spf13/cobra"
)

type registerResult struct {
	TxHash   string `json:"txHash"`
	URI      string `json:"uri"`
	Chain    string `json:"chain"`
	Network  string `json:"network"`
	ChainID  int64  `json:"chainId"`
	Registry string `json:"registry"`
}

type registerOptions struct {
	agentURI   string
	chainName  string
	network    string
	privateKey string
	jsonOutput bool
}

func newRegisterCommand() *cobra.Command {
	opts := &registerOptions{}

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

			return printRegisterResult(cmd, result, opts.jsonOutput)
		},
	}

	flags := cmd.Flags()

	flags.StringVar(&opts.agentURI, "agent-uri", "", "Published agent metadata URI, for example ipfs://...")
	flags.StringVar(&opts.chainName, "chain", "", "Chain name, for example Celo")
	flags.StringVar(&opts.network, "network", network.NetworkTestnet, "Network name. Defaults to testnet")
	flags.StringVar(&opts.privateKey, "private-key", "", "Private key. Defaults to PRIVATE_KEY environment variable")
	flags.BoolVar(&opts.jsonOutput, "json", false, "Print register result as JSON")

	return cmd
}

func runRegister(cmd *cobra.Command, opts *registerOptions) (*registerResult, error) {
	if err := opts.normalizeAndValidate(); err != nil {
		return nil, err
	}

	if err := validatePrivateKey(opts.privateKey); err != nil {
		return nil, err
	}

	resolved, err := network.ResolveNetworkClient(cmd.Context(), network.NetworkClientConfig{
		Chain:   opts.chainName,
		Network: opts.network,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to resolve network client: %w", err)
	}

	auth, err := chain.MakeAuth(opts.privateKey, resolved.ChainID)
	if err != nil {
		return nil, fmt.Errorf("failed to create transaction auth: %w", err)
	}

	tx, err := resolved.Client.RegisterAgent(auth, opts.agentURI)
	if err != nil {
		return nil, fmt.Errorf("failed to register agent: %w", err)
	}

	return &registerResult{
		TxHash:   tx.Hash().Hex(),
		URI:      opts.agentURI,
		Chain:    resolved.Chain,
		Network:  resolved.Network,
		ChainID:  resolved.ChainID,
		Registry: resolved.IdentityRegistryAddress,
	}, nil
}

func validatePrivateKey(privateKey string) error {
	privateKey = strings.TrimPrefix(strings.TrimSpace(privateKey), "0x")

	if privateKey == "" {
		return fmt.Errorf("private key is required")
	}

	if _, err := crypto.HexToECDSA(privateKey); err != nil {
		return fmt.Errorf("invalid private key: %w", err)
	}

	return nil
}

func (opts *registerOptions) normalizeAndValidate() error {
	opts.agentURI = strings.TrimSpace(opts.agentURI)
	opts.chainName = strings.TrimSpace(opts.chainName)
	opts.network = strings.TrimSpace(opts.network)
	opts.privateKey = strings.TrimSpace(opts.privateKey)

	if opts.privateKey == "" {
		opts.privateKey = strings.TrimSpace(os.Getenv("PRIVATE_KEY"))
	}

	if opts.agentURI == "" {
		return fmt.Errorf("agent uri is required")
	}

	if opts.chainName == "" {
		return fmt.Errorf("chain is required")
	}

	if opts.network == "" {
		opts.network = network.NetworkTestnet
	}

	if opts.privateKey == "" {
		return fmt.Errorf("private key is required")
	}

	return nil
}

func printRegisterResult(cmd *cobra.Command, result *registerResult, jsonOutput bool) error {
	out := cmd.OutOrStdout()

	if jsonOutput {
		encoded, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			return fmt.Errorf("failed to encode register result: %w", err)
		}

		fmt.Fprintln(out, string(encoded))
		return nil
	}

	fmt.Fprintln(out, "Agent registration submitted.")
	fmt.Fprintf(out, "Tx: %s\n", result.TxHash)
	fmt.Fprintf(out, "URI: %s\n", result.URI)
	fmt.Fprintf(out, "Chain: %s\n", result.Chain)
	fmt.Fprintf(out, "Network: %s\n", result.Network)
	fmt.Fprintf(out, "Chain ID: %d\n", result.ChainID)
	fmt.Fprintf(out, "Registry: %s\n", result.Registry)

	return nil
}
