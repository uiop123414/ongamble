CREATE TABLE IF NOT EXISTS articles (
    id bigserial PRIMARY KEY,
    name text NOT NULL,
    publish boolean NOT NULL,
    reading_time text NOT NULL,
    username text NOT NULL,
    html_list text NOT NULL,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    updated_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    version integer NOT NULL DEFAULT 1
);