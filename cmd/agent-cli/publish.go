package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/ronexlemon/agent-kit/pkg/storage"
	"github.com/ronexlemon/agent-kit/pkg/types"
	"github.com/spf13/cobra"
)

func newPublishCommand() *cobra.Command {
	var filePath string
	var jwt string
	var name string
	var jsonOutput bool

	cmd := &cobra.Command{
		Use:          "publish",
		Short:        "Publish agent metadata JSON to Pinata",
		Long:         "Publish a local ERC-8004 agent metadata JSON file to Pinata and return an ipfs:// URI.",
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			filePath = strings.TrimSpace(filePath)
			if filePath == "" {
				return fmt.Errorf("metadata file is required")
			}

			cleanPath := filepath.Clean(filePath)

			raw, err := os.ReadFile(cleanPath)
			if err != nil {
				return fmt.Errorf("failed to read metadata file: %w", err)
			}

			var metadata types.AgentMetadata
			if err := json.Unmarshal(raw, &metadata); err != nil {
				return fmt.Errorf("failed to decode agent metadata json: %w", err)
			}

			jwt = strings.TrimSpace(jwt)
			if jwt == "" {
				jwt = strings.TrimSpace(os.Getenv("PINATA_JWT"))
			}

			pinata, err := storage.NewPinataClient(jwt)
			if err != nil {
				return err
			}

			if strings.TrimSpace(name) == "" {
				name = filepath.Base(cleanPath)
			}

			result, err := pinata.PublishJSON(context.Background(), storage.PublishJSONInput{
				Name: name,
				Data: metadata,
			})
			if err != nil {
				return err
			}

			if jsonOutput {
				encoded, err := json.MarshalIndent(result, "", "  ")
				if err != nil {
					return fmt.Errorf("failed to encode publish result: %w", err)
				}

				fmt.Fprintln(cmd.OutOrStdout(), string(encoded))
				return nil
			}

			fmt.Fprintln(cmd.OutOrStdout(), "Agent metadata uploaded to Pinata.")
			fmt.Fprintf(cmd.OutOrStdout(), "CID: %s\n", result.CID)
			fmt.Fprintf(cmd.OutOrStdout(), "URI: %s\n", result.URI)

			return nil
		},
	}

	cmd.Flags().StringVarP(&filePath, "file", "f", "", "Path to agent metadata JSON file")
	cmd.Flags().StringVar(&jwt, "jwt", "", "Pinata JWT. Defaults to PINATA_JWT environment variable")
	cmd.Flags().StringVar(&name, "name", "", "Optional Pinata metadata name")
	cmd.Flags().BoolVar(&jsonOutput, "json", false, "Print publish result as JSON")

	return cmd
}
