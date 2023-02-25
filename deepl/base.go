package deepl

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"

	json "github.com/bytedance/sonic"
	log "github.com/sirupsen/logrus"
)

var Token = ""

func init() {
	log.Info("Deepl API Token: ", Token)
}

type Option func(*url.Values)

// WithNothing 什么都不做
func WithNothing() Option {
	return func(v *url.Values) {}
}

func WithPreserveFormatting() Option {
	return func(v *url.Values) {
		v.Set("preserve_formatting", "1")
	}
}

func WithFormality(formality string) Option {
	return func(v *url.Values) {
		v.Set("formality", formality)
	}
}

func WithHTML() Option {
	return func(v *url.Values) {
		v.Set("tag_handling", "html")
	}
}

func WithIgnoreTags(tags ...string) Option {
	return func(v *url.Values) {
		v.Set("ignore_tags", strings.Join(tags, ","))
	}
}

func Translate(ctx context.Context, text string, targetLang string, options ...Option) (string, error) {
	v := url.Values{}
	v.Set("text", text)
	v.Set("target_lang", targetLang)
	for _, opt := range options {
		opt(&v)
	}
	return translate(ctx, v)
}

func translate(ctx context.Context, v url.Values) (string, error) {
	body := strings.NewReader(v.Encode())
	req, err := http.NewRequestWithContext(ctx, "POST", "https://api.deepl.com/v2/translate", body)
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", "DeepL-Auth-Key "+Token)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return parseTranslateResponse(bytes)
}

type translateResponse struct {
	Translations []struct {
		Text string `json:"text"`
	} `json:"translations"`
}

func parseTranslateResponse(resp []byte) (string, error) {
	tr := &translateResponse{}
	if err := json.Unmarshal(resp, tr); err != nil {
		return "", err
	}
	return tr.Translations[0].Text, nil
}

func IsValidLang(lang string) bool {
	switch lang {
	case "BG", "CS", "DA", "DE", "EL", "EN", "ES", "ET", "FI", "FR", "HU", "ID", "IT", "JA", "KO", "LT", "LV", "NB", "NL", "PL", "PT", "RO", "RU", "SK", "SL", "SV", "TR", "UK", "ZH":
		return true
	default:
		return false
	}
}
