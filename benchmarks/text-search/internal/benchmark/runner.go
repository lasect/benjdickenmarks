package benchmark

import (
	"context"
	"fmt"
	"time"

	"github.com/benjdickenmarks/text-search/internal/config"
	"github.com/benjdickenmarks/text-search/internal/metrics"
	"github.com/benjdickenmarks/text-search/internal/results"
)

type Benchmarker interface {
	Name() string
	Setup(ctx context.Context) error
	LoadData(ctx context.Context, data interface{}) error
	BuildIndex(ctx context.Context) error
	Warmup(ctx context.Context, queries []string) error
	Benchmark(ctx context.Context, queries []string, concurrency int) (metrics.LatencyStats, error)
	WriteTest(ctx context.Context, data interface{}) (float64, error)
	Cleanup(ctx context.Context) error
	GetIndexSize() (int64, error)
}

type Runner struct {
	cfg        *config.Config
	benchmarks []Benchmarker
	results    []results.BenchmarkResult
}

func NewRunner(cfg *config.Config) *Runner {
	return &Runner{
		cfg:        cfg,
		benchmarks: make([]Benchmarker, 0),
		results:    make([]results.BenchmarkResult, 0),
	}
}

func (r *Runner) Register(b Benchmarker) {
	r.benchmarks = append(r.benchmarks, b)
}

func (r *Runner) Run() error {
	queries := r.loadQueries()

	for _, b := range r.benchmarks {
		fmt.Printf("\n=== Running benchmark: %s ===\n", b.Name())

		ctx := context.Background()

		if err := b.Setup(ctx); err != nil {
			fmt.Printf("Setup failed: %v\n", err)
			continue
		}
		defer b.Cleanup(ctx)

		data := r.generateData()
		fmt.Printf("Loading data...\n")
		loadStart := time.Now()
		if err := b.LoadData(ctx, data); err != nil {
			fmt.Printf("Load failed: %v\n", err)
			continue
		}
		loadTime := time.Since(loadStart)
		fmt.Printf("Data loaded in %v\n", loadTime)

		fmt.Printf("Building index...\n")
		indexStart := time.Now()
		if err := b.BuildIndex(ctx); err != nil {
			fmt.Printf("Index build failed: %v\n", err)
			continue
		}
		indexTime := time.Since(indexStart)
		fmt.Printf("Index built in %v\n", indexTime)

		indexSize, _ := b.GetIndexSize()
		fmt.Printf("Index size: %d bytes\n", indexSize)

		fmt.Printf("Warming up (%d queries)...\n", r.cfg.Benchmark.WarmupQueries)
		warmupQueries := queries[:r.cfg.Benchmark.WarmupQueries]
		if err := b.Warmup(ctx, warmupQueries); err != nil {
			fmt.Printf("Warmup failed: %v\n", err)
		}

		for _, conc := range r.cfg.Benchmark.Concurrency {
			fmt.Printf("Benchmarking with concurrency %d...\n", conc)
			stats, err := b.Benchmark(ctx, queries, conc)
			if err != nil {
				fmt.Printf("Benchmark failed: %v\n", err)
				continue
			}

			result := results.BenchmarkResult{
				Tool:        b.Name(),
				Dataset:     r.cfg.Benchmark.Dataset,
				Size:        r.cfg.Benchmark.DatasetSize,
				Concurrency: conc,
				IndexTime:   indexTime,
				IndexSize:   indexSize,
				QueryStats:  stats,
			}
			r.results = append(r.results, result)

			fmt.Printf("  p50: %v, p95: %v, p99: %v, tps: %.2f\n",
				stats.Median, stats.P95, stats.P99, stats.TPS)
		}

		fmt.Printf("Running write test...\n")
		writeTPS, _ := b.WriteTest(ctx, data)
		if len(r.results) > 0 {
			r.results[len(r.results)-1].WriteTPS = writeTPS
		}
		fmt.Printf("  Write TPS: %.2f\n", writeTPS)
	}

	return r.saveResults()
}

func (r *Runner) loadQueries() []string {
	return []string{
		"premium wireless headphones",
		"how to optimize postgresql queries",
		"history of artificial intelligence",
		"best way to handle errors in go",
		"understanding quantum mechanics",
		"modern javascript frameworks",
		"python machine learning",
		"docker container best practices",
		"rest api design patterns",
		"data structures algorithms",
	}
}

func (r *Runner) generateData() interface{} {
	return nil
}

func (r *Runner) saveResults() error {
	filename := fmt.Sprintf("results_%s.csv", time.Now().Format("20060102_150405"))
	return results.WriteResults(r.results, filename)
}
