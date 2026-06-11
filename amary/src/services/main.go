package services

import (
	"context"
	"time"
)

func WithErrorRetry(ctx context.Context, fun func() error, delay time.Duration) {
	maxRetry := 3
	for attempt := 0; attempt <= maxRetry; attempt++ {
		if err := fun(); err == nil {
			return
		}
		if attempt == maxRetry {
			break
		}

		select {
		case <-ctx.Done():
			return
		case <-time.After(delay):
		}
	}
}
