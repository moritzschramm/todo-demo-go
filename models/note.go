package models

import (
	"database/sql"
	"time"
)

const (
	TABLE = "notes"
)

type Note struct {
	DB *sql.DB 				`json:"-"`
	ID int 					`json:"id"`
	Note string 			`json:"note"`
	Created_at time.Time 	`json:"created_at,omitempty"`
	Updated_at time.Time 	`json:"updated_at,omitempty"`
}

func Notes(db *sql.DB) ([]*Note, error) {

	rows, err := db.Query("SELECT id, note, created_at, updated_at FROM "+TABLE)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	notes := make([]*Note, 100)

	i := 0
	for rows.Next() {
		var id int
        var note string
        var created_at, updated_at time.Time
        if err := rows.Scan(&id, &note, &created_at, &updated_at); err != nil {
        	return nil, err
        }

        notes[i] = &Note{DB: db, ID: id, Note: note, Created_at: created_at, Updated_at: updated_at}
        i++
	}

	return notes[:i], nil
}

func FindNoteById(db *sql.DB, id int) (*Note, error) {

	var note string
	var created_at, updated_at time.Time
	err := db.QueryRow("SELECT note, created_at, updated_at FROM "+TABLE+" WHERE id = ?", id).Scan(&note, &created_at, &updated_at)
	if err != nil {
		return nil, err
	}

	return &Note{DB: db, ID: id, Note: note, Created_at: created_at, Updated_at: updated_at}, nil
}

func CreateNote(db *sql.DB, note string) (*Note, error) {

	created_at := time.Now();
	_, err := db.Query("INSERT INTO "+TABLE+" (note, created_at, updated_at) VALUES (?, ?, ?)", note, created_at, created_at)
	if err != nil {
		return nil, err
	}

	var id int
	err = db.QueryRow("SELECT id FROM "+TABLE+" ORDER BY created_at DESC LIMIT 1").Scan(&id)
	if err != nil {
		return nil, err
	}

	return &Note{DB: db, ID: id, Note: note, Created_at: created_at, Updated_at: created_at}, nil
}

func (n *Note) Edit(note string) error {

	n.Note = note
	n.Updated_at = time.Now()
	_, err := n.DB.Query("UPDATE "+TABLE+" SET note = ?, updated_at = ? WHERE id = ?", n.Note, n.Updated_at, n.ID)

	return err
}

func (n *Note) Delete() error {

	_, err := n.DB.Query("DELETE FROM "+TABLE+" WHERE id = ?", n.ID)
	n = nil

	return err
}