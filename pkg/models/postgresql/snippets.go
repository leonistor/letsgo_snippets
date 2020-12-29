package postgresql

import (
	"database/sql"
	"errors"
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

	id := 0
	err := m.DB.
		QueryRow(stmt, title, content, expires).
		Scan(&id)

	return id, err
}

// Get return a specific snippet.
func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	stmt := `select id, title, content, created, expires
	FROM snippets
	WHERE expires > now() and id=$1`

	row := m.DB.QueryRow(stmt, id)
	s := &models.Snippet{}
	err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		}
		return nil, err
	}
	return s, nil
}

// Latest returns the most 10 recently created snippets.
func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	return nil, nil
}
