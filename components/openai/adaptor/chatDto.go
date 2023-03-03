package adaptor

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gtoxlili/give-advice/common/stream"
	"github.com/gtoxlili/give-advice/common/validate"
	"github.com/gtoxlili/give-advice/components/openai/completion"
)

type ChatDto struct {
	Topic   string `json:"topic"`
	Records []struct {
		Q string `json:"q" validate:"required"`
		A string `json:"a"`
	} `json:"records" validate:"required"`
}

func (dto *ChatDto) Bind(_ *http.Request) error {
	return validate.Struct(dto)
}

func (dto *ChatDto) Completion(ctx context.Context) <-chan stream.Result[string] {
	msgItems := make([]completion.Message, 0, len(dto.Records))
	if dto.Topic != "" {
		msgItems = append(msgItems, completion.Message{
			Role:    "system",
			Content: dto.Topic,
		})
	}
	for _, record := range dto.Records {
		msgItems = append(msgItems, completion.Message{
			Role:    "user",
			Content: record.Q,
		}, completion.Message{
			Role:    "assistant",
			Content: record.A,
		})
	}
	return stream.AsyncReflow(ctx, completion.Do(msgItems))
}

// Review 需要被审查的内容
func (dto *ChatDto) Review() string {
	return fmt.Sprintf("%s\n%s", dto.Topic, dto.Records[len(dto.Records)-1].Q)
}
