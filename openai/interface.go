package openai

import (
	"context"
	"fmt"
	"github.com/gtoxlili/give-advice/common/stream"
	"strings"
	"unicode"
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

func GeneralAdvice(ctx context.Context, noun string, description string) <-chan stream.Result[string] {
	prompt := fmt.Sprintf(detectLanguage(description, generalPrompt), noun, description)
	return stream.AsyncReflow(ctx, completions(prompt))
}

func QuestionArticles(ctx context.Context, noun string, description string) <-chan stream.Result[string] {
	// noun : {title} || {question}
	arr := strings.Split(noun, "||")

	title := strings.TrimSpace(arr[0])
	question := strings.TrimSpace(arr[1])

	prompt := fmt.Sprintf(detectLanguage(question, articlePrompt), title, description, question)
	return stream.AsyncReflow(ctx, completions(prompt))
}

func detectLanguage(input string, format func(lang string) string) string {
	var (
		enCount int
		zhCount int
		jaCount int
	)

	for _, r := range input {
		switch {
		case unicode.In(r, unicode.Han):
			zhCount++
		case unicode.In(r, unicode.Katakana, unicode.Hiragana):
			jaCount++
		case unicode.In(r, unicode.Letter, unicode.Punct):
			enCount++
		}
	}

	if zhCount > jaCount && zhCount > enCount/2 {
		return format("Chinese")
	} else if jaCount > zhCount && jaCount > enCount/2 {
		return format("Japanese")
	} else {
		return format("English")
	}
}

func generalPrompt(lang string) string {
	switch lang {
	case "Chinese":
		return "请用中文通过以下的内容填充为一篇完整的%s，用 markdown 格式以分点叙述的形式输出：%s。"
	case "Japanese":
		return "以下の内容を日本語で記入して、完全な%sを形成し、箇条書きでマークダウン形式で出力してください：%s。"
	default:
		return "Please fill in the following content in English to form a complete %s, and output it in markdown format in the form of bullet points: %s."
	}
}

func articlePrompt(lang string) string {
	switch lang {
	case "Chinese":
		return "请阅读题为「%s」的文章\n'''\n%s\n'''\n然后用中文回答以下问题：%s。"
	case "Japanese":
		return "「%s」というタイトルの記事を読んでください。\n'''\n%s\n'''\nそして、以下の質問に日本語で答えてください：%s。"
	default:
		return "Please read the article titled \"%s\"'''\n%s\n'''\n nand answer the following question in English: %s"
	}
}
