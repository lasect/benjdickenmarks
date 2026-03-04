package benchmark

import (
	"context"

	"github.com/benjdickenmarks/text-search/internal/config"
	"github.com/benjdickenmarks/text-search/internal/metrics"
)

type MeilisearchBenchmark struct {
	cfg *config.Config
}

func NewMeilisearchBenchmark(cfg *config.Config) *MeilisearchBenchmark {
	return &MeilisearchBenchmark{cfg: cfg}
}

func (b *MeilisearchBenchmark) Name() string                                         { return "meilisearch" }
func (b *MeilisearchBenchmark) Setup(ctx context.Context) error                      { return nil }
func (b *MeilisearchBenchmark) LoadData(ctx context.Context, data interface{}) error { return nil }
func (b *MeilisearchBenchmark) BuildIndex(ctx context.Context) error                 { return nil }
func (b *MeilisearchBenchmark) Warmup(ctx context.Context, queries []string) error   { return nil }
func (b *MeilisearchBenchmark) Benchmark(ctx context.Context, queries []string, concurrency int) (metrics.LatencyStats, error) {
	return metrics.LatencyStats{}, nil
}
func (b *MeilisearchBenchmark) WriteTest(ctx context.Context, data interface{}) (float64, error) {
	return 0, nil
}
func (b *MeilisearchBenchmark) GetIndexSize() (int64, error)      { return 0, nil }
func (b *MeilisearchBenchmark) Cleanup(ctx context.Context) error { return nil }
