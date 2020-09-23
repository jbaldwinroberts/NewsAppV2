package feed

import (
	"github.com/mmcdole/gofeed"
	"github.com/pkg/errors"
)

const GetFeedsFailedErrorMessage = "get feeds failed"

//go:generate mockgen -destination=mocks/feed.go -source=feed.go
type Parser interface {
	ParseUrl(url string) (*gofeed.Feed, error)
	//ParseUrlWithContext(url string, ctx context.Context) (*gofeed.Feed, error)
}

type Getter interface {
	GetUrls() []string
}

type Client struct {
	g Getter
	p Parser
}

func New(getter Getter, parser Parser) Client {
	return Client{
		g: getter,
		p: parser,
	}
}

func (c Client) GetFeeds() ([]*gofeed.Feed, error) {
	var feeds []*gofeed.Feed

	for _, v := range c.g.GetUrls() {
		feed, err := c.p.ParseUrl(v)
		if err != nil {
			return nil, errors.Wrap(err, GetFeedsFailedErrorMessage)
		}

		feeds = append(feeds, feed)
	}

	return feeds, nil
}
