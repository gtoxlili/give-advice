package adaptor

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gtoxlili/advice-hub/common/stream"
	"github.com/gtoxlili/advice-hub/common/validate"
	"github.com/gtoxlili/advice-hub/components/openai/completion"
)

const lastRecords = 5

type ChatDto struct {
	Topic    string `json:"topic"`
	BaseInfo string `json:"baseInfo"`
	Records  []struct {
		Q string `json:"q" validate:"required"`
		A string `json:"a"`
	} `json:"records" validate:"required"`
}

func (dto *ChatDto) Bind(_ *http.Request) error {
	return validate.Struct(dto)
}

func (dto *ChatDto) Completion(ctx context.Context) <-chan stream.Result[string] {
	// 仅使用最后 5 条记录
	l := len(dto.Records) - lastRecords
	if l < 0 {
		l = 0
	}
	msgItems := make([]completion.Message, 0, l+1)
	if dto.Topic != "" {
		msgItems = append(msgItems, completion.Message{
			Role:    "system",
			Content: dto.BaseInfo,
		})
	}
	for _, record := range dto.Records[l:] {
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
