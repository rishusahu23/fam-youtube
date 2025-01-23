CREATE INDEX IF NOT EXISTS
    records_search_idx ON records USING GIN (to_tsvector('english', title || ' ' || description));