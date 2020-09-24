package api

import (
	"encoding/json"
	"net/http"

	"github.com/mmcdole/gofeed"
)

type Getter interface {
	GetFeeds() ([]*gofeed.Feed, error)
}

type GetFeedsHandler struct {
	g Getter
}

func New(getter Getter) GetFeedsHandler {
	return GetFeedsHandler{
		g: getter,
	}
}

func (h GetFeedsHandler) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	feeds, err := h.g.GetFeeds()
	if err != nil {
		//TODO
	}
	_ = json.NewEncoder(w).Encode(feeds)
}
