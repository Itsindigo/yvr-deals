package app_config

import (
	"fmt"
	"log"
	"log/slog"
	"os"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type SlackConfig struct {
	WebhookID string `env:"SLACK_WEBHOOK_ID,required"`
}

type AppConfig struct {
	EnableDebugLogs bool `env:"ENABLE_DEBUG_LOGS"`
	Slack           SlackConfig
}

func ConfigureApp() AppConfig {
	if os.Getenv("IS_REMOTE_ENVIRONMENT") == "" {
		err := godotenv.Load()

		if err != nil {
			log.Fatalf("Error loading .env file: %v", err)
		}
	}

	cfg := AppConfig{}
	err := env.Parse(&cfg)

	if err != nil {
		log.Fatalf("Error mounting config: %v", err)
	}

	if cfg.EnableDebugLogs {
		slog.SetLogLoggerLevel(slog.LevelDebug)
	}

	return cfg
}

func obfuscateSecret(s string, reveal_n int) string {
	if reveal_n > len(s) {
		return "***"
	}
	return fmt.Sprintf("%s***", s[0:reveal_n])
}

func (c AppConfig) String() string {
	return fmt.Sprintf(
		"AppConfig(EnableDebugLogs: %t, Slack(WebhookID: %s))",
		c.EnableDebugLogs,
		obfuscateSecret(c.Slack.WebhookID, 12),
	)
}
