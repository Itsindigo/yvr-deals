package feedreader

import (
	"context"
	"fmt"
	"time"
)

const YVR_DEALS_RSS_URL = "https://ydeals.com/output/rss/yvr.xml"

type YvrHandler struct {
	FeedReader FeedReader
}

type Deal struct {
	Title       string
	Link        string
	PublishedAt time.Time
}

func (yvr *YvrHandler) GetPastNDayDeals(ctx context.Context, n int) ([]Deal, error) {
	feed, err := yvr.FeedReader.ByURL(ctx, YVR_DEALS_RSS_URL)

	if err != nil {
		return []Deal{}, fmt.Errorf("error GetNDaysDealls: %w", err)
	}

	deals := make([]Deal, 0)
	until := time.Now().Add(-24 * time.Hour * time.Duration(n))

	for _, item := range feed.Items {
		deal := Deal{
			Title:       item.Title,
			Link:        item.Link,
			PublishedAt: *item.PublishedParsed,
		}

		if deal.PublishedAt.Before(until) {
			break
		}

		deals = append(deals, deal)
	}

	return deals, nil
}

func NewYvrHandler() YvrHandler {
	return YvrHandler{
		FeedReader: NewFeedReader(),
	}
}
