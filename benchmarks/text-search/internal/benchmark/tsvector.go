package benchmark

import (
	"context"
	"fmt"
	"time"

	"github.com/benjdickenmarks/text-search/internal/config"
	"github.com/benjdickenmarks/text-search/internal/loader"
	"github.com/benjdickenmarks/text-search/internal/metrics"
	"github.com/jackc/pgx/v5/pgxpool"
)

type TSVectorBenchmark struct {
	cfg   *config.Config
	pool  *pgxpool.Pool
	table string
}

func NewTSVectorBenchmark(cfg *config.Config) *TSVectorBenchmark {
	return &TSVectorBenchmark{
		cfg:   cfg,
		table: "benchmark_tsvector",
	}
}

func (b *TSVectorBenchmark) Name() string {
	return "tsvector"
}

func (b *TSVectorBenchmark) Setup(ctx context.Context) error {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
		b.cfg.Postgres.User, b.cfg.Postgres.Password,
		b.cfg.Postgres.Host, b.cfg.Postgres.Port, b.cfg.Postgres.Database)

	pool, err := pgxpool.New(ctx, connStr)
	if err != nil {
		return err
	}
	b.pool = pool

	_, err = b.pool.Exec(ctx, fmt.Sprintf(`
		DROP TABLE IF EXISTS %s;
		CREATE TABLE %s (
			id BIGSERIAL PRIMARY KEY,
			content TEXT
		);
	`, b.table, b.table))
	return err
}

func (b *TSVectorBenchmark) LoadData(ctx context.Context, data interface{}) error {
	var products []loader.Product
	switch d := data.(type) {
	case []loader.Product:
		products = d
	default:
		products = loader.GenerateProducts(100000)
	}

	for i := 0; i < len(products); i += 1000 {
		end := i + 1000
		if end > len(products) {
			end = len(products)
		}

		batch := make([]string, end-i)
		for j := i; j < end; j++ {
			batch[j-i] = products[j].Name + " " + products[j].Description
		}

		_, err := b.pool.Exec(ctx, fmt.Sprintf(`
			INSERT INTO %s (content) SELECT * FROM UNNEST($1::text[])
		`, b.table), batch)
		if err != nil {
			return err
		}
	}

	return nil
}

func (b *TSVectorBenchmark) BuildIndex(ctx context.Context) error {
	_, err := b.pool.Exec(ctx, fmt.Sprintf(`
		ALTER TABLE %s ADD COLUMN tsv tsvector;
		UPDATE %s SET tsv = to_tsvector('english', content);
		CREATE INDEX idx_tsvector ON %s USING GIN(tsv);
	`, b.table, b.table, b.table))
	return err
}

func (b *TSVectorBenchmark) Warmup(ctx context.Context, queries []string) error {
	for _, q := range queries {
		_, _ = b.pool.Query(ctx, fmt.Sprintf(`
			SELECT * FROM %s WHERE tsv @@ to_tsquery('english', $1)
			LIMIT 10
		`, b.table), q)
	}
	return nil
}

func (b *TSVectorBenchmark) Benchmark(ctx context.Context, queries []string, concurrency int) (metrics.LatencyStats, error) {
	results := make([]metrics.QueryResult, 0, len(queries)*10)

	for i := 0; i < len(queries)*10; i++ {
		query := queries[i%len(queries)]
		start := time.Now()

		rows, err := b.pool.Query(ctx, fmt.Sprintf(`
			SELECT * FROM %s WHERE tsv @@ to_tsquery('english', $1)
			LIMIT 10
		`, b.table), query)

		latency := time.Since(start)

		if err == nil {
			rows.Close()
			results = append(results, metrics.QueryResult{
				Latency: latency,
				Success: true,
			})
		} else {
			results = append(results, metrics.QueryResult{
				Latency: latency,
				Success: false,
			})
		}
	}

	return metrics.CalculateStats(results), nil
}

func (b *TSVectorBenchmark) WriteTest(ctx context.Context, data interface{}) (float64, error) {
	return 0, nil
}

func (b *TSVectorBenchmark) GetIndexSize() (int64, error) {
	return 0, nil
}

func (b *TSVectorBenchmark) Cleanup(ctx context.Context) error {
	if b.pool != nil {
		b.pool.Close()
	}
	return nil
}
