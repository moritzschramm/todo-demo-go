package main

import (
    "github.com/julienschmidt/httprouter"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "net/http"
    "log"
    "time"
    "github.com/moritzschramm/todo-demo-go/controllers"
)

func Index(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {

    res.Header().Set("x-content-type-options", "nosniff")
    res.Header().Set("x-frame-options", "SAMEORIGIN")
    res.Header().Set("x-xss-protection", "1; mode=block")
    http.ServeFile(res, req, "public/index.html")
}

func main() {

    // database connection setup
    db, err := sql.Open("mysql", "homestead:secret@(mysql)/homestead?parseTime=true")
    if err != nil {
        log.Fatal("Error connecting to DB: ", err.Error())
        return
    }
    // check database connection
    for i := 0; i < 5; i++ {                                        // try connecting to database 5 times. 
        err = db.Ping()                                             // this is needed when the mysql docker 
        if err != nil {                                             // container first starts and is creating
            log.Println("Error connecting to DB: ", err.Error())    // the database. 
            log.Println("Trying to connect again in 5 seconds")
        } else {
            break
        }
        time.Sleep(5 * time.Second)                                 // wait for 5 seconds after each try
    }
    if err != nil {
        log.Fatal("Could not connect to DB: ", err.Error())
        return
    }

    defer db.Close()

    // setup routes
    router := httprouter.New()

    // static content
    router.GET("/", Index)
    router.ServeFiles("/assets/*filepath", http.Dir("public"))

    // api
    todo := &controllers.Todo{DB: db}
    router.POST("/todos",               todo.ShowAll)
    router.POST("/todo",                todo.Create)
    router.POST("/edit/todo/:id",       todo.Edit)
    router.POST("/delete/todo/:id",     todo.Delete)

    log.Println("Starting server on http://localhost:8000")
    log.Fatal(http.ListenAndServe(":8000", router))
}