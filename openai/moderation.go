package openai

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"

	json "github.com/bytedance/sonic"
)

// curl https://api.openai.com/v1/moderations \
// -X POST \
// -H "Content-Type: application/json" \
// -H "Authorization: Bearer $OPENAI_API_KEY" \
// -d '{"input": "Sample text goes here"}'

type modRes struct {
	Results []struct {
		Flagged    bool            `json:"flagged"`
		Categories map[string]bool `json:"categories"`
	} `json:"results"`
}

type modReq struct {
	Input string `json:"input"`
}

func generateModBodyReader(text string) io.Reader {
	body := &modReq{Input: text}
	j, _ := json.Marshal(body)
	return bytes.NewReader(j)
}

func Moderation(ctx context.Context, text string) error {
	req, _ := http.NewRequest("POST", "https://api.openai.com/v1/moderations", generateModBodyReader(text))
	req.Header.Set("Authorization", "Bearer "+token(ctx))
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("moderations server anomalies status: %s", resp.Status)
	}
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	res := &modRes{}
	json.Unmarshal(b, res)
	if res.Results[0].Flagged {
		for k, v := range res.Results[0].Categories {
			if v {
				return fmt.Errorf("violation content: %s", k)
			}
		}
	}
	return nil
}

func ModerationChan(ctx context.Context, text string) <-chan error {
	ch := make(chan error, 1)
	go func() {
		err := Moderation(ctx, text)
		if err != nil {
			ch <- err
		}
	}()
	return ch
}
