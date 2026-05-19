package main

import (
	"fmt"
	"os"

	"github.com/ronexlemon/agent-kit/cmd/agent-cli/commands"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "agent-kit",
	Short: "Developer tooling for ERC-8004 agents",
}

func Execute() {
	rootCmd.AddCommand(commands.NewCreateCommand())
	rootCmd.AddCommand(commands.NewPublishCommand())
	rootCmd.AddCommand(commands.NewRegisterCommand())

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
