package api

import (
	"encoding/json"
	"fmt"
	"github.com/Nageshbansal/web-crawler/internal/crawler"
	"github.com/Nageshbansal/web-crawler/internal/logger"
	"github.com/Nageshbansal/web-crawler/pkg/telemetry"
	"github.com/Nageshbansal/web-crawler/pkg/types"
	"github.com/google/uuid"
	"net/http"
)

// CrawlHandler handles the crawl requests
func CrawlHandler(w http.ResponseWriter, r *http.Request, crawlerMetrics *telemetry.WebCrawlerMetrics) {

	logger.Info("[CrawlHandler]: CrawlHandler started")
	var reqBody types.CrawlRequest
	ctx := r.Context()
	// Decode JSON request body
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Validate URL (basic check)
	if reqBody.URL == "" {
		http.Error(w, "URL is required", http.StatusBadRequest)
		return
	}

	// Initialize crawler and start crawling
	id := uuid.New().String()
	url := reqBody.URL
	maxDepth := reqBody.MaxDepth

	err = crawler.Crawl(url, id, maxDepth, crawlerMetrics)
	if err != nil {
		crawlerMetrics.ErrorRequestCounter.Add(ctx, 1)
		http.Error(w, fmt.Sprintf("Error crawling URL: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	// Send success response
	response := types.CrawlResponse{
		Message: fmt.Sprintf("Crawling Done for URL: %s", reqBody.URL),
		CrawlID: id,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
	return
}
