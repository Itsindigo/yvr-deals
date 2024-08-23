package feedreader

import (
	"context"
	"fmt"

	"github.com/mmcdole/gofeed"
)

type FeedReader struct{}

func (fr *FeedReader) ByURL(ctx context.Context, url string) (*gofeed.Feed, error) {
	fp := gofeed.NewParser()
	feed, err := fp.ParseURLWithContext(url, ctx)

	if err != nil {
		return &gofeed.Feed{}, fmt.Errorf("error fetching feed from %s: %w", url, err)
	}

	return feed, nil
}

func NewFeedReader() FeedReader {
	return FeedReader{}
}
