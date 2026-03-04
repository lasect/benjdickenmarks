package benchmark

import (
	"context"

	"github.com/benjdickenmarks/text-search/internal/config"
	"github.com/benjdickenmarks/text-search/internal/metrics"
)

type PgTrgmBenchmark struct {
	cfg *config.Config
}

func NewPgTrgmBenchmark(cfg *config.Config) *PgTrgmBenchmark {
	return &PgTrgmBenchmark{cfg: cfg}
}

func (b *PgTrgmBenchmark) Name() string                                         { return "pg_trgm" }
func (b *PgTrgmBenchmark) Setup(ctx context.Context) error                      { return nil }
func (b *PgTrgmBenchmark) LoadData(ctx context.Context, data interface{}) error { return nil }
func (b *PgTrgmBenchmark) BuildIndex(ctx context.Context) error                 { return nil }
func (b *PgTrgmBenchmark) Warmup(ctx context.Context, queries []string) error   { return nil }
func (b *PgTrgmBenchmark) Benchmark(ctx context.Context, queries []string, concurrency int) (metrics.LatencyStats, error) {
	return metrics.LatencyStats{}, nil
}
func (b *PgTrgmBenchmark) WriteTest(ctx context.Context, data interface{}) (float64, error) {
	return 0, nil
}
func (b *PgTrgmBenchmark) GetIndexSize() (int64, error)      { return 0, nil }
func (b *PgTrgmBenchmark) Cleanup(ctx context.Context) error { return nil }
