package main

import (
	"testing"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	// Set up viper configuration for testing
	viper.SetConfigFile("test_config.yaml") // Use a test config file

	// Mock configuration values
	viper.Set("interval", 10)
	viper.Set("owner", "test-owner")
	viper.Set("repo", "test-repo")
	viper.Set("branch", "test-branch")

	// Load configuration
	err := loadConfig()

	// Assert that there were no errors
	assert.NoError(t, err)

	// Assert that viper has loaded the correct values
	assert.Equal(t, 10, viper.GetInt("interval"))
	assert.Equal(t, "test-owner", viper.GetString("owner"))
	assert.Equal(t, "test-repo", viper.GetString("repo"))
	assert.Equal(t, "test-branch", viper.GetString("branch"))
}
