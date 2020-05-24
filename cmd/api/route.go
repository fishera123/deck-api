package api

import (
	"github.com/gorilla/mux"
)

func (app *Application) Routes(router *mux.Router) {
	apiRouter := router.PathPrefix("/v1").Subrouter()

	apiRouter.HandleFunc("/health", app.checkHealth).Methods("GET")

	apiRouter.PathPrefix("/").HandlerFunc(app.routerNotFound)
}
