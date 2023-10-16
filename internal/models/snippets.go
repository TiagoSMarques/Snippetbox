package models

import (
	"database/sql"
	"errors"
	"time"
)

// Define a snippet struct to hold the data for the individual snippet. The fields of the struct correspond to the fields of the MySQL snippets table
type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
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
	stmt := `SELECT id, title, content, created, expires FROM snippets
			 WHERE expires > UTC_TIMESTAMP AND id = ?`

	//use queryrow on the conn pool to execute the sql statement - this returns a pointer to a sql.row object which holds the result from de db
	row := m.DB.QueryRow(stmt, id)

	//initialize a pointer to a new zeroed snippet struct.
	s := &Snippet{}

	//use row.scan to copy values from each field in sql.row to the corrensponding field in the cstruct
	//the args to row.scan are pointers to where we want to copy the data into, and the number of args must be the same as the number of cols returned by our statment
	err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
	// If the query returns no rows, then row.Scan() will return a sql.ErrNoRows error
	//We use the errors.Is() function check for that error specifically, and return our own ErrNoRecord error instead
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		} else {
			return nil, err
		}
	}
	return s, nil
}

// this will return the 10 most recently created snippets
func (m *SnippetModel) Lastest() ([]*Snippet, error) {
	stmt := `SELECT id, title, content, created, expires FROM snippets
			WHERE expires > UTC_TIMESTAMP() ORDER BY id DESC LIMIT 10`

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	// We need to ensure that the sql.Rows resultset id always properly closed before the method lastest returns. This should always be called after checking for errors
	//If we dont defer the connection to the db will remain opened
	defer rows.Close()

	snipets := []*Snippet{}

	// rows.Next iterates over the resulset
	for rows.Next() {

		s := &Snippet{}

		err = rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)

		if err != nil {
			return nil, err
		}
		snipets = append(snipets, s)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return snipets, nil
}
