package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/ronexlemon/agent-kit/pkg/client"
	"github.com/ronexlemon/agent-kit/pkg/types"
	"github.com/spf13/cobra"
)

func printSummary(cmd *cobra.Command, metadata types.AgentMetadata) {
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

func writeMetadataFile(path string, metadata types.AgentMetadata) error {
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

func printJSON(cmd *cobra.Command, metadata types.AgentMetadata) error {
	encoded, err := json.MarshalIndent(metadata, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to encode metadata: %w", err)
	}

	fmt.Fprintln(cmd.OutOrStdout(), string(encoded))
	return nil
}

func newCreateCommand() *cobra.Command {
	var input types.CreateAgentInput
	var outPath string
	var jsonOutput bool
	var noInteractive bool
	var services []string

	createCmd := &cobra.Command{
		Use:   "create",
		Short: "Create local metadata for a new ERC-8004 agent",
		Long:  "Create local agent metadata that can later be published to an ERC-8004 Identity Registry.",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := collectMissingRequiredInput(&input, noInteractive); err != nil {
				return err
			}
			parsedServices, err := parseServices(services)
			if err != nil {
				return err
			}

			input.Services = parsedServices

			metadata, err := client.CreateAgentMetadata(input)
			if err != nil {
				return err
			}

			if outPath != "" {
				if err := writeMetadataFile(outPath, metadata); err != nil {
					return err
				}

				if !jsonOutput {
					fmt.Fprintf(cmd.OutOrStdout(), "Agent metadata created successfully.\nOutput: %s\n", outPath)
					return nil
				}
			}

			if jsonOutput {
				return printJSON(cmd, metadata)
			}

			printSummary(cmd, metadata)
			return nil
		},
	}

	createCmd.Flags().StringVarP(&input.Name, "name", "n", "", "Name of the agent")
	createCmd.Flags().StringVarP(&input.Description, "desc", "d", "", "Description of the agent")
	createCmd.Flags().StringVar(&input.Image, "image", "", "Image URI for the agent")
	createCmd.Flags().StringArrayVar(
		&services,
		"service",
		nil,
		"Service format: name,endpoint[,version]",
	)
	createCmd.Flags().StringVarP(&outPath, "out", "o", "", "Write agent metadata JSON to a file")
	createCmd.Flags().BoolVar(&jsonOutput, "json", false, "Print agent metadata as JSON")
	createCmd.Flags().BoolVar(&noInteractive, "no-interactive", false, "Fail instead of prompting when required fields are missing")

	return createCmd
}

func collectMissingRequiredInput(input *types.CreateAgentInput, noInteractive bool) error {
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

func parseServices(values []string) ([]types.AgentService, error) {
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
		fmt.Println("LEn", len(parts))

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
