// pkg/api/router.go

package api

import (
	"context"
	"github.com/Nageshbansal/web-crawler/pkg/telemetry"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// NewRouter creates a new HTTP router
func NewRouter() *mux.Router {

	ctx := context.Background()

	meter, err := telemetry.NewMeter(ctx)
	if err != nil {
		log.Fatal(err)
	}

	crawlerMetrics, err := telemetry.IntializeMetrics(meter)
	if err != nil {
		log.Fatal(err)
	}

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/crawl", func(w http.ResponseWriter, r *http.Request) {
		log.Print("Crawler called")
		CrawlHandler(w, r, crawlerMetrics)
	}).Methods("POST")

	router.HandleFunc("/sitemap", func(w http.ResponseWriter, r *http.Request) {
		SitemapHandler(w, r, crawlerMetrics)
	}).Methods("POST")

	router.HandleFunc("/configure", func(w http.ResponseWriter, r *http.Request) {
		//crawlerMetrics.RequestLatencyHistogram.Record(ctx, 3)
		ConfigureHandler(w, r, crawlerMetrics)
	}).Methods("POST")
	return router
}
