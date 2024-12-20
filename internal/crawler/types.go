package crawler

import (
	"context"
	"github.com/Nageshbansal/web-crawler/pkg/telemetry"
	"os"
	"sync"
)

type ConfigCrawler struct {
	Domain         string
	Visited        sync.Map
	Queue          chan UrlDepth
	File           *os.File
	WaitGroup      sync.WaitGroup
	Context        context.Context
	MaxDepth       int
	crawlerMetrics *telemetry.WebCrawlerMetrics
	lastProcessed  chan int
}

type UrlDepth struct {
	Url   string
	Depth int
}
