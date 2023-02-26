package openai

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"sync"

	json "github.com/bytedance/sonic"
	"github.com/gtoxlili/give-advice/common/fail"
	"github.com/gtoxlili/give-advice/common/stream"
	log "github.com/sirupsen/logrus"
)

var Token = ""

func init() {
	log.Info("OpenAI API Token: ", Token)
}

type completionsReq struct {
	Model            string  `json:"model"`
	Prompt           string  `json:"prompt"`
	MaxTokens        int     `json:"max_tokens"`
	Temperature      float32 `json:"temperature"`
	TopP             float32 `json:"top_p"`
	FrequencyPenalty float32 `json:"frequency_penalty"`
	PresencePenalty  float32 `json:"presence_penalty"`
	Stream           bool    `json:"stream"`
	N                int     `json:"n"`
}

type completionsRes struct {
	Choices []struct {
		Text         string `json:"text"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
}

var completionsReqPool = sync.Pool{
	New: func() interface{} {
		return &completionsReq{
			"text-davinci-003",
			"",
			1536,
			0.7,
			1.0,
			0.0,
			0.0,
			true,
			1,
		}
	},
}

func generateCompletionsBodyReader(prompt string) io.Reader {
	body := completionsReqPool.Get().(*completionsReq)
	body.Prompt = prompt
	defer completionsReqPool.Put(body)

	byteBody, _ := json.Marshal(body)
	return bytes.NewReader(byteBody)
}

func completions(prompt string) func(entry *stream.Entry[string]) {
	return func(entry *stream.Entry[string]) {
		req, err := http.NewRequestWithContext(
			entry.Ctx,
			"POST",
			"https://api.openai.com/v1/completions",
			generateCompletionsBodyReader(prompt),
		)
		if err != nil {
			entry.Panic(err)
		}

		req.Header.Set("Authorization", "Bearer "+token(entry.Ctx))
		req.Header.Set("Content-Type", "application/json")
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Warning(err.Error())
			entry.Panic(err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			if resp.StatusCode == http.StatusTooManyRequests {
				if fail.Retry(entry.Ctx, func(ct context.Context) {
					log.Warning("正在尝试重新访问 completions 接口 ：第", ct.Value("retry").(int), "次")
					completions(prompt)(entry.WithContext(ct))
				}) != nil {
					log.Warning(err.Error())
					entry.Panic(err)
				}
				return
			}
			log.Warning(resp.Status)
			entry.Panic(fmt.Errorf("completions server anomalies status: %s", resp.Status))
		}
		res := &completionsRes{}
		// 返回的是一个 stream，需要读取到最后一个 \n\n 为止
		err = stream.ReadFlow(resp.Body, []byte{'\n', '\n'},
			func(msg []byte) bool {
				if err := json.Unmarshal(msg[5:], res); err != nil {
					// FinishReason 没有出现但返回了 [DONE] 的情况
					if bytes.Contains(msg[5:], []byte("[DONE]")) {
						return true
					}
					log.Warning(err.Error())
					entry.Panic(err)
				}
				if res.Choices[0].Text == "\n" {
					entry.Send("\\n")
				} else {
					entry.Send(res.Choices[0].Text)
				}
				return res.Choices[0].FinishReason == "stop" || res.Choices[0].FinishReason == "length"
			},
		)
		if err != nil {
			log.Warning(err.Error())
			entry.Panic(err)
		}
	}
}

func token(ctx context.Context) string {
	// 尝试从上下文中获取 token，取不到则从环境变量中获取
	if token, ok := ctx.Value("OpenAI-Auth-Key").(string); ok && token != "" {
		return token
	}
	return Token
}
