package main

import (
	"context"
	"log/slog"

	"github.com/itsindigo/yvr-deals/internal/app_config"
	"github.com/itsindigo/yvr-deals/internal/feed_reader"
	"github.com/itsindigo/yvr-deals/internal/slack"
)

func main() {
	ctx := context.Background()
	config := app_config.ConfigureApp()
	slog.Debug("Config", slog.String("config", config.String()))

	slackClient := slack.NewSlack(config.Slack.WebhookID)
	reporter := slack.NewReporter(slackClient)
	yvr := feedreader.NewYvrHandler()

	deals, err := yvr.GetPastNDayDeals(ctx, 1)

	if err != nil {
		slog.Error("Error getting deals", slog.String("error", err.Error()))
		reporter.ReportParsingError(ctx, err)
		return
	}

	for _, d := range deals {
		slog.Info("Found Deal", slog.String("title", d.Title))
		reporter.ReportYVRDeal(ctx, d)
	}

	slog.Info("Job finished, exiting.")
}
