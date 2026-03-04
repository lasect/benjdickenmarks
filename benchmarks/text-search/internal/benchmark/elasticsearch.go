package benchmark

import (
	"context"

	"github.com/benjdickenmarks/text-search/internal/config"
	"github.com/benjdickenmarks/text-search/internal/metrics"
)

type ElasticsearchBenchmark struct {
	cfg *config.Config
}

func NewElasticsearchBenchmark(cfg *config.Config) *ElasticsearchBenchmark {
	return &ElasticsearchBenchmark{cfg: cfg}
}

func (b *ElasticsearchBenchmark) Name() string                                         { return "elasticsearch" }
func (b *ElasticsearchBenchmark) Setup(ctx context.Context) error                      { return nil }
func (b *ElasticsearchBenchmark) LoadData(ctx context.Context, data interface{}) error { return nil }
func (b *ElasticsearchBenchmark) BuildIndex(ctx context.Context) error                 { return nil }
func (b *ElasticsearchBenchmark) Warmup(ctx context.Context, queries []string) error   { return nil }
func (b *ElasticsearchBenchmark) Benchmark(ctx context.Context, queries []string, concurrency int) (metrics.LatencyStats, error) {
	return metrics.LatencyStats{}, nil
}
func (b *ElasticsearchBenchmark) WriteTest(ctx context.Context, data interface{}) (float64, error) {
	return 0, nil
}
func (b *ElasticsearchBenchmark) GetIndexSize() (int64, error)      { return 0, nil }
func (b *ElasticsearchBenchmark) Cleanup(ctx context.Context) error { return nil }
