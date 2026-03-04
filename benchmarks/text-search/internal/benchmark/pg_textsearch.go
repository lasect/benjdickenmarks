package benchmark

import (
	"context"

	"github.com/benjdickenmarks/text-search/internal/config"
	"github.com/benjdickenmarks/text-search/internal/metrics"
)

type PgTextsearchBenchmark struct {
	cfg *config.Config
}

func NewPgTextsearchBenchmark(cfg *config.Config) *PgTextsearchBenchmark {
	return &PgTextsearchBenchmark{cfg: cfg}
}

func (b *PgTextsearchBenchmark) Name() string                                         { return "pg_textsearch" }
func (b *PgTextsearchBenchmark) Setup(ctx context.Context) error                      { return nil }
func (b *PgTextsearchBenchmark) LoadData(ctx context.Context, data interface{}) error { return nil }
func (b *PgTextsearchBenchmark) BuildIndex(ctx context.Context) error                 { return nil }
func (b *PgTextsearchBenchmark) Warmup(ctx context.Context, queries []string) error   { return nil }
func (b *PgTextsearchBenchmark) Benchmark(ctx context.Context, queries []string, concurrency int) (metrics.LatencyStats, error) {
	return metrics.LatencyStats{}, nil
}
func (b *PgTextsearchBenchmark) WriteTest(ctx context.Context, data interface{}) (float64, error) {
	return 0, nil
}
func (b *PgTextsearchBenchmark) GetIndexSize() (int64, error)      { return 0, nil }
func (b *PgTextsearchBenchmark) Cleanup(ctx context.Context) error { return nil }
