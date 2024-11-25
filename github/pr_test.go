package github

import (
	"errors"
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/google/go-github/v66/github"
	"github.com/rs/zerolog"
)

// Mocking GitHub client methods
type MockGitHubClient struct {
	mock.Mock
}

func (m *MockGitHubClient) PullRequestsCreate(owner, repo string, pr *github.NewPullRequest) (*github.PullRequest, *github.Response, error) {
	args := m.Called(owner, repo, pr)
	return args.Get(0).(*github.PullRequest), args.Get(1).(*github.Response), args.Error(2)
}

func TestCreatePullRequest(t *testing.T) {
	mockClient := new(MockGitHubClient)
	logger := zerolog.New(zerolog.NewConsoleWriter())

	// Setup the mock to return expected values
	mockClient.On("PullRequestsCreate", "owner", "repo", mock.Anything).Return(&github.PullRequest{}, &github.Response{}, nil)

	err := CreatePullRequest("owner", "repo", "branch", "Title", "Body", "content", &logger)

	// Assert that there were no errors and the method was called as expected
	assert.NoError(t, err)
	mockClient.AssertExpectations(t)
}

func TestCreatePullRequest_Error(t *testing.T) {
	mockClient := new(MockGitHubClient)
	logger := zerolog.New(zerolog.NewConsoleWriter())

	// Setup the mock to return an error
	mockClient.On("PullRequestsCreate", "owner", "repo", mock.Anything).Return(nil, nil, errors.New("GitHub API error"))

	err := CreatePullRequest("owner", "repo", "branch", "Title", "Body", "content", &logger)

	// Assert that there was an error
	assert.Error(t, err)
}
