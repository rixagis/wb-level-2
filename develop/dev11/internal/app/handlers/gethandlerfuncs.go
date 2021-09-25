package handlers

import (
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/rixagis/wb-level-2/develop/dev11/internal/app/models"
)

// parseDate получает дату из параметров в url и возвращает ее перевод в Time и ошибку
func parseDate(u *url.URL) (time.Time, error) {
	dateParam := u.Query().Get("date")
	return time.Parse(models.TimeFormat, dateParam)
}

// eventsForDayHandler обрабатывает запросы на получение событий в указанный день
func (h *Handler) eventsForDayHandler(w http.ResponseWriter, r *http.Request) {
	date, err := parseDate(r.URL)
	log.Println(date)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(MakeJSONErrorResponse(MsgWrongDateFormat))
		return
	}

	events := h.storage.ReadEventsForDay(date)
	w.WriteHeader(http.StatusOK)
	w.Write(MakeJSONResultResponse(events))
}

// eventsForWeekHandler обрабатывает запросы на получение событий в указанную неделю
func (h *Handler) eventsForWeekHandler(w http.ResponseWriter, r *http.Request) {
	date, err := parseDate(r.URL)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(MakeJSONErrorResponse(MsgWrongDateFormat))
		return
	}

	events := h.storage.ReadEventsForWeek(date)
	w.WriteHeader(http.StatusOK)
	w.Write(MakeJSONResultResponse(events))
}

// eventsForMonthHandler обрабатывает запросы на получение событий в указанный месяц
func (h *Handler) eventsForMonthHandler(w http.ResponseWriter, r *http.Request) {
	date, err := parseDate(r.URL)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(MakeJSONErrorResponse(MsgWrongDateFormat))
		return
	}

	events := h.storage.ReadEventsForMonth(date)
	w.WriteHeader(http.StatusOK)
	w.Write(MakeJSONResultResponse(events))
}
