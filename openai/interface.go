package openai

import (
	"context"
	"fmt"
	"github.com/gtoxlili/give-advice/common/stream"
	"strings"
)

type AIFunc func(ctx context.Context, noun string, description string) <-chan stream.Result[string]

func Factory(typ string) AIFunc {
	switch typ {
	case "advice":
		return GeneralAdvice
	case "article":
		return QuestionArticles
	default:
		return nil
	}
}

// GeneralAdvice
// 请用中文通过以下的内容填充为一篇完整的{}，用 markdown 格式以分点叙述的形式输出：{}
func GeneralAdvice(ctx context.Context, noun string, description string) <-chan stream.Result[string] {
	prompt := fmt.Sprintf("请通过以下的内容填充为一篇完整的%s，用 markdown 格式以分点叙述的形式输出：%s。", noun, description)
	return stream.AsyncReflow(ctx, completions(prompt))
}

// QuestionArticles
// 请通过这篇名为“{}”的文章
// ”“”
// {}
// ”“”
// 用中文回答我这个问题：{}
func QuestionArticles(ctx context.Context, noun string, description string) <-chan stream.Result[string] {
	// noun : {title} || {question}
	arr := strings.Split(noun, "||")

	title := strings.TrimSpace(arr[0])
	question := strings.TrimSpace(arr[1])

	prompt := fmt.Sprintf("请通过这篇名为“%s”的文章:\n”“”\n%s\n”“”\n用中文回答我这个问题：%s。", title, description, question)
	return stream.AsyncReflow(ctx, completions(prompt))
}
