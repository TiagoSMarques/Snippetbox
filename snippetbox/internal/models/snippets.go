package models

import (
	"database/sql"
	"time"
)

// Define a snippet struct to hold the data for the individual snippet. The fields of the struct correspond to the fields of the MySQL snippets table
type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expire  time.Time
}

// Define a SnippetModel type which wraps a sql.DB connection pool.
type SnippetModel struct {
	DB *sql.DB
}

// this will insert a new snippet in the database
func (m *SnippetModel) Insert(tile string, content string, expires int) (int, error) {
	return 0, nil
}

// this will return a specific snippet based on its id
func (m *SnippetModel) Get(id int) (*Snippet, error) {
	return nil, nil
}

// this will return the 10 most recently created snippets
func (m *SnippetModel) Lastest() ([]*Snippet, error) {
	return nil, nil
}
