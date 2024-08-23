package slack

import (
	"fmt"

	feedreader "github.com/itsindigo/yvr-deals/internal/feed_reader"
)

func GetDealMessageBlocks(deal feedreader.Deal) (Blocks, error) {
	message := fmt.Sprintf("*YVR Deals:*\n*%s*\n%s", deal.Title, deal.Link)
	blocks := []SectionBlock{
		{Type: "section", Text: TextBlock{Type: "mrkdwn", Text: message}},
	}

	return NewBlocksMap(blocks)
}

func GetErrorMessageBlocks(error string) (Blocks, error) {
	message := fmt.Sprintf("*Error occurred while fetching data:* %s", error)
	blocks := []SectionBlock{
		{Type: "section", Text: TextBlock{Type: "mrkdwn", Text: message}},
	}

	return NewBlocksMap(blocks)
}
