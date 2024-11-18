package models

import (
	"database/sql"
	"time"
)

// snippet type for holding data of individual snippet
type Snippet struct {
	Id      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

type SnippetModel struct {
	DB *sql.DB
}

// insert snippet into db
func (m *SnippetModel) Insert(title, content string, expires int) (int, error) {
	// prepare statement
	stmt := `INSERT INTO snippets (title, content, created, expires)
  VALUES ($1, $2, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP + INTERVAL '1 day' * $3)
  RETURNING id`

	var id int64
	if err := m.DB.QueryRow(stmt, title, content, expires).Scan(&id); err != nil {
		return 0, err
	}

	// assert int type to convert int64 to int
	return int(id), nil
}

// return specific snippet based on id
func (m *SnippetModel) Get(id int) (Snippet, error) {
	return Snippet{}, nil
}

// return the 10 most recently created snippets
func (m *SnippetModel) Latest() ([]Snippet, error) {
	return nil, nil
}
