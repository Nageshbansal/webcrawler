package api

import (
	"encoding/json"
	"fmt"
	"github.com/Nageshbansal/web-crawler/pkg/telemetry"
	"github.com/Nageshbansal/web-crawler/pkg/types"
	"log"
	"net/http"
)

// ConfigureHandler handles the configure requests
func ConfigureHandler(w http.ResponseWriter, r *http.Request, crawlerMetrics *telemetry.WebCrawlerMetrics) {

	log.Println("ConfigureHandler called")
	// Send success response
	response := types.ConfigureReponse{
		Version:  "0.0.1",
		Response: fmt.Sprintf("Configured successfully"),
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
	return
}
