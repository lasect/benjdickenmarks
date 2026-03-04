# Text Search Benchmarks

Comparing PostgreSQL text search tools and external search engines.

## Quick Start

```bash
cd benchmarks/text-search
make up
make install-extensions
```

## Tools Benchmarked

| # | Tool | Type | Port |
|---|------|------|------|
| 1 | tsvector | Native PostgreSQL | 5432 |
| 2 | pg_trgm | PostgreSQL Extension | 5432 |
| 3 | pgroonga | PostgreSQL Extension | 5432 |
| 4 | ParadeDB (pg_search) | PostgreSQL Extension | 5433 |
| 5 | pg_textsearch | PostgreSQL Extension | 5432 |
| 6 | Elasticsearch | External | 9200 |
| 7 | Meilisearch | External | 7700 |
