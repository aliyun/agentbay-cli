// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/agentbay/agentbay-cli/internal/client"
	"github.com/spf13/cobra"
)

// printRequestIDIfVerbose prints "[DEBUG] RequestId: <id>" to stderr when -v is set and id is non-empty.
func printRequestIDIfVerbose(cmd *cobra.Command, requestID string) {
	if requestID == "" {
		return
	}
	verbose, _ := cmd.Flags().GetBool("verbose")
	if !verbose {
		return
	}
	fmt.Fprintf(os.Stderr, "[DEBUG] RequestId: %s\n", requestID)
}

// printRequestIDFromErrIfVerbose extracts client.ErrWithRequestID from err and prints RequestId when -v.
func printRequestIDFromErrIfVerbose(cmd *cobra.Command, err error) {
	if err == nil {
		return
	}
	verbose, _ := cmd.Flags().GetBool("verbose")
	if !verbose {
		return
	}
	var errWithID *client.ErrWithRequestID
	if errors.As(err, &errWithID) && errWithID.RequestID != "" {
		fmt.Fprintf(os.Stderr, "[DEBUG] RequestId: %s\n", errWithID.RequestID)
	}
}
