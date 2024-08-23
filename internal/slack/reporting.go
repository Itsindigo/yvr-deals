package slack

import (
	"context"
	"fmt"
)

type DealReporter struct {
	Slack Slack
}

func (dr *DealReporter) ReportDeal(ctx context.Context) error {
	message, err := GetDealInfoBlocks()

	if err != nil {
		return fmt.Errorf("could not create deal message: %w", err)
	}

	_, err = dr.Slack.SendMessage(ctx, message)

	if err != nil {
		return err
	}

	return nil
}

func NewDealReporter(slack Slack) *DealReporter {
	return &DealReporter{
		Slack: slack,
	}
}
