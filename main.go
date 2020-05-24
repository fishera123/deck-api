package main

import (
	"deck-api/cmd/api"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime|log.Lmicroseconds)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Llongfile)

	app := &api.Application{
		ErrorLog: errorLog,
		InfoLog:  infoLog,
	}

	router := mux.NewRouter()

	router.Use(app.LogRequest)
	router.Use(app.RecoverPanic)

	infoLog.Printf("Starting server on %s", ":8080")
	err := http.ListenAndServe(":8080", router) // todo: read from env variable
	errorLog.Fatal(err)
}
