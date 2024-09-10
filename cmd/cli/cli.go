package cli

import (
	"fmt"
	"strings"

	"github.com/Amman30/EnvPilot/pkg/pilot"
	"github.com/spf13/cobra"
)

// NewCmdRoot creates the root command for the CLI
func NewCmdRoot() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "pilot",
		Short: "A CLI tool for managing environment variables",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Use 'pilot set <key>=<value> [flags]' to set environment variables")
		},
	}

	cmd.AddCommand(NewCmdSet())
	return cmd
}

// NewCmdSet creates the 'set' command for the CLI
func NewCmdSet() *cobra.Command {
	var valueType string
	var filename string

	var cmd = &cobra.Command{
		Use:   "set <key>=<value>",
		Short: "Set an environment variable",
		Args:  cobra.ExactArgs(1), // Expecting one positional argument
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return fmt.Errorf("invalid syntax. Use: pilot set <key>=<value> [flags]")
			}

			keyValue := args[0]
			parts := strings.SplitN(keyValue, "=", 2)
			if len(parts) != 2 {
				return fmt.Errorf("invalid syntax. Use: <key>=<value>")
			}
			key := parts[0]
			value := parts[1]

			if filename == "" {
				filename = ".env" 
			}

			err := pilot.Env.SetEnvValue(key, value, valueType, filename)
			if err != nil {
				return fmt.Errorf("error setting value: %s", err)
			}

			fmt.Printf("Successfully set %s=%s as %s in file %s\n", key, value, valueType, filename)
			return nil
		},
	}

	cmd.Flags().StringVarP(&valueType, "type", "t", "string", "Type of the value (string, int, bool, float)")
	cmd.Flags().StringVarP(&filename, "file", "f", "", "The file to save the environment variable (default is .env)")
	return cmd
}
