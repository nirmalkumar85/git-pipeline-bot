package utils

import (
	"testing"
	"time"
	"github.com/stretchr/testify/assert"
	"github.com/rs/zerolog"
)

func TestRunPeriodicTask(t *testing.T) {
	// Mock logger
	logger := zerolog.New(zerolog.NewConsoleWriter())

	// Track the number of times the task has run
	runCount := 0
	task := func() {
		runCount++
	}

	// Run the task every 100ms (to speed up the test)
	RunPeriodicTask(task, 100*time.Millisecond, &logger)

	// Give it some time to run
	time.Sleep(350 * time.Millisecond)

	// Assert that the task ran at least 3 times
	assert.GreaterOrEqual(t, runCount, 3)
}
