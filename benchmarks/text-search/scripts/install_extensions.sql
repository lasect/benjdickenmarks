-- Install PostgreSQL extensions for benchmarking

-- pg_trgm (should be available in contrib)
CREATE EXTENSION IF NOT EXISTS pg_trgm;

-- Note: pgroonga requires separate installation
-- It's not available in standard PostgreSQL and needs to be installed via:
-- https://pgroonga.github.io/install/

-- Note: pg_textsearch requires PostgreSQL 17+ and separate installation
-- https://github.com/timescale/pg_textsearch

-- Note: ParadeDB's pg_search is available in the paradedb container
