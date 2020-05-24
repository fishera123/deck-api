package api

import (
	dto "deck-api/pkg/dtos"
	"encoding/json"
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

func (app *Application) createDeck(w http.ResponseWriter, req *http.Request) {
	var bodyJson struct {
		Shuffled bool   `json:"shuffled"`
		Cards    string `json:"cards"`
	}
	err := json.NewDecoder(req.Body).Decode(&bodyJson)
	if err != nil {
		app.handleClientError(w, http.StatusBadRequest, err)
		return
	}

	deck, err := app.CreateDeck(bodyJson.Cards, bodyJson.Shuffled)

	// todo: CreateDeck service function may fail database operation, so just returning 400 is wrong, should be checked what kind
	// error it is and based on that should set the status code. 500 or 400
	if err != nil {
		app.handleClientError(w, 400, err)
		return
	}

	app.httpOutOk(dto.ToDeckDto(deck), w)
}
