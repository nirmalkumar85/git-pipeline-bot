package cmd

import (
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestCreatePullRequestCmd(t *testing.T) {
	// Create a new command
	cmd := NewCreatePullRequestCmd(nil)

	// Set the flags for testing
	cmd.Flags().Set("owner", "test-owner")
	cmd.Flags().Set("repo", "test-repo")
	cmd.Flags().Set("branch", "test-branch")
	cmd.Flags().Set("title", "Test PR")
	cmd.Flags().Set("body", "Test PR body")
	cmd.Flags().Set("interval", "10")

	// Parse flags
	err := cmd.Execute()

	// Assert that no error occurred during flag parsing
	assert.NoError(t, err)

	// Assert that the values are correctly loaded into viper
	assert.Equal(t, "test-owner", viper.GetString("owner"))
	assert.Equal(t, "test-repo", viper.GetString("repo"))
	assert.Equal(t, "test-branch", viper.GetString("branch"))
	assert.Equal(t, "Test PR", viper.GetString("title"))
	assert.Equal(t, "Test PR body", viper.GetString("body"))
	assert.Equal(t, 10, viper.GetInt("interval"))
}
