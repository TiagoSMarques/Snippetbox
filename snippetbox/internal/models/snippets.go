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
	//the SQL statement we want to exeute. split over 2 lines thats why ``
	stmt := `INSERT INTO snippets (title, content, created, expires) 
			 VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`

	//use de Exec method on the embeded connection pool
	result, err := m.DB.Exec(stmt, tile, content, expires)
	if err != nil {

		return 0, nil
	}

	//get the id of the newly inserted record to the snippets table
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

// this will return a specific snippet based on its id
func (m *SnippetModel) Get(id int) (*Snippet, error) {
	return nil, nil
}

// this will return the 10 most recently created snippets
func (m *SnippetModel) Lastest() ([]*Snippet, error) {
	return nil, nil
}
