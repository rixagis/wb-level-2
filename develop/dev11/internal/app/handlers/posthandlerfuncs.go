package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/rixagis/wb-level-2/develop/dev11/internal/app/models"
)

// decodeEvent получает объект event из тела запроса
func decodeEvent(r *http.Request) (models.Event, error) {
	var event models.Event
	err := json.NewDecoder(r.Body).Decode(&event)
	log.Println(err)
	return event, err
}

// createEventHandler обрабатывает запрос на создание события
func (h *Handler) createEventHandler(w http.ResponseWriter, r *http.Request) {
	event, err := decodeEvent(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(MakeJSONErrorResponse(MsgWrongInputData))
		return
	}

	_, err = h.storage.CreateEvent(event)
	if err != nil {
		w.WriteHeader(503)
		w.Write(MakeJSONErrorResponse(err.Error()))
		return
	}
	
	w.WriteHeader(http.StatusOK)
	w.Write(MakeJSONPostSuccessResultResponse(MsgCreateSuccess))
}

// updateEventHandler обрабатыает запрос на обновление события
func (h *Handler) updateEventHandler(w http.ResponseWriter, r *http.Request) {
	event, err := decodeEvent(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(MakeJSONErrorResponse(MsgWrongInputData))
		return
	}

	_, err = h.storage.UpdateEvent(event.EventID, event)
	if err != nil {
		w.WriteHeader(503)
		w.Write(MakeJSONErrorResponse(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(MakeJSONPostSuccessResultResponse(MsgUpdateSuccess))
}

// deleteEventHandler обрабатывает запрос на удаление события
func (h *Handler) deleteEventHandler(w http.ResponseWriter, r *http.Request) {
	event, err := decodeEvent(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(MakeJSONErrorResponse(MsgWrongInputData))
		return
	}

	err = h.storage.DeleteEvent(event.EventID)
	if err != nil {
		w.WriteHeader(503)
		w.Write(MakeJSONErrorResponse(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(MakeJSONPostSuccessResultResponse(MsgDeleteSuccess))
}