create table if not exists records (
    id text not null,
    title text not null,
    description text not null,
    metadata jsonb,
    published_at TIMESTAMPTZ,
    created_at      		TIMESTAMP WITH TIME ZONE DEFAULT now(),
    updated_at      		TIMESTAMP WITH TIME ZONE DEFAULT now(),
    PRIMARY KEY 			(id)
);

CREATE INDEX records_title_idx ON records (title);
CREATE INDEX records_description_idx ON records (description);
