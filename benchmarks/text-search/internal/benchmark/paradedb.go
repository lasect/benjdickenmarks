package benchmark

import (
	"context"

	"github.com/benjdickenmarks/text-search/internal/config"
	"github.com/benjdickenmarks/text-search/internal/metrics"
)

type ParadeDBBenchmark struct {
	cfg *config.Config
}

func NewParadeDBBenchmark(cfg *config.Config) *ParadeDBBenchmark {
	return &ParadeDBBenchmark{cfg: cfg}
}

func (b *ParadeDBBenchmark) Name() string                                         { return "paradedb" }
func (b *ParadeDBBenchmark) Setup(ctx context.Context) error                      { return nil }
func (b *ParadeDBBenchmark) LoadData(ctx context.Context, data interface{}) error { return nil }
func (b *ParadeDBBenchmark) BuildIndex(ctx context.Context) error                 { return nil }
func (b *ParadeDBBenchmark) Warmup(ctx context.Context, queries []string) error   { return nil }
func (b *ParadeDBBenchmark) Benchmark(ctx context.Context, queries []string, concurrency int) (metrics.LatencyStats, error) {
	return metrics.LatencyStats{}, nil
}
func (b *ParadeDBBenchmark) WriteTest(ctx context.Context, data interface{}) (float64, error) {
	return 0, nil
}
func (b *ParadeDBBenchmark) GetIndexSize() (int64, error)      { return 0, nil }
func (b *ParadeDBBenchmark) Cleanup(ctx context.Context) error { return nil }
