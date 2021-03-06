package api

import (
	"github.com/gorilla/mux"
)

func (app *Application) Routes(router *mux.Router) {
	apiRouter := router.PathPrefix("/v1").Subrouter()

	apiRouter.HandleFunc("/health", app.checkHealth).Methods("GET")
	apiRouter.HandleFunc("/decks", app.createDeck).Methods("POST")
	apiRouter.HandleFunc("/decks/{id}", app.showDeck).Methods("GET")
	apiRouter.HandleFunc("/decks/{id}/draw", app.drawDeck).Methods("POST")

	apiRouter.PathPrefix("/").HandlerFunc(app.routerNotFound)
}
