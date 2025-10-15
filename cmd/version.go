// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"fmt"

	"github.com/agentbay/agentbay-cli/internal/config"
	"github.com/spf13/cobra"
)

var (
	Version   = "dev"
	GitCommit = "unknown"
	BuildDate = "unknown"
)

var VersionCmd = &cobra.Command{
	Use:     "version",
	Short:   "Show version information",
	Long:    "Display version, git commit, and build date information",
	GroupID: "core",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Printf("AgentBay CLI version %s\n", Version)
		fmt.Printf("Git commit: %s\n", GitCommit)
		fmt.Printf("Build date: %s\n", BuildDate)

		// Show environment information
		env := config.GetEnvironment()
		envConfig := config.GetEnvironmentConfig()
		fmt.Printf("Environment: %s\n", env)
		fmt.Printf("Endpoint: %s\n", envConfig.Endpoint)

		return nil
	},
}
