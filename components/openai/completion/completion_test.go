package completion

import (
	"context"
	"fmt"
	"testing"

	"github.com/gtoxlili/advice-hub/common/stream"
)

func TestDto_Completion(t *testing.T) {
	ctx := context.Background()
	for v := range stream.AsyncReflow(ctx, Do(
		[]Message{
			{
				Role:    "system",
				Content: "你是一个精通习近平思想的人",
			},
			{
				Role:    "user",
				Content: "请围绕对习近平新时代中国特色社会主义思想的认识情况和掌握水平，对不同时代特别是中国特色社会主义新时代青年学生使命担当的认识和实践情况，写一篇不少于1500字论文。题目自拟",
			},
		}),
	) {
		if v.Err != nil {
			t.Fatal(v.Err)
		}
		fmt.Print(v.Val)
	}
}
