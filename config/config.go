package config

import (
	"log"

	"github.com/go-ini/ini"
)

func LoadAPIConfig() (string, string, int) {
	cfg, err := ini.Load("/etc/cai/config.ini")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	apiKey := cfg.Section("openai").Key("api_key").String()
	model := cfg.Section("openai").Key("model").String()
	maxTokens, err := cfg.Section("openai").Key("max_tokens").Int()
	if err != nil {
		maxTokens = 100 // default fallback
	}

	return apiKey, model, maxTokens
}
