// Package crawler provides a simple web crawler that traverses websites up to a specified depth.
package crawler

import (
	"context"
	"fmt"
	"github.com/Nageshbansal/web-crawler/internal/logger"
	"github.com/Nageshbansal/web-crawler/pkg/telemetry"
	"net/url"
	"os"
	"sync"
	"time"
)

// Crawl starts the crawling process from a given baseURL and logs the visited URLs to a file.
// It takes a baseURL and an id to create a unique output file.
// Errors during the process are logged, and the function returns an error if the initial setup fails.
func Crawl(baseURL string, id string, maxDepth int, crawlerMetrics *telemetry.WebCrawlerMetrics) error {

	u, err := url.ParseRequestURI(baseURL)
	if err != nil {
		//crawlerMetrics.ErrorRequestCounter.Add(ctx, 1)
		logger.Errorf("[Crawler]: Invalid start URL: %v\n", err)
		return fmt.Errorf("invalid start URL: %v", err)
	}
	// redhat.com
	domain := u.String()

	file, err := os.Create(fmt.Sprintf("%s.txt", id))
	if err != nil {
		logger.Errorf("[Crawler]: Error creating output file: %v\n", err)
		return err // Return error for handling
	}
	defer file.Close()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	crawlJob := ConfigCrawler{
		Domain:         domain,
		Visited:        sync.Map{},
		Queue:          make(chan UrlDepth, 100),
		File:           file,
		WaitGroup:      sync.WaitGroup{},
		Context:        ctx,
		MaxDepth:       maxDepth,
		crawlerMetrics: crawlerMetrics,
	}

	crawlJob.WaitGroup.Add(1)
	go func() {
		defer crawlJob.WaitGroup.Done()
		crawlJob.Run()
	}()

	crawlJob.Queue <- UrlDepth{Url: baseURL, Depth: 0}
	time.Sleep(1 * time.Second)

	cancel()
	crawlJob.WaitGroup.Wait()

	close(crawlJob.Queue)
	time.Sleep(1 * time.Second)
	logger.Info("[Crawler]: Crawling completed")
	return nil
}
