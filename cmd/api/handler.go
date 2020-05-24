package api

import (
	dto "deck-api/pkg/dtos"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
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

func (app *Application) showDeck(w http.ResponseWriter, req *http.Request) {

	params := mux.Vars(req)
	id := params["id"]
	// todo: move to service, controllers shouldn't be using models directly
	deck, err := app.DeckModel.Get(id)

	if err != nil {
		app.handleClientError(w, http.StatusNotFound, err)
		return
	}

	if deck.IsEmpty() {
		app.handleNotFoundError(w)
		return
	}

	app.httpOutOk(dto.ToOpenDeckDto(deck), w)
}

func (app *Application) drawDeck(w http.ResponseWriter, req *http.Request) {
	var bodyJson struct {
		Count int64 `json:"count"`
	}
	err := json.NewDecoder(req.Body).Decode(&bodyJson)
	if err != nil {
		app.handleClientError(w, http.StatusBadRequest, err)
		return
	}

	if (bodyJson.Count == 0 || bodyJson.Count > 52) || err != nil {
		app.handleClientError(w, http.StatusUnprocessableEntity, errors.New("invalid count query param"))
		return
	}
	params := mux.Vars(req)
	id := params["id"]

	if err != nil {
		app.handleServerError(w, err)
		return
	}

	deck, err := app.DeckModel.Get(id)
	if err != nil {
		app.handleNotFoundError(w)
		return
	}

	drawnCards, err := app.DrawDeck(deck, bodyJson.Count) //todo get count query

	if err != nil {
		app.handleClientError(w, 400, err)
	}

	app.httpOutOk(dto.ToCardsDto(drawnCards), w)
}
