CREATE TABLE IF NOT EXISTS snippets (
  id serial PRIMARY KEY,
  title varchar(100) NOT NULL,
  content text NOT NULL,
  created timestamptz NOT NULL DEFAULT NOW(),
  expires timestamptz NOT NULL DEFAULT NOW() + INTERVAL '1' year
);

CREATE INDEX IF NOT EXISTS idx_snippets_created ON snippets(created);
