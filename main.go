// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/agentbay/agentbay-cli/cmd"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:               "agentbay",
	Short:             "AgentBay CLI",
	Long:              "Command line interface for AgentBay services",
	DisableAutoGenTag: true,
	SilenceUsage:      false,
	SilenceErrors:     false,
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

func init() {
	// Add command groups
	rootCmd.AddGroup(&cobra.Group{ID: "core", Title: "Core Commands"})
	rootCmd.AddGroup(&cobra.Group{ID: "management", Title: "Management Commands"})

	// Add commands
	rootCmd.AddCommand(cmd.VersionCmd)
	rootCmd.AddCommand(cmd.LoginCmd)
	rootCmd.AddCommand(cmd.LogoutCmd)
	rootCmd.AddCommand(cmd.ImageCmd)

	// Global flags
	rootCmd.CompletionOptions.HiddenDefaultCmd = true
	rootCmd.PersistentFlags().BoolP("help", "", false, "help for agentbay")
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "Enable verbose output")
	rootCmd.Flags().BoolP("version", "", false, "Display the version of AgentBay CLI")

	// Handle version flag and verbose flag
	rootCmd.PersistentPreRun = func(command *cobra.Command, args []string) {
		// Set up logging based on verbose flag
		verbose, _ := command.Flags().GetBool("verbose")
		if verbose {
			log.SetLevel(log.DebugLevel)
		} else {
			log.SetLevel(log.InfoLevel)
		}

		// Set log format to be more CLI-friendly
		log.SetFormatter(&log.TextFormatter{
			DisableTimestamp: true,
			DisableColors:    false,
		})
	}

	// Handle version flag
	rootCmd.PreRun = func(command *cobra.Command, args []string) {
		versionFlag, _ := command.Flags().GetBool("version")
		if versionFlag {
			err := cmd.VersionCmd.RunE(command, []string{})
			if err != nil {
				log.Fatal(err)
			}
			os.Exit(0)
		}
	}
}

func main() {
	// Load environment variables
	_ = godotenv.Load()

	// Execute root command
	err := rootCmd.Execute()
	if err != nil {
		// Error messages are already displayed by cobra
		os.Exit(1)
	}
}
