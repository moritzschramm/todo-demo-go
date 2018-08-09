package controllers

/**
  * todo_controller.go:
  * - JSON API to show all, create, edit and delete notes
  * - requires database connection (see struct)
  * - logs errors via log and returns http error (status code 500)
  */

import (
    "net/http"
    "github.com/julienschmidt/httprouter"
    "database/sql"
    "encoding/json"
    "fmt"
    "log"
    "github.com/moritzschramm/todo-demo-go/models"
    "strconv"
)

type Todo struct {
    DB *sql.DB
}

/**
  * set content type header in response
  */
func setHeaders(res http.ResponseWriter) {

    res.Header().Set("Content-Type", "application/json; charset=utf-8")
}

/**
  * return json array of all saved notes
  */
func (t *Todo) ShowAll(res http.ResponseWriter, req *http.Request, params httprouter.Params) {

    setHeaders(res)

    notes, err := models.Notes(t.DB)
    if err != nil {
        log.Println("Error loading notes: ", err)
        http.Error(res, "Internal Server Error", 500)
        return
    }
    
    notesJson, err := json.Marshal(notes)
    if err != nil {
        log.Println("Error loading notes: ", err)
        http.Error(res, "Internal Server Error", 500)
        return
    }
    fmt.Fprint(res, string(notesJson))
}

/**
  * create note, return json representation of note (with new id)
  */
func (t *Todo) Create(res http.ResponseWriter, req *http.Request, params httprouter.Params) {

    setHeaders(res)

    text := req.FormValue("note")
    note, err := models.CreateNote(t.DB, text)
    if err != nil {
        log.Println("Error creating note: ", err)
        http.Error(res, "Internal Server Error", 500)
        return
    }

    noteJson, err := json.Marshal(note)
    if err != nil {
        log.Println("Error creating note: ", err)
        http.Error(res, "Internal Server Error", 500)
        return
    }

    fmt.Fprint(res, string(noteJson))
}

/**
  * try to update note, reutrn OK if found and updated, not found otherwise
  */
func (t *Todo) Edit(res http.ResponseWriter, req *http.Request, params httprouter.Params) {

    setHeaders(res)

    note := t.getNoteById(res, req, params)
    if note != nil {
        
        text := req.FormValue("note")
        done := false
        if req.FormValue("done") == "true" {
            done = true
        }
        note.Edit(text, done)

        res.WriteHeader(http.StatusOK)
    } else {
        res.WriteHeader(http.StatusNotFound)
    }
}

/**
  * delete note if it exists, return not found otherwise
  */
func (t *Todo) Delete(res http.ResponseWriter, req *http.Request, params httprouter.Params) {

    setHeaders(res)

    note := t.getNoteById(res, req, params)

    if note != nil {
        note.Delete()
        res.WriteHeader(http.StatusOK)
    } else {
        res.WriteHeader(http.StatusNotFound)
    }
}

/**
  * try to find note by id and return it, return not found otherwise
  */
func (t *Todo) getNoteById(res http.ResponseWriter, req *http.Request, params httprouter.Params) *models.Note {

    id, err := strconv.Atoi(params.ByName("id"))
    if err != nil {
        log.Println("Error getting note by id: ", err)
        http.NotFound(res, req)
        return nil
    }
    note, err := models.FindNoteById(t.DB, id)
    if err != nil {
        log.Println("Error getting note by id: ", err)
        http.NotFound(res, req)
        return nil
    }

    return note
}