package postgresql

import (
	"database/sql"
	"snippetbox/pkg/models"
)

// SnippetModel wraps a sql.DB connection pool.
type SnippetModel struct {
	DB *sql.DB
}

// Insert a new snippet into the database.
func (m *SnippetModel) Insert(title, content, expires string) (int, error) {
	stmt := `INSERT INTO snippets (title, content, expires)
	VALUES ($1, $2, (NOW() + '1 day'::INTERVAL * $3)::timestamptz)
	RETURNING id`

	// VALUES ($1, $2, NOW() + $3::interval)

	id := 0
	err := m.DB.
		QueryRow(stmt, title, content, expires).
		Scan(&id)

	return id, err
}

// Get return a specific snippet.
func (m *SnippetModel) Get(id int) (int, error) {
	return 0, nil
}

// Latest returns the most 10 recently created snippets.
func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	return nil, nil
}
