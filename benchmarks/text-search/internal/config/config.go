package config

import (
	"os"
	"strconv"
)

type Config struct {
	Postgres      PostgresConfig
	Paradedb      PostgresConfig
	Elasticsearch ElasticsearchConfig
	Meilisearch   MeilisearchConfig
	Benchmark     BenchmarkConfig
}

type PostgresConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
}

type ElasticsearchConfig struct {
	Addresses []string
	Index     string
}

type MeilisearchConfig struct {
	Host   string
	APIKey string
	Index  string
}

type BenchmarkConfig struct {
	Dataset          string
	DatasetSize      int
	WarmupQueries    int
	BenchmarkQueries int
	Concurrency      []int
}

func Load() *Config {
	return &Config{
		Postgres: PostgresConfig{
			Host:     getEnv("POSTGRES_HOST", "localhost"),
			Port:     getEnvInt("POSTGRES_PORT", 5432),
			User:     getEnv("POSTGRES_USER", "benchmark"),
			Password: getEnv("POSTGRES_PASSWORD", "benchmark"),
			Database: getEnv("POSTGRES_DB", "benchmark"),
		},
		Paradedb: PostgresConfig{
			Host:     getEnv("PARADEDB_HOST", "localhost"),
			Port:     getEnvInt("PARADEDB_PORT", 5433),
			User:     getEnv("PARADEDB_USER", "benchmark"),
			Password: getEnv("PARADEDB_PASSWORD", "benchmark"),
			Database: getEnv("PARADEDB_DB", "benchmark"),
		},
		Elasticsearch: ElasticsearchConfig{
			Addresses: []string{"http://localhost:9200"},
			Index:     "benchmark",
		},
		Meilisearch: MeilisearchConfig{
			Host:   getEnv("MEILISEARCH_HOST", "http://localhost:7700"),
			APIKey: getEnv("MEILISEARCH_KEY", "benchmark"),
			Index:  "benchmark",
		},
		Benchmark: BenchmarkConfig{
			Dataset:          getEnv("DATASET", "ecommerce"),
			DatasetSize:      getEnvInt("DATASET_SIZE", 100000),
			WarmupQueries:    getEnvInt("WARMUP_QUERIES", 50),
			BenchmarkQueries: getEnvInt("BENCHMARK_QUERIES", 1000),
			Concurrency:      []int{1, 10, 50},
		},
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intVal, err := strconv.Atoi(value); err == nil {
			return intVal
		}
	}
	return defaultValue
}
