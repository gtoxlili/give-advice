package completion

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"

	json "github.com/bytedance/sonic"
	"github.com/gtoxlili/advice-hub/common/fail"
	"github.com/gtoxlili/advice-hub/common/ht"
	"github.com/gtoxlili/advice-hub/common/stream"
	"github.com/gtoxlili/advice-hub/constants"
	log "github.com/sirupsen/logrus"
)

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type completionReq struct {
	Model       string    `json:"model"`
	Messages    []Message `json:"messages"`
	Temperature float32   `json:"temperature"`
	Stream      bool      `json:"stream"`
}

type completionRes struct {
	Choices []struct {
		Delta struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		} `json:"delta"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
}

func bodyReader(msgItems []Message) io.Reader {
	body := &completionReq{
		"gpt-3.5-turbo",
		msgItems,
		0.7,
		true,
	}
	byteBody, _ := json.Marshal(body)
	return bytes.NewReader(byteBody)
}

func Do(msgItems []Message) func(entry *stream.Entry[string]) {
	return func(entry *stream.Entry[string]) {
		req, err := http.NewRequestWithContext(
			entry.Ctx,
			"POST",
			"https://api.openai.com/v1/chat/completions",
			bodyReader(msgItems),
		)
		if err != nil {
			entry.Panic(err)
		}
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+constants.CalcOpenAIToken(entry.Ctx))

		resp, err := ht.Client.Do(req)
		if err != nil {
			log.Warning(err.Error())
			entry.Panic(err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			if resp.StatusCode == http.StatusTooManyRequests {
				if fail.Retry(entry.Ctx, func(ct context.Context) {
					// Trying to re-access the completions interface: 3rd time
					log.Warning("Trying to re-access the [Completions] interface: ", ct.Value("retry").(int), "th time")
					Do(msgItems)(entry.WithContext(ct))
				}) != nil {
					log.Warning(err.Error())
					entry.Panic(err)
				}
				return
			}
			log.Warning(resp.Status)
			entry.Panic(fmt.Errorf("[Completions] anomalies : %s", resp.Status))
		}

		res := &completionRes{}
		if stream.ReadFlow(resp.Body, []byte{'\n', '\n'},
			func(msg []byte) bool {
				if res.Choices != nil {
					res.Choices[0].Delta.Content = ""
				}
				if err = json.Unmarshal(msg[5:], res); err != nil {
					// FinishReason 没有出现但返回了 [DONE] 的情况
					if bytes.Contains(msg[5:], []byte("[DONE]")) {
						return true
					}
					log.Warning(err.Error())
					entry.Panic(err)
				}
				if res.Choices[0].Delta.Content != "" {
					entry.Send(res.Choices[0].Delta.Content)
				}
				return res.Choices[0].FinishReason == "stop" || res.Choices[0].FinishReason == "length"
			},
		) != nil {
			log.Warning(err.Error())
			entry.Panic(err)
		}
	}
}
