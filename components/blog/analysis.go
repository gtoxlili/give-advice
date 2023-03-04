package blog

import (
	"bytes"
	"context"
	"fmt"

	"github.com/gtoxlili/advice-hub/common/ht"
)

// AnalysisResult 解析结果
type AnalysisResult struct {
	// 标题
	Title string `json:"title"`
	// 正文
	Content string `json:"content"`
	// 梗概
	Abstract string `json:"abstract"`
}

func Analysis(ctx context.Context, url string) (*AnalysisResult, error) {
	viewSource, err := ht.RequestWithUnmarshalHandler(ctx,
		"GET", url, nil, nil,
		func(buf *bytes.Buffer) string {
			return buf.String()
		})
	if err != nil {
		return nil, err
	}
	fmt.Println(viewSource.Data)
	return nil, nil
}
