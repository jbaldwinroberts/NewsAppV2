package parser

import (
	"context"

	"github.com/mmcdole/gofeed"
)

//go:generate mockgen -destination=../mocks/parser/gofeedparser.go -source=gofeedparser.go
type GofeedParser interface {
	ParseURL(feedURL string) (feed *gofeed.Feed, err error)
	ParseURLWithContext(feedURL string, ctx context.Context) (feed *gofeed.Feed, err error)
}
