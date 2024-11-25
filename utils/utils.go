package utils

import (
	"time"
	"github.com/rs/zerolog"
)

// RunPeriodicTask runs a function periodically based on the interval
func RunPeriodicTask(task func(), interval time.Duration, logger *zerolog.Logger) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			task()
		}
	}
}
