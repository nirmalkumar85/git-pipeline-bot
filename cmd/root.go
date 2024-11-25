package cmd

import (
	"time"

	"github.com/nirmalkumar85/git-pipeline-bot/github"
	"github.com/nirmalkumar85/git-pipeline-bot/template"
	"github.com/nirmalkumar85/git-pipeline-bot/utils"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// NewCreatePullRequestCmd creates the command for periodic pull request creation
func NewCreatePullRequestCmd(logger *zerolog.Logger) *cobra.Command {
	var repoOwner, repoName, branchName, prTitle, prBody string
	var interval int

	cmd := &cobra.Command{
		Use:   "create-pullrequest",
		Short: "Create a pull request from a forked branch every interval",
		Run: func(cmd *cobra.Command, args []string) {
			// Retrieve configuration from viper
			interval = viper.GetInt("interval")
			repoOwner = viper.GetString("owner")
			repoName = viper.GetString("repo")
			branchName = viper.GetString("branch")
			prTitle = viper.GetString("title")
			prBody = viper.GetString("body")

			// Ensure required config values are set
			if repoOwner == "" || repoName == "" || branchName == "" {
				logger.Fatal().Msg("GitHub configuration (owner, repo, branch) must be provided.")
			}

			// Periodically create a pull request
			utils.RunPeriodicTask(func() {
				// Generate CRD spec content using the Go template
				content, err := template.GenerateProbeCRDSpec()
				if err != nil {
					logger.Error().Err(err).Msg("Error generating CRD spec")
					return
				}

				// Create the pull request
				err = github.CreatePullRequest(repoOwner, repoName, branchName, prTitle, prBody, content, logger)
				if err != nil {
					logger.Error().Err(err).Msg("Error creating pull request")
				} else {
					logger.Info().Msg("Pull request created successfully!")
				}
			}, time.Duration(interval)*time.Minute, logger)
		},
	}

	// Allow flags for command-line overrides
	cmd.Flags().Int("interval", 5, "Interval in minutes between pull requests")
	cmd.Flags().String("owner", "", "GitHub repository owner")
	cmd.Flags().String("repo", "", "GitHub repository name")
	cmd.Flags().String("branch", "", "Forked branch name")
	cmd.Flags().String("title", "New Probe CRD Spec", "Pull request title")
	cmd.Flags().String("body", "This is an automated pull request to update probe CRD spec.", "Pull request body")

	// Bind flags to viper
	viper.BindPFlags(cmd.Flags())

	return cmd
}
