package main

import (
	"github.com/julienschmidt/httprouter"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
    "net/http"
    "log"
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
	db, err := sql.Open("mysql", "homestead:secret@/homestead?parseTime=true")
	if err != nil {
		log.Fatal("Error connecting to DB: ", err.Error())
		return
	}
	// check database connection
	err = db.Ping()
	if err != nil {
		log.Fatal("Error connecting to DB: ", err.Error())
		return
	}
	defer db.Close()

	// setup routes
	router := httprouter.New()

	// static content
    router.GET("/", Index)
    router.ServeFiles("/assets/*filepath", http.Dir("public"))

    // api
    todo := &controllers.Todo{db}
    router.POST("/todos", 				todo.ShowAll)
    router.POST("/todo", 				todo.Create)
    router.POST("/edit/todo/:id", 		todo.Edit)
    router.POST("/delete/todo/:id", 	todo.Delete)

    log.Println("Starting server on http://localhost:8000")
    log.Fatal(http.ListenAndServe(":8000", router))
}