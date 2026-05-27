// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

package cmd_test

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"

	"github.com/agentbay/agentbay-cli/cmd"
)

func TestDockerCmd(t *testing.T) {
	t.Run("docker command has correct metadata", func(t *testing.T) {
		assert.Equal(t, "docker", cmd.DockerCmd.Use)
		assert.Equal(t, "Docker image build & push operations", cmd.DockerCmd.Short)
		assert.Equal(t, "management", cmd.DockerCmd.GroupID)
	})

	t.Run("docker has all subcommands including share/unshare/list-shares", func(t *testing.T) {
		children := cmd.DockerCmd.Commands()
		names := make([]string, 0, len(children))
		for _, c := range children {
			names = append(names, c.Name())
		}
		assert.Contains(t, names, "login")
		assert.Contains(t, names, "tag")
		assert.Contains(t, names, "push")
		assert.Contains(t, names, "share")
		assert.Contains(t, names, "unshare")
		assert.Contains(t, names, "list-shares")
	})
}

func TestDockerShareCmd(t *testing.T) {
	var shareCmd = findSubCmd(cmd.DockerCmd, "share")

	t.Run("share command has correct metadata", func(t *testing.T) {
		assert.NotNil(t, shareCmd)
		assert.Equal(t, "share [<target-uid>]", shareCmd.Use)
		assert.Equal(t, "Share the Docker repo with another Alibaba Cloud account", shareCmd.Short)
	})

	t.Run("share command has --target-uid flag", func(t *testing.T) {
		assert.NotNil(t, shareCmd)
		flag := shareCmd.Flags().Lookup("target-uid")
		assert.NotNil(t, flag)
		assert.Equal(t, "0", flag.DefValue)
	})

	t.Run("share command accepts positional arg or --target-uid flag", func(t *testing.T) {
		assert.NotNil(t, shareCmd)
		flag := shareCmd.Flags().Lookup("target-uid")
		assert.NotNil(t, flag)
		// Not strictly required as flag since positional arg is also accepted
		annotations := flag.Annotations
		_, isRequired := annotations["cobra_annotation_bash_completion_one_required_flag"]
		assert.False(t, isRequired, "--target-uid should not be marked as required since positional arg is supported")
		// Args allows 0 or 1 positional arg, rejects 2+
		assert.NoError(t, shareCmd.Args(nil, []string{}))
		assert.NoError(t, shareCmd.Args(nil, []string{"1234567890"}))
		assert.Error(t, shareCmd.Args(nil, []string{"1", "2"}))
	})
}

func TestDockerUnshareCmd(t *testing.T) {
	var unshareCmd = findSubCmd(cmd.DockerCmd, "unshare")

	t.Run("unshare command has correct metadata", func(t *testing.T) {
		assert.NotNil(t, unshareCmd)
		assert.Equal(t, "unshare [<target-uid>]", unshareCmd.Use)
		assert.Equal(t, "Cancel sharing the Docker repo with another Alibaba Cloud account", unshareCmd.Short)
	})

	t.Run("unshare command has --target-uid flag", func(t *testing.T) {
		assert.NotNil(t, unshareCmd)
		flag := unshareCmd.Flags().Lookup("target-uid")
		assert.NotNil(t, flag)
		assert.Equal(t, "0", flag.DefValue)
	})

	t.Run("unshare command accepts positional arg or --target-uid flag", func(t *testing.T) {
		assert.NotNil(t, unshareCmd)
		flag := unshareCmd.Flags().Lookup("target-uid")
		assert.NotNil(t, flag)
		// Not strictly required as flag since positional arg is also accepted
		annotations := flag.Annotations
		_, isRequired := annotations["cobra_annotation_bash_completion_one_required_flag"]
		assert.False(t, isRequired, "--target-uid should not be marked as required since positional arg is supported")
		// Args allows 0 or 1 positional arg, rejects 2+
		assert.NoError(t, unshareCmd.Args(nil, []string{}))
		assert.NoError(t, unshareCmd.Args(nil, []string{"1234567890"}))
		assert.Error(t, unshareCmd.Args(nil, []string{"1", "2"}))
	})
}

func TestDockerListSharesCmd(t *testing.T) {
	var listSharesCmd = findSubCmd(cmd.DockerCmd, "list-shares")

	t.Run("list-shares command has correct metadata", func(t *testing.T) {
		assert.NotNil(t, listSharesCmd)
		assert.Equal(t, "list-shares", listSharesCmd.Use)
		assert.Equal(t, "List Docker repo sharing information", listSharesCmd.Short)
	})

	t.Run("list-shares command has --direction flag", func(t *testing.T) {
		assert.NotNil(t, listSharesCmd)
		flag := listSharesCmd.Flags().Lookup("direction")
		assert.NotNil(t, flag)
		assert.Equal(t, "", flag.DefValue)
	})

	t.Run("list-shares command --direction is required", func(t *testing.T) {
		assert.NotNil(t, listSharesCmd)
		flag := listSharesCmd.Flags().Lookup("direction")
		assert.NotNil(t, flag)
		annotations := flag.Annotations
		_, isRequired := annotations["cobra_annotation_bash_completion_one_required_flag"]
		assert.True(t, isRequired, "--direction should be marked as required")
	})

	t.Run("list-shares command has --output / -o flag", func(t *testing.T) {
		assert.NotNil(t, listSharesCmd)
		flag := listSharesCmd.Flags().Lookup("output")
		assert.NotNil(t, flag)
		assert.Equal(t, "", flag.DefValue)
		assert.Equal(t, "o", flag.Shorthand)
	})
}

// findSubCmd finds a subcommand by name within a parent command.
func findSubCmd(parent interface{ Commands() []*cobra.Command }, name string) *cobra.Command {
	for _, c := range parent.Commands() {
		if c.Name() == name {
			return c
		}
	}
	return nil
}
