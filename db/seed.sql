TRUNCATE TABLE snippets;

INSERT INTO
  snippets(title, content)
VALUES
  (
    'An old silent pond',
    E'An old silent pond...\nA frog jumps into the pond,\nsplash! Silence again.\n\n– Matsuo Bashō'
  ),
  (
    'Over the wintry forest',
    E'Over the wintry\nforest, winds howl in rage\nwith no leaves to blow.\n\n– Natsume Soseki'
  );

INSERT INTO
  snippets(title, content, expires)
VALUES
  (
    'First autumn morning',
    E'First autumn morning\nthe mirror I stare into\nshows my father''s face.\n\n– Murakami Kijo',
    NOW() + INTERVAL '7' DAY
  )
