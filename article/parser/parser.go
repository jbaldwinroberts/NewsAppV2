package parser

import (
	"github.com/pkg/errors"

	"github.com/mmcdole/gofeed"
)

const ParseErrorMessage = "parsing failed"

type Parser struct {
	fp GofeedParser
}

func New(fp GofeedParser) Parser {
	return Parser{
		fp: fp,
	}
}

func (p Parser) ParseUrl(url string) (*gofeed.Feed, error) {
	feed, err := p.fp.ParseURL(url)
	if err != nil {
		return nil, errors.Wrap(err, ParseErrorMessage)
	}

	return feed, nil
}
