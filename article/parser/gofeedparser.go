package parser

import "github.com/mmcdole/gofeed"

//go:generate mockgen -destination=../mocks/gofeedparser.go -source=gofeedparser.go
type GofeedParser interface {
	ParseURL(feedURL string) (feed *gofeed.Feed, err error)
}
