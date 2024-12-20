package reader

import (
	"fmt"
	"github.com/Nageshbansal/web-crawler/internal/connector"
	"github.com/Nageshbansal/web-crawler/internal/extractor"
	"github.com/Nageshbansal/web-crawler/internal/logger"
	"github.com/sirupsen/logrus"
	"net/http"
)

// Read fetches the HTML content from the specified URL and extracts all links from it.
// It returns a slice of strings containing the extracted links and any error encountered.
func Read(domain string, urlStr string) ([]string, error) {
	logger.InfoWithValues("[Reader]: Reading URL", logrus.Fields{
		"domain": domain,
		"url":    urlStr,
	})

	wc := connector.NewWebsiteConnector()
	response, err := wc.FetchURL(domain, urlStr)

	if err != nil {
		logger.Errorf("[Reader]: Error fetching URL %s: %v", urlStr, err)
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		logger.Errorf("[Reader]: Non-OK HTTP status: %d for URL %s", response.StatusCode, urlStr)
		return nil, fmt.Errorf("non-OK HTTP status: %d", response.StatusCode)
	}

	links, err := extractor.ExtractLinks(response.Body)
	if err != nil {
		logger.Errorf("[Reader]: Error extracting links from URL %s: %v", urlStr, err)
		return nil, err
	}
	return links, nil
}
