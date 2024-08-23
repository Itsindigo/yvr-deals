package slack

import (
	"context"
	"fmt"

	"github.com/itsindigo/yvr-deals/internal/feed_reader"
)

type Reporter struct {
	Slack Slack
}

func (dr *Reporter) ReportYVRDeal(ctx context.Context, deal feedreader.Deal) error {
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

func (dr *Reporter) ReportParsingError(ctx context.Context, err error) error {
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

func NewReporter(slack Slack) *Reporter {
	return &Reporter{
		Slack: slack,
	}
}
