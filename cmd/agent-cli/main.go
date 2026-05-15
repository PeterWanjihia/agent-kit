package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func main() {
	var agentName string
	var agentDesc string

	//TODO RESTRUCTURE
	var createCmd = &cobra.Command{
		Use:   "create",
		Short: "Register a new ERC-8004 agent",
		Run: func(cmd *cobra.Command, args []string) {
			reader := bufio.NewReader(os.Stdin)

			// 2. Interactive input for Name (if flag is empty)
			if agentName == "" {
				fmt.Print("Enter Agent Name: ")
				name, _ := reader.ReadString('\n')
				agentName = strings.TrimSpace(name)
			}

			// 3. Interactive input for Description (if flag is empty)
			if agentDesc == "" {
				fmt.Print("Enter Agent Description: ")
				desc, _ := reader.ReadString('\n')
				agentDesc = strings.TrimSpace(desc)
			}

		
			fmt.Println("\n--- Agent Registration Details ---")
			fmt.Printf("Name: %s\n", agentName)
			fmt.Printf("Description: %s\n", agentDesc)
		},
	}

	createCmd.Flags().StringVarP(&agentName, "name", "n", "", "Name of the agent")
	createCmd.Flags().StringVarP(&agentDesc, "desc", "d", "", "Description of the agent")

	var rootCmd = &cobra.Command{Use: "agent-kit"}
	rootCmd.AddCommand(createCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}