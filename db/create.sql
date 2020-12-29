CREATE TABLE IF NOT EXISTS snippets (
  id serial PRIMARY KEY,
  title varchar(100) NOT NULL,
  content text NOT NULL,
  created timestamptz NOT NULL DEFAULT NOW(),
  expires timestamptz NOT NULL DEFAULT NOW() + INTERVAL '1' year
);

CREATE INDEX IF NOT EXISTS idx_snippets_created ON snippets(created);

GRANT ALL ON snippets TO snippetbox;

GRANT ALL ON snippets_id_seq TO snippetbox;

CREATE TABLE IF NOT EXISTS users (
  id serial PRIMARY KEY,
  name varchar(255) NOT NULL,
  email varchar(255) NOT NULL,
  hashed_password char(60) NOT NULL,
  created timestamptz NOT NULL DEFAULT NOW(),
  active boolean NOT NULL DEFAULT TRUE
);

ALTER TABLE
  users
ADD
  CONSTRAINT users_uc_email UNIQUE(email);

GRANT ALL ON users TO snippetbox;

GRANT ALL ON users_id_seq TO snippetbox;
