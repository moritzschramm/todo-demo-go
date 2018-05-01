package controllers

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

func setHeaders(res http.ResponseWriter) {

    res.Header().Set("Content-Type", "application/json; charset=utf-8")
}

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