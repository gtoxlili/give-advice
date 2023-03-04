package constants

import (
	"context"

	log "github.com/sirupsen/logrus"
)

var (
	DeeplToken    = ""
	OpenAIToken   = ""
	RedisAddr     = ""
	RedisPassword = ""
	ProxyAddr     = ""
)

func init() {
	log.Info("Deepl API Token: ", DeeplToken)
	log.Info("OpenAI API Token: ", OpenAIToken)
	log.Info("Redis Addr: ", RedisAddr)
	log.Info("Redis Password: ", RedisPassword)
	log.Info("Proxy Addr: ", ProxyAddr)
}

func CalcDeeplToken(ctx context.Context) string {
	if token, ok := ctx.Value("Deepl-Auth-Key").(string); ok && token != "" {
		return token
	}
	return DeeplToken
}

func CalcOpenAIToken(ctx context.Context) string {
	if token, ok := ctx.Value("OpenAI-Auth-Key").(string); ok && token != "" {
		return token
	}
	return OpenAIToken
}
