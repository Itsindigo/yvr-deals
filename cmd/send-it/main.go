package main

import (
	"context"
	"log/slog"

	"github.com/itsindigo/yvr-deals/internal/app_config"
	"github.com/itsindigo/yvr-deals/internal/slack"
)

func main() {
	ctx := context.Background()
	config := app_config.ConfigureApp()
	slog.Debug("Config", slog.String("config", config.String()))

	slackClient := slack.NewSlack(config.Slack.WebhookID)
	dealReporter := slack.NewDealReporter(slackClient)

	dealReporter.ReportDeal(ctx)

	slog.Info("Job finished, exiting.")
}
