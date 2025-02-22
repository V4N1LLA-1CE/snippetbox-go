-- Create a `snippets` table
CREATE TABLE snippets (
  id SERIAL PRIMARY KEY,
  title VARCHAR(100) NOT NULL,
  content TEXT NOT NULL,
  created TIMESTAMP NOT NULL,
  expires TIMESTAMP NOT NULL
);

-- Add an index on the created column
CREATE INDEX idx_snippets_created ON snippets(created);

-- Add dummy data
INSERT INTO snippets (title, content, created, expires) VALUES (
  'An old silent pond',
  E'An old silent pond...\nA frog jumps into the pond,\nsplash! Silence again.\n\n- Matsui Bashō',
  CURRENT_TIMESTAMP,
  CURRENT_TIMESTAMP + INTERVAL '365 days'
);

INSERT INTO snippets (title, content, created, expires) VALUES (
  'Over the wintry forest',
  E'Over the wintry\nforest, winds howl in range\nwith no leaves to blow.\n\n- Natsume Soseki',
  CURRENT_TIMESTAMP,
  CURRENT_TIMESTAMP + INTERVAL '365 days'
);

INSERT INTO snippets (title, content, created, expires) VALUES (
  'First autumn morning',
  E'First autumn morning\nthe mirror I stare into\nshows my father''s face.\n\n- Murakami Kijo',
  CURRENT_TIMESTAMP,
  CURRENT_TIMESTAMP + INTERVAL '7 days'
);
