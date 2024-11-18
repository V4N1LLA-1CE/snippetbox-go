package models

import (
	"database/sql"
	"errors"
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
	stmt := `SELECT id, title, content, created, expires FROM snippets
  WHERE expires > CURRENT_TIMESTAMP AND id = $1`

	// create new snippet struct and write to each field
	var s Snippet
	if err := m.DB.QueryRow(stmt, id).Scan(&s.Id, &s.Title, &s.Content, &s.Created, &s.Expires); err != nil {
		// if no rows found, we just return ErrNoRecord
		if errors.Is(err, sql.ErrNoRows) {
			return Snippet{}, ErrNoRecord
		} else {
			return Snippet{}, err
		}
	}

	// return snippet if everything is ok
	return s, nil
}

// return the 10 most recently created snippets
func (m *SnippetModel) Latest() ([]Snippet, error) {
	return nil, nil
}
