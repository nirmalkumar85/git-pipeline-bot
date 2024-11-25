package github

import (
	"context"
	"fmt"
	"github.com/google/go-github/v66/github" // Use v66 for GitHub API v3
	"golang.org/x/oauth2"
	"github.com/spf13/viper"
	"github.com/rs/zerolog"
)

// CreatePullRequest creates a GitHub pull request from a forked branch
func CreatePullRequest(owner, repo, branch, title, body, content string, logger *zerolog.Logger) error {
	client, err := newGitHubClient(logger)
	if err != nil {
		return err
	}

	// Prepare the pull request data
	pr := &github.NewPullRequest{
		Title: &title,
		Body:  &body,
		Base:  github.String("main"), // Target branch to merge into
		Head:  github.String(branch), // Source branch in the fork
	}

	// Create the pull request on GitHub
	_, _, err = client.PullRequests.Create(context.Background(), owner, repo, pr)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to create pull request")
		return err
	}

	logger.Info().Msg("Pull request created successfully!")
	return nil
}

// newGitHubClient creates and returns an authenticated GitHub client
func newGitHubClient(logger *zerolog.Logger) (*github.Client, error) {
	token := viper.GetString("github_token")
	if token == "" {
		logger.Fatal().Msg("GitHub token is required but not provided")
		return nil, fmt.Errorf("GitHub token not provided")
	}

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(context.Background(), ts)
	return github.NewClient(tc), nil
}
