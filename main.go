package main

import (
	"os"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/rs/zerolog"
	"github.com/yourusername/pullrequest-cli/cmd"
	"github.com/yourusername/pullrequest-cli/github"
	"github.com/yourusername/pullrequest-cli/template"
	"github.com/yourusername/pullrequest-cli/utils"
)

func main() {
	// Initialize the logger
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()

	// Load configuration using Viper
	if err := loadConfig(); err != nil {
		logger.Fatal().Err(err).Msg("Failed to load config")
	}

	// Initialize the Cobra root command
	rootCmd := &cobra.Command{
		Use:   "pullrequest-cli",
		Short: "A CLI tool to automate pull requests from forked branches.",
	}

	// Add the command for creating pull requests
	rootCmd.AddCommand(cmd.NewCreatePullRequestCmd(logger))

	// Execute the command
	if err := rootCmd.Execute(); err != nil {
		logger.Fatal().Err(err).Msg("Error executing command")
	}
}

// loadConfig loads configuration from a file or environment variables
func loadConfig() error {
	viper.SetConfigName("config") // Name of config file (without extension)
	viper.AddConfigPath(".")      // Look for the config in the current directory
	viper.AutomaticEnv()          // Automatically read environment variables
	viper.SetDefault("interval", 5)

	return viper.ReadInConfig()
}
