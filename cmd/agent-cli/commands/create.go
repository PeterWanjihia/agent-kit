package commands

import (
	"fmt"
	"github.com/ronexlemon/agent-kit/pkg/client"
	"github.com/ronexlemon/agent-kit/pkg/types"
	"github.com/ronexlemon/agent-kit/utils/agent"
	"github.com/spf13/cobra"
)




func NewCreateCommand() *cobra.Command {
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
			if err := agent.CollectMissingRequiredInput(&input, noInteractive); err != nil {
				return err
			}
			parsedServices, err := agent.ParseServices(services)
			if err != nil {
				return err
			}

			input.Services = parsedServices

			metadata, err := client.CreateAgentMetadata(input)
			if err != nil {
				return err
			}

			if outPath != "" {
				if err := agent.WriteMetadataFile(outPath, metadata); err != nil {
					return err
				}

				if !jsonOutput {
					fmt.Fprintf(cmd.OutOrStdout(), "Agent metadata created successfully.\nOutput: %s\n", outPath)
					return nil
				}
			}

			if jsonOutput {
				return agent.PrintJSON(cmd, metadata)
			}

			agent.PrintSummary(cmd, metadata)
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


