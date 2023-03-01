package openai

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gtoxlili/give-advice/common/ht"
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

func Moderation(ctx context.Context, text string) error {
	res, err := ht.Request[modRes](ctx, "POST", "https://api.openai.com/v1/moderations",
		&modReq{Input: text}, ht.Header{
			"Authorization": "Bearer " + token(ctx),
			"Content-Type":  "application/json",
		})
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("[Moderation] anomalies : %s", res.Status)
	}
	data := res.Data
	if data.Results[0].Flagged {
		for k, v := range data.Results[0].Categories {
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
