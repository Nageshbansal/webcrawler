package types

// CrawlRequest represents the structure of JSON request body for crawl request
type CrawlRequest struct {
	URL      string `json:"url"`
	MaxDepth int    `json:"max_depth" default:"5"`
}

// CrawlResponse represents the structure of JSON response for crawl request
type CrawlResponse struct {
	Message string `json:"message"`
	CrawlID string `json:"crawl_id"`
	Error   string `json:"error,omitempty"`
}
