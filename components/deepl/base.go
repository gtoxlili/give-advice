package deepl

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/gtoxlili/give-advice/common/ht"
	"github.com/gtoxlili/give-advice/constants"
)

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

type tranRes struct {
	Translations []struct {
		Text string `json:"text"`
	} `json:"translations"`
}

func translate(ctx context.Context, v url.Values) (string, error) {
	res, err := ht.Post[tranRes](ctx,
		"https://api.deepl.com/v2/translate", v,
		ht.H{
			"Authorization": "DeepL-Auth-Key " + constants.CalcDeeplToken(ctx),
		})
	if err != nil {
		return "", err
	}
	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("[Translate] anomalies : %s", res.Status)
	}
	return res.Data.Translations[0].Text, nil
}

func IsValidLang(lang string) bool {
	switch lang {
	case "BG", "CS", "DA", "DE", "EL", "EN", "ES", "ET", "FI", "FR", "HU", "ID", "IT", "JA", "KO", "LT", "LV", "NB", "NL", "PL", "PT", "RO", "RU", "SK", "SL", "SV", "TR", "UK", "ZH":
		return true
	default:
		return false
	}
}
