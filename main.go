package main

/**
  * main.go:
  * - database connection setup
  * - router and middleware setup
  * - server setup
  * - Index handler, Not found handler, Header middleware
  */

import (
    "github.com/julienschmidt/httprouter"
    "github.com/urfave/negroni"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "net/http"
    "log"
    "time"
    "github.com/moritzschramm/todo-demo-go/controllers"
)


/**
  * returns html index file
  */
func Index(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
    http.ServeFile(res, req, "frontend/index.html")
}

/**
  * headers which will be sent on every response
  */
func HeaderMiddleware(res http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
  
    res.Header().Set("x-content-type-options", "nosniff")
    res.Header().Set("x-frame-options", "SAMEORIGIN")
    res.Header().Set("x-xss-protection", "1; mode=block")
    next(res, req)
}

func NotFoundHandler (res http.ResponseWriter, req *http.Request) {
    http.ServeFile(res, req, "frontend/404.html")
}

/**
  * open database connection
  * setup routes and logging middleware
  * start server (listen on localhost:8000)
  */
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

    // setup router
    router := httprouter.New()
    router.NotFound = http.HandlerFunc(NotFoundHandler)

    // static files
    router.GET("/", Index)
    router.ServeFiles("/dist/*filepath", http.Dir("frontend/dist"))

    // api
    todo := &controllers.Todo{DB: db}
    router.POST("/todos",               todo.ShowAll)
    router.POST("/todo",                todo.Create)
    router.POST("/edit/todo/:id",       todo.Edit)
    router.POST("/delete/todo/:id",     todo.Delete)

    n := negroni.New()
    n.Use(negroni.NewLogger())
    n.Use(negroni.NewRecovery())
    n.Use(negroni.HandlerFunc(HeaderMiddleware))
    n.UseHandler(router)

    log.Println("Starting server on http://localhost:8000")
    log.Fatal(http.ListenAndServe(":8000", n))
}