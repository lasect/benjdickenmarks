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

	if err := runner.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Benchmarks complete!")
}
