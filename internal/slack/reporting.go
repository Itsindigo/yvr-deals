package slack

import (
	"context"
	"fmt"

	"github.com/itsindigo/yvr-deals/internal/feed_reader"
)

type DealReporter struct {
	Slack Slack
}

func (dr *DealReporter) ReportDeal(ctx context.Context, deal feedreader.Deal) error {
	message, err := GetDealMessageBlocks(deal)

	if err != nil {
		return fmt.Errorf("could not create deal message: %w", err)
	}

	_, err = dr.Slack.SendMessage(ctx, message)

	if err != nil {
		return err
	}

	return nil
}

func (dr *DealReporter) ReportParsingError(ctx context.Context, err error) error {
	message, err := GetErrorMessageBlocks(err.Error())

	if err != nil {
		return fmt.Errorf("could not create no action message: %w", err)
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
