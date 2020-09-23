package parser

import (
	"context"
	"os"

	"github.com/pkg/errors"

	"github.com/mmcdole/gofeed"
)

const ParseErrorMessage = "parsing failed"
const RequestTimeOutErrorMessage = "request time out"

type Parser struct {
	fp GofeedParser
}

func New(fp GofeedParser) Parser {
	return Parser{
		fp: fp,
	}
}

func (p Parser) ParseUrl(url string) (*gofeed.Feed, error) {
	return p.ParseUrlWithContext(url, context.Background())
}

func (p Parser) ParseUrlWithContext(url string, ctx context.Context) (*gofeed.Feed, error) {
	feed, err := p.fp.ParseURLWithContext(url, ctx)
	if err != nil {
		if os.IsTimeout(err) {
			return nil, errors.Wrap(err, RequestTimeOutErrorMessage)
		}
		return nil, errors.Wrap(err, ParseErrorMessage)
	}
	return feed, nil
}
