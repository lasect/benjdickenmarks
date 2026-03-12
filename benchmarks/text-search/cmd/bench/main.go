package main

import (
	"fmt"
	"os"

	"github.com/benjdickenmarks/text-search/internal/benchmark"
	"github.com/benjdickenmarks/text-search/internal/config"
)

func main() {
	cfg := config.Load()

	runner := benchmark.NewRunner(cfg)
	runner.Register(benchmark.NewTSVectorBenchmark(cfg))
	runner.Register(benchmark.NewPgTrgmBenchmark(cfg))
	runner.Register(benchmark.NewPgroongaBenchmark(cfg))
	runner.Register(benchmark.NewParadeDBBenchmark(cfg))
	runner.Register(benchmark.NewPgTextsearchBenchmark(cfg))
	runner.Register(benchmark.NewElasticsearchBenchmark(cfg))
	runner.Register(benchmark.NewMeilisearchBenchmark(cfg))

	fmt.Printf("Running benchmarks with dataset: %s, size: %d\n",
		cfg.Benchmark.Dataset, cfg.Benchmark.DatasetSize)

	if err := runner.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Benchmarks complete!")
	fmt.Println("")
	fmt.Println("Note: To use real Wikipedia data:")
	fmt.Println("  pip install pandas pyarrow")
	fmt.Println("  python3 -c \"import pandas as pd; df = pd.read_parquet('data/wikipedia.parquet'); df.to_json('data/wikipedia.json', orient='records')\"")
}
