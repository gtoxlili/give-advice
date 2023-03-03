package openai

import (
	"context"
	"net/http"
	"strings"

	"github.com/gtoxlili/give-advice/common/stream"
	"github.com/gtoxlili/give-advice/components/openai/adaptor"
)

type Adaptor interface {
	Completion(context.Context) <-chan stream.Result[string]
	Bind(*http.Request) error
	Review() string
}

// Factory 根据 Type 创建对应的 Adaptor
func Factory(typ string) Adaptor {
	switch strings.ToLower(typ) {
	case "chat":
		return &adaptor.ChatDto{}
	}
	return nil
}
