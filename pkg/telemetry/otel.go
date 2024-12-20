package telemetry

import (
	"context"
	"fmt"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/metric"
	metricSDK "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
	"os"
	"time"
)

type WebCrawlerMetrics struct {
	// Existing metrics
	ErrorRequestCounter        metric.Int64Counter
	RequestLatencyHistogram    metric.Float64Histogram
	CrawledSuccessPagesCounter metric.Int64Counter
	CrawledFailedPageCounter   metric.Int64Counter
	TotalPagesCounter          metric.Int64Counter
	// New Gauge for tracking URL queue size
	URLQueueSizeGauge    metric.Int64Gauge
	PageFetchLatency     metric.Int64Histogram
	StuckCrawlerDetected metric.Int64Counter
}

func IntializeMetrics(meter metric.Meter) (*WebCrawlerMetrics, error) {

	errorRequestCounter, err := meter.Int64Counter("webcrawler.error_requests_total",
		metric.WithDescription("Total number of error request during crawling"),
		metric.WithUnit("{request}"))
	if err != nil {
		return nil, err
	}
	// Request latency histogram
	requestLatencyHistogram, err := meter.Float64Histogram("webcrawler_request_latency_seconds",
		metric.WithDescription("Distribution of request latencies"),
		metric.WithUnit("s"))
	if err != nil {
		return nil, err
	}

	// Crawled pages counter
	crawledPagesCounter, err := meter.Int64Counter("webcrawler_crawled_pages_success",
		metric.WithDescription("Total number of pages crawled"),
		metric.WithUnit("{page}"))
	if err != nil {
		return nil, err
	}

	// Active crawlers gauge
	urlQueueSizeGauge, err := meter.Int64Gauge("webcrawler_url_queue",
		metric.WithDescription("Crawler Queue Size"),
		metric.WithUnit("{pages}"))
	if err != nil {
		return nil, err
	}

	crawledFailedPageCounter, err := meter.Int64Counter("webcrawler_crawled_failed_page",
		metric.WithDescription("Crawler Failed Page"),
		metric.WithUnit("{page}"))
	if err != nil {
		return nil, err
	}

	totalPagesCounter, err := meter.Int64Counter("webcrawler_total_pages",
		metric.WithDescription("Total number of pages"),
		metric.WithUnit("{page}"))
	if err != nil {
		return nil, err
	}
	pageFetchLatency, err := meter.Int64Histogram("webcrawler_page_fetch_latency",
		metric.WithDescription("Distribution of page fetch latency"),
		metric.WithUnit("ms"))
	if err != nil {
		return nil, err
	}

	stuckCrawlerDetected, err := meter.Int64Counter("webcrawler_stuck_crawlers",
		metric.WithDescription("Number of instances where the crawler was stuck or unresponsive."),
		metric.WithUnit("{instance}"))

	return &WebCrawlerMetrics{
		ErrorRequestCounter:        errorRequestCounter,
		RequestLatencyHistogram:    requestLatencyHistogram,
		CrawledSuccessPagesCounter: crawledPagesCounter,
		URLQueueSizeGauge:          urlQueueSizeGauge,
		CrawledFailedPageCounter:   crawledFailedPageCounter,
		TotalPagesCounter:          totalPagesCounter,
		PageFetchLatency:           pageFetchLatency,
		StuckCrawlerDetected:       stuckCrawlerDetected,
	}, nil
}

func NewMeter(ctx context.Context) (metric.Meter, error) {
	provider, err := newMeterProvider(ctx)
	if err != nil {
		return nil, err
	}
	return provider.Meter("webcrawler"), nil
}

func newMeterProvider(ctx context.Context) (metric.MeterProvider, error) {

	metricExporter, err := getOtelMetricsCollectorExporter(ctx)
	if err != nil {
		return nil, err
	}
	res, err := newResource()
	if err != nil {
		return nil, err
	}

	meterProvider := metricSDK.NewMeterProvider(
		metricSDK.WithResource(res),
		metricSDK.WithReader(metricSDK.NewPeriodicReader(metricExporter,
			metricSDK.WithInterval(3*time.Second))))

	return meterProvider, nil

}

func newResource() (*resource.Resource, error) {
	return resource.Merge(resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName("webcrawler"),
			semconv.ServiceVersion("1.0.0")))
}

func getOtelMetricsCollectorExporter(ctx context.Context) (metricSDK.Exporter, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	OTEL_EXPORTER_OTLP_ENDPOINT := os.Getenv("OTEL_EXPORTER_OTLP_ENDPOINT")

	fmt.Printf("Using OTLP Endpoint: %s\n", OTEL_EXPORTER_OTLP_ENDPOINT)
	if OTEL_EXPORTER_OTLP_ENDPOINT == "" {
		OTEL_EXPORTER_OTLP_ENDPOINT = "0.0.0.0:4317"
	}
	exporter, err := otlpmetricgrpc.New(ctx,
		otlpmetricgrpc.WithEndpoint(OTEL_EXPORTER_OTLP_ENDPOINT),
		otlpmetricgrpc.WithInsecure(),
	)
	//exporter, err := stdoutmetric.New()
	if err != nil {
		return nil, fmt.Errorf("could not create metric exporter: %w", err)
	}

	return exporter, nil
}
