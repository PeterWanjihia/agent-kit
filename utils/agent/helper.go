package agent

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ronexlemon/agent-kit/pkg/types"
	"github.com/spf13/cobra"
)

func ParseServices(values []string) ([]types.AgentService, error) {
	if len(values) == 0 {
		return nil, nil
	}

	var services []types.AgentService

	for _, v := range values {
		v = strings.TrimSpace(v)
		if v == "" {
			continue
		}

		parts := strings.Split(v, ",")

		if len(parts) < 2 {
			return nil, fmt.Errorf(
				"invalid service format: %q (expected name,endpoint[,version])",
				v,
			)
		}

		service := types.AgentService{
			Name:     strings.TrimSpace(parts[0]),
			Endpoint: strings.TrimSpace(parts[1]),
		}

		if service.Name == "" || service.Endpoint == "" {
			return nil, fmt.Errorf("service name and endpoint cannot be empty: %q", v)
		}

		if len(parts) >= 3 {
			service.Version = strings.TrimSpace(parts[2])
		}

		services = append(services, service)
	}

	return services, nil
}



func CollectMissingRequiredInput(input *types.CreateAgentInput, noInteractive bool) error {
	input.Name = strings.TrimSpace(input.Name)
	input.Description = strings.TrimSpace(input.Description)

	if noInteractive {
		return nil
	}

	reader := bufio.NewReader(os.Stdin)

	if input.Name == "" {
		value, err := prompt(reader, "Enter Agent Name: ")
		if err != nil {
			return err
		}
		input.Name = value
	}

	if input.Description == "" {
		value, err := prompt(reader, "Enter Agent Description: ")
		if err != nil {
			return err
		}
		input.Description = value
	}

	return nil
}

func PrintJSON(cmd *cobra.Command, metadata types.AgentMetadata) error {
	encoded, err := json.MarshalIndent(metadata, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to encode metadata: %w", err)
	}

	fmt.Fprintln(cmd.OutOrStdout(), string(encoded))
	return nil
}



func PrintSummary(cmd *cobra.Command, metadata types.AgentMetadata) {
	fmt.Fprintln(cmd.OutOrStdout(), "\n--- Agent Metadata ---")
	fmt.Fprintf(cmd.OutOrStdout(), "Name: %s\n", metadata.Name)
	fmt.Fprintf(cmd.OutOrStdout(), "Description: %s\n", metadata.Description)

	if metadata.Image != "" {
		fmt.Fprintf(cmd.OutOrStdout(), "Image: %s\n", metadata.Image)
	}

	if len(metadata.Services) > 0 {
		fmt.Fprintln(cmd.OutOrStdout(), "Services:")
		for _, service := range metadata.Services {
			fmt.Fprintf(cmd.OutOrStdout(), "  - %s: %s", service.Name, service.Endpoint)

			if service.Version != "" {
				fmt.Fprintf(cmd.OutOrStdout(), " (%s)", service.Version)
			}

			fmt.Fprintln(cmd.OutOrStdout())
		}
	}
}

func prompt(reader *bufio.Reader, label string) (string, error) {
	fmt.Print(label)

	value, err := reader.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("failed to read input: %w", err)
	}

	return strings.TrimSpace(value), nil
}

func WriteMetadataFile(path string, metadata types.AgentMetadata) error {
	cleanPath := filepath.Clean(path)

	info, err := os.Stat(cleanPath)
	if err == nil && info.IsDir() {
		return fmt.Errorf("output path points to a directory: %s", cleanPath)
	}

	encoded, err := json.MarshalIndent(metadata, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to encode metadata: %w", err)
	}

	encoded = append(encoded, '\n')

	if err := os.WriteFile(cleanPath, encoded, 0o644); err != nil {
		return fmt.Errorf("failed to write metadata file: %w", err)
	}

	return nil
}

func PrintRegisterResult(cmd *cobra.Command, result *types.RegisterResult, jsonOutput bool) error {
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


func ValidatePrivateKey(privateKey string) error {
	privateKey = strings.TrimPrefix(strings.TrimSpace(privateKey), "0x")

	if privateKey == "" {
		return fmt.Errorf("private key is required")
	}

	if _, err := crypto.HexToECDSA(privateKey); err != nil {
		return fmt.Errorf("invalid private key: %w", err)
	}

	return nil
}

