package results

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"

	"github.com/benjdickenmarks/text-search/internal/metrics"
)

type BenchmarkResult struct {
	Tool        string
	Dataset     string
	Size        int
	Concurrency int

	IndexTime time.Duration
	IndexSize int64

	QueryStats metrics.LatencyStats

	WriteTPS float64
}

func WriteResults(results []BenchmarkResult, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	header := []string{
		"tool", "dataset", "size", "concurrency",
		"index_time_ms", "index_size_bytes",
		"query_min_ms", "query_max_ms", "query_mean_ms", "query_median_ms",
		"query_p95_ms", "query_p99_ms", "query_stddev_ms",
		"query_total_ms", "query_count", "query_success", "query_failed", "query_tps",
		"write_tps",
	}
	writer.Write(header)

	for _, r := range results {
		row := []string{
			r.Tool,
			r.Dataset,
			fmt.Sprintf("%d", r.Size),
			fmt.Sprintf("%d", r.Concurrency),
			fmt.Sprintf("%.2f", float64(r.IndexTime.Milliseconds())),
			fmt.Sprintf("%d", r.IndexSize),
			fmt.Sprintf("%.2f", float64(r.QueryStats.Min.Milliseconds())),
			fmt.Sprintf("%.2f", float64(r.QueryStats.Max.Milliseconds())),
			fmt.Sprintf("%.2f", float64(r.QueryStats.Mean.Milliseconds())),
			fmt.Sprintf("%.2f", float64(r.QueryStats.Median.Milliseconds())),
			fmt.Sprintf("%.2f", float64(r.QueryStats.P95.Milliseconds())),
			fmt.Sprintf("%.2f", float64(r.QueryStats.P99.Milliseconds())),
			fmt.Sprintf("%.2f", float64(r.QueryStats.StdDev.Milliseconds())),
			fmt.Sprintf("%.2f", float64(r.QueryStats.Total.Milliseconds())),
			fmt.Sprintf("%d", r.QueryStats.Count),
			fmt.Sprintf("%d", r.QueryStats.Success),
			fmt.Sprintf("%d", r.QueryStats.Failed),
			fmt.Sprintf("%.2f", r.QueryStats.TPS),
			fmt.Sprintf("%.2f", r.WriteTPS),
		}
		writer.Write(row)
	}

	return writer.Error()
}
