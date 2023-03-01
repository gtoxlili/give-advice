package deepl

import (
	"context"
	"fmt"
	"testing"
)

func TestTranslate(t *testing.T) {
	ctx := context.Background()
	text := "hello world"
	result, err := Translate(ctx, text, "ZH")
	fmt.Println(result, err)
}
