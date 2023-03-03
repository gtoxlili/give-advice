package completion

import (
	"context"
	"testing"

	"github.com/gtoxlili/give-advice/common/stream"
)

func TestDto_Completion(t *testing.T) {
	ctx := context.Background()
	for _ = range stream.AsyncReflow(ctx, Do(
		[]Message{
			{
				Role:    "user",
				Content: "我想买一台电脑",
			},
		}),
	) {
		// do nothing
	}
}
