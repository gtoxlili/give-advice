package openai

import (
	"context"
	"fmt"
	"strings"
	"testing"

	json "github.com/bytedance/sonic"
)

func TestOpenai(t *testing.T) {
	err := Moderation(context.Background(), "Sample text goes here")
	if err != nil {
		fmt.Println(err)
	}
}

type abc struct {
	Results string `json:"results"`
}

func TestJson(t *testing.T) {
	bytes := []byte(`{"results":"\n"}`)
	abc := &abc{}
	// EscapeHTML
	_ = json.Unmarshal(bytes, abc)
	fmt.Println(abc)
}

type Pattern struct {
	prefix   string
	suffix   string
	wildcard bool
}

func NewPattern(value string) Pattern {
	p := Pattern{}
	if i := strings.IndexByte(value, '*'); i >= 0 {
		p.wildcard = true
		p.prefix = value[0:i]
		p.suffix = value[i+1:]
	} else {
		p.prefix = value
	}
	return p
}

func (p Pattern) Match(v string) bool {
	if !p.wildcard {
		if p.prefix == v {
			return true
		} else {
			return false
		}
	}
	return len(v) >= len(p.prefix+p.suffix) && strings.HasPrefix(v, p.prefix) && strings.HasSuffix(v, p.suffix)
}

func TestPattern(t *testing.T) {
	p := NewPattern("")
	fmt.Println(p.Match(""))
}
