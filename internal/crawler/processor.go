package crawler

import (
	"github.com/Nageshbansal/web-crawler/internal/logger"
	"github.com/Nageshbansal/web-crawler/internal/reader"
	"github.com/Nageshbansal/web-crawler/internal/util"
	"time"
)

// Run initiates the crawling process, listening for URLs to crawl until the context is cancelled.
func (c *ConfigCrawler) Run() {
	for {
		select {
		case <-c.Context.Done():
			// Handle context cancellation gracefully
			logger.Info("[Processor]: Crawler context cancelled")
			return
		case ud := <-c.Queue:
			if ud.Depth <= c.MaxDepth {
				c.WaitGroup.Add(1)
				go c.crawl(ud.Url, ud.Depth)
			}
		}
	}
}

// crawl processes a single URL, fetching links and enqueuing them for further crawling.
func (c *ConfigCrawler) crawl(currentURL string, depth int) {
	defer c.WaitGroup.Done()
	startTime := time.Now()
	links, err := reader.Read(c.Domain, currentURL)
	duration := time.Since(startTime)
	c.crawlerMetrics.PageFetchLatency.Record(
		c.Context,
		duration.Milliseconds(),
	)
	if err != nil {
		c.crawlerMetrics.ErrorRequestCounter.Add(c.Context, 1)
		c.crawlerMetrics.CrawledFailedPageCounter.Add(c.Context, 1)
		logger.Errorf("[Processor]: Error reading URL %s: %v", currentURL, err)
		return
	}
	c.crawlerMetrics.CrawledSuccessPagesCounter.Add(c.Context, 1)
	c.crawlerMetrics.TotalPagesCounter.Add(c.Context, int64(len(links)))
	for _, link := range links {
		absoluteURL := util.ResolveURL(link, currentURL)
		if util.IsSameDomain(absoluteURL, c.Domain) {
			if _, loaded := c.Visited.LoadOrStore(absoluteURL, true); !loaded {
				logger.Infof("[Processor]: Found new link: %s", absoluteURL)
				c.Queue <- UrlDepth{Url: absoluteURL, Depth: depth + 1}
				c.crawlerMetrics.URLQueueSizeGauge.Record(c.Context, int64(len(c.Queue)))
				if _, err := c.File.WriteString(absoluteURL + "\n"); err != nil {
					logger.Errorf("[Processor]: Error writing URL %s to file: %v", absoluteURL, err)
				}
			}
		}
	}
}
