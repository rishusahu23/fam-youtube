CREATE INDEX IF NOT EXISTS
    records_created_at_idx ON records (created_at);
CREATE INDEX IF NOT EXISTS
    records_updated_at_idx ON records (updated_at);
CREATE INDEX IF NOT EXISTS
    records_published_at_idx ON records (published_at);