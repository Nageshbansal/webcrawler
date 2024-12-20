package api

import (
	"encoding/json"
	"fmt"
	"github.com/Nageshbansal/web-crawler/internal/logger"
	"github.com/Nageshbansal/web-crawler/internal/sitemap"
	"github.com/Nageshbansal/web-crawler/pkg/telemetry"
	"github.com/Nageshbansal/web-crawler/pkg/types"
	"net/http"
	"os"
)

func SitemapHandler(w http.ResponseWriter, r *http.Request, crawlerMetrics *telemetry.WebCrawlerMetrics) {

	logger.Info("[SitemapHandler]: SitemapHandler started")
	var reqBody types.SiteMapRequest

	// Decode JSON request body
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Validate ID (basic check)
	if reqBody.ID == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}
	// Parse URLs from the file
	pwd, _ := os.Getwd()
	filepath := fmt.Sprintf("%s/%s.txt", pwd, reqBody.ID)
	urls, err := sitemap.ParseURLsFromFile(filepath)
	if err != nil {
		logger.Fatalf("[Sitemap]: Error parsing URLs from file: %v", err)
	}

	// Generate and print sitemap
	sm := sitemap.NewSitemap(urls)

	// Send success response
	response := types.SiteMapResponse{
		Message: fmt.Sprintf("Sitemap generated for ID: %s", reqBody.ID),
		Data:    sm.String(),
		// Dummy crawl ID
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
	return
}
