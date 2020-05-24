package main

import (
	"deck-api/cmd/api"
	"deck-api/pkg/database"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime|log.Lmicroseconds)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Llongfile)

	// todo: inject database instance in models
	db, err := database.ConnectDb()
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()

	app := &api.Application{
		ErrorLog:  errorLog,
		InfoLog:   infoLog,
		DeckModel: &database.DeckModel{DB: db},
	}

	router := mux.NewRouter()

	router.Use(app.LogRequest)
	router.Use(app.RecoverPanic)
	app.Routes(router)

	infoLog.Printf("Starting server on %s", ":8080")
	err = http.ListenAndServe(":8080", router) // todo: read from env variable
	errorLog.Fatal(err)
}
