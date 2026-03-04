package metrics

import (
	"math"
	"sort"
	"time"
)

type QueryResult struct {
	Latency time.Duration
	Success bool
}

type LatencyStats struct {
	Min     time.Duration
	Max     time.Duration
	Mean    time.Duration
	Median  time.Duration
	P95     time.Duration
	P99     time.Duration
	StdDev  time.Duration
	Total   time.Duration
	Count   int
	Success int
	Failed  int
	TPS     float64
}

func CalculateStats(results []QueryResult) LatencyStats {
	if len(results) == 0 {
		return LatencyStats{}
	}

	latencies := make([]float64, 0, len(results))
	var total time.Duration
	success := 0
	failed := 0

	for _, r := range results {
		latencies = append(latencies, r.Latency.Seconds()*1000)
		total += r.Latency
		if r.Success {
			success++
		} else {
			failed++
		}
	}

	sort.Float64s(latencies)

	n := float64(len(latencies))
	mean := total.Seconds() * 1000 / n

	var stdDev float64
	for _, l := range latencies {
		stdDev += (l - mean) * (l - mean)
	}
	stdDev = math.Sqrt(stdDev / n)

	stats := LatencyStats{
		Min:     time.Duration(latencies[0]) * time.Millisecond,
		Max:     time.Duration(latencies[len(latencies)-1]) * time.Millisecond,
		Mean:    time.Duration(mean) * time.Millisecond,
		Median:  time.Duration(latencies[int(n*0.5)]) * time.Millisecond,
		P95:     time.Duration(latencies[int(n*0.95)]) * time.Millisecond,
		P99:     time.Duration(latencies[int(n*0.99)]) * time.Millisecond,
		StdDev:  time.Duration(stdDev) * time.Millisecond,
		Total:   total,
		Count:   len(results),
		Success: success,
		Failed:  failed,
	}

	if total > 0 {
		stats.TPS = float64(success) / total.Seconds()
	}

	return stats
}
