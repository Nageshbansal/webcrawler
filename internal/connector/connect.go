package connector

import (
	"fmt"
	"github.com/Nageshbansal/web-crawler/internal/logger"
	"github.com/Nageshbansal/web-crawler/internal/util"
	"github.com/sirupsen/logrus"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// NewWebsiteConnector creates a new WebsiteConnector instance with a default HTTP client.
func NewWebsiteConnector() *WebsiteConnector {
	return &WebsiteConnector{
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

// FetchURL fetches the HTML content of a given URL, ensuring it matches the specified domain.
func (wc *WebsiteConnector) FetchURL(domain string, rawurl string) (*http.Response, error) {

	logger.InfoWithValues("[Connector]: Attempting to fetch URL: %s", logrus.Fields{
		"domain": domain,
		"url":    rawurl,
	})

	// Validate the domain
	if !util.IsSameDomain(rawurl, domain) {
		logger.Errorf("Invalid domain for URL: %s", rawurl)
		return &http.Response{}, fmt.Errorf("invalid domain")
	}

	parsedURL, err := url.Parse(rawurl)
	if err != nil {
		logger.Errorf("[Connector]: Failed to parse URL: %v", err)
		return &http.Response{}, fmt.Errorf("failed to parse URL: %v", err)
	}

	// Ensure the URL scheme is set
	if parsedURL.Scheme == "" {
		parsedURL.Scheme = "http"
	}

	// Ensure the URL path has a leading slash
	if !strings.HasPrefix(parsedURL.Path, "/") {
		parsedURL.Path = "/" + parsedURL.Path
	}

	// Make the request
	req, err := http.NewRequest(http.MethodGet, parsedURL.String(), nil)
	if err != nil {
		logger.Errorf("[Connector]: Failed to create HTTP request for URL %s: %v", rawurl, err)
		return &http.Response{}, fmt.Errorf("failed to create HTTP request: %v", err)
	}

	// Perform the request
	resp, err := wc.client.Do(req)
	if err != nil {
		logger.Errorf("[Connector]: Failed to fetch URL %s: %v", rawurl, err)
		return &http.Response{}, fmt.Errorf("failed to fetch URL: %v", err)
	}

	// Check the status code
	if resp.StatusCode != http.StatusOK {
		logger.Errorf("[Connector]: Unexpected status code: %d for URL %s", resp.StatusCode, rawurl)
		return &http.Response{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	logger.InfoWithValues("[Connector]: Successfully fetched URL", logrus.Fields{"url": rawurl})

	return resp, nil
}
