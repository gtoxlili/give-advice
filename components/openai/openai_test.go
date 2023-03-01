package openai

import (
	"context"
	"fmt"
	"testing"
)

func TestOpenai(t *testing.T) {
	err := Moderation(context.Background(), "Sample text goes here")
	if err != nil {
		fmt.Println(err)
	}
}
