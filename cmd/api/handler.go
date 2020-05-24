package api

import (
	"net/http"
)

func (app *Application) routerNotFound(w http.ResponseWriter, r *http.Request) {
	app.handleNotFoundError(w)
}

func (app *Application) checkHealth(w http.ResponseWriter, req *http.Request) {
	app.httpOutOk(map[string]string{
		"status": "ok",
	}, w)
}
