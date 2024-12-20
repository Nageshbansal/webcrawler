package connector

import (
	"net/http"
)

// WebsiteConnector represents a client for fetching HTML content from URLs.
type WebsiteConnector struct {
	client *http.Client
}
