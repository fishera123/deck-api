package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"runtime/debug"
)

func httpOut(payload interface{}, status int, w http.ResponseWriter) {
	obj, _ := json.MarshalIndent(payload, "", " ")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(obj)
}

func (app *Application) httpOutOk(payload interface{}, w http.ResponseWriter) {
	httpOut(payload, http.StatusOK, w)
}

func (app *Application) handleServerError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.ErrorLog.Output(2, trace)
	httpOut(map[string]string{
		"error": http.StatusText(http.StatusInternalServerError),
	}, http.StatusInternalServerError, w)
}

func (app *Application) handleClientError(w http.ResponseWriter, status int, err error) {
	httpOut(map[string]string{
		"error": err.Error(),
	}, status, w)
}

func (app *Application) handleNotFoundError(w http.ResponseWriter) {
	app.handleClientError(w, http.StatusNotFound, errors.New(http.StatusText(http.StatusNotFound)))
}
