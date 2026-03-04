package benchmark

import (
	"context"

	"github.com/benjdickenmarks/text-search/internal/config"
	"github.com/benjdickenmarks/text-search/internal/metrics"
)

type PgroongaBenchmark struct {
	cfg *config.Config
}

func NewPgroongaBenchmark(cfg *config.Config) *PgroongaBenchmark {
	return &PgroongaBenchmark{cfg: cfg}
}

func (b *PgroongaBenchmark) Name() string                                         { return "pgroonga" }
func (b *PgroongaBenchmark) Setup(ctx context.Context) error                      { return nil }
func (b *PgroongaBenchmark) LoadData(ctx context.Context, data interface{}) error { return nil }
func (b *PgroongaBenchmark) BuildIndex(ctx context.Context) error                 { return nil }
func (b *PgroongaBenchmark) Warmup(ctx context.Context, queries []string) error   { return nil }
func (b *PgroongaBenchmark) Benchmark(ctx context.Context, queries []string, concurrency int) (metrics.LatencyStats, error) {
	return metrics.LatencyStats{}, nil
}
func (b *PgroongaBenchmark) WriteTest(ctx context.Context, data interface{}) (float64, error) {
	return 0, nil
}
func (b *PgroongaBenchmark) GetIndexSize() (int64, error)      { return 0, nil }
func (b *PgroongaBenchmark) Cleanup(ctx context.Context) error { return nil }
