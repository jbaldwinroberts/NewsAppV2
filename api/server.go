package api

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/mmcdole/gofeed"

	"NewsAppV2/feed"
	"NewsAppV2/feed/parser"
	"NewsAppV2/feed/store"
)

type Server struct {
}

func (s Server) Serve() {
	router := mux.NewRouter()

	router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		_ = json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	})

	store := store.New()
	parser := parser.New(gofeed.NewParser())
	client := feed.New(store, parser)
	getFeeds := New(client)
	router.PathPrefix("/api/feeds").Methods(http.MethodGet).Handler(getFeeds)

	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
