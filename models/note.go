package models

/**
  * note.go:
  * - model representation of a note (stores text, done flag, creation and update time as well as unique id)
  * - note can be easily formatted to JSON via json.Marshal(note)
  * - requires database connection
  * - methods to fetch all, find one, create, edit and delete note(s)
  */

import (
    "database/sql"
    "time"
)

const (
    TABLE = "notes"
    SQL_FALSE = "\x00"  // how the mysql implementation represents "false"...
)

type Note struct {
    DB *sql.DB              `json:"-"`      // ignore in json representation
    ID int                  `json:"id"`
    Note string             `json:"note"`
    Done bool               `json:"done"`
    Created_at time.Time    `json:"created_at,omitempty"`
    Updated_at time.Time    `json:"updated_at,omitempty"`
}

/**
  * select all notes from database and return slice of notes
  */
func Notes(db *sql.DB) ([]*Note, error) {

    rows, err := db.Query("SELECT id, note, done, created_at, updated_at FROM "+TABLE)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    notes := make([]*Note, 100)

    i := 0
    for rows.Next() {
        var id int
        var note string
        var done string
        var created_at, updated_at time.Time
        if err := rows.Scan(&id, &note, &done, &created_at, &updated_at); err != nil {
            return nil, err
        }

        notes[i] = &Note{DB: db, ID: id, Note: note, Done: done != SQL_FALSE, Created_at: created_at, Updated_at: updated_at}
        i++
    }

    return notes[:i], nil
}

/**
  * query database to find note by id and return note object
  */
func FindNoteById(db *sql.DB, id int) (*Note, error) {

    var note string
    var done string
    var created_at, updated_at time.Time
    err := db.QueryRow("SELECT note, done, created_at, updated_at FROM "+TABLE+" WHERE id = ?", id).Scan(&note, &done, &created_at, &updated_at)
    if err != nil {
        return nil, err
    }

    return &Note{DB: db, ID: id, Note: note, Done: done != SQL_FALSE, Created_at: created_at, Updated_at: updated_at}, nil
}

/**
  * insert note into database and return note object with id
  */
func CreateNote(db *sql.DB, note string) (*Note, error) {

    created_at := time.Now();
    rows, err := db.Query("INSERT INTO "+TABLE+" (note, done, created_at, updated_at) VALUES (?, false, ?, ?)", note, created_at, created_at)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var id int
    err = db.QueryRow("SELECT id FROM "+TABLE+" ORDER BY created_at DESC LIMIT 1").Scan(&id)
    if err != nil {
        return nil, err
    }

    return &Note{DB: db, ID: id, Note: note, Done: false, Created_at: created_at, Updated_at: created_at}, nil
}

/**
  * update note in database
  */
func (n *Note) Edit(note string, done bool) error {

    n.Note = note
    n.Done = done
    n.Updated_at = time.Now()
    rows, err := n.DB.Query("UPDATE "+TABLE+" SET note = ?, done = ?, updated_at = ? WHERE id = ?", n.Note, n.Done, n.Updated_at, n.ID)
    defer rows.Close()

    return err
}

/**
  * delete note from database
  */
func (n *Note) Delete() error {

    rows, err := n.DB.Query("DELETE FROM "+TABLE+" WHERE id = ?", n.ID)
    defer rows.Close()

    n = nil

    return err
}