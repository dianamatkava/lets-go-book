package models

import (
	"database/sql"
	"errors"
	"time"
)

type Snippet struct {
	ID int
	Title string
	Content string
	Created time.Time
	Expires time.Time
}

type SnippetModel struct {
	DB *sql.DB
}

func (m *SnippetModel) Insert(title string, content string, expires int) (int, error) {
	query := `
		INSERT INTO snippets (title, content, created, expires)
		VALUES ($1, $2, NOW(), NOW() + ($3 * INTERVAL '1 day')) returning id;
	`
	var id int
	err := m.DB.QueryRow(query, title, content, expires).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (m *SnippetModel) Get(id int) (Snippet, error) {
	query := `
		SELECT id, title, content, created, expires 
		FROM snippets
		WHERE expires > NOW() AND id = $1;
	`
	var s Snippet
	err := m.DB.QueryRow(query, id).Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Snippet{}, ErrNoRecord
		} else {
			return Snippet{}, nil
		}
	}
	
	return s, nil
}

func (m *SnippetModel) Latest() ([]Snippet, error) {
	query := `
		SELECT id, title, content, created, expires FROM snippets
		WHERE expires > NOW() ORDER BY id DESC LIMIT 10
	`
	var snippets []Snippet

	rows, err := m.DB.Query(query)
	if err != nil {
		return []Snippet{}, nil
	}

	defer rows.Close()

	for rows.Next() {
		var s Snippet
		err = rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
		if err != nil {
			return nil, err
		}
		snippets = append(snippets, s)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return snippets, nil
}