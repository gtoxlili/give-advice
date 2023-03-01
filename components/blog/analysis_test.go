package blog

import (
	"context"
	"fmt"
	"testing"
)

func TestAnalysis(t *testing.T) {
	ctx := context.Background()
	_, err := Analysis(ctx, "https://www.yystv.cn/p/10509")
	if err != nil {
		fmt.Println(fmt.Sprintf("Analysis err: %v", err))
	}
}
