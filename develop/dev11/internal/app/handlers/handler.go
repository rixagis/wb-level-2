package handlers

import (
	"log"
	"net/http"
	"time"

	"github.com/rixagis/wb-level-2/develop/dev11/internal/app/storage"
)

// Handler - структура, обеспечивающая обработку http запросов
type Handler struct {
	storage *storage.EventStorage
}

// NewHandler - конструктор Handler
func NewHandler(storage *storage.EventStorage) *Handler {
	return &Handler{storage: storage}
}

// InitRoutes настраивает все роуты и возвращает настроенный мутекс
func (h *Handler) InitRoutes() *http.ServeMux {
	mux := http.ServeMux{}

	mux.HandleFunc("/create_event", LogRequest(WithMethod(http.HandlerFunc(h.createEventHandler), http.MethodPost)))
	mux.HandleFunc("/update_event", LogRequest(WithMethod(http.HandlerFunc(h.updateEventHandler), http.MethodPost)))
	mux.HandleFunc("/delete_event", LogRequest(WithMethod(http.HandlerFunc(h.deleteEventHandler), http.MethodPost)))
	mux.HandleFunc("/events_for_day", LogRequest(WithMethod(http.HandlerFunc(h.eventsForDayHandler), http.MethodGet)))
	mux.HandleFunc("/events_for_week", LogRequest(WithMethod(http.HandlerFunc(h.eventsForWeekHandler), http.MethodGet)))
	mux.HandleFunc("/events_for_month", LogRequest(WithMethod(http.HandlerFunc(h.eventsForMonthHandler), http.MethodGet)))

	return &mux
}

var (
	MsgWrongDateFormat = "wrong date format"
	MsgMethodNotAllowed = "method not allowed"
	MsgWrongInputData = "wrong input data"
	MsgCreateSuccess = "successfully created event"
	MsgUpdateSuccess = "successfully updated event"
	MsgDeleteSuccess = "successfully deleted event"
)

// WithMethod перехватывает запрос и возвращает клиенту ошибку, если запрос имеет неверный метод
func WithMethod(next http.Handler, method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == method {
			next.ServeHTTP(w, r)
		} else {
			w.WriteHeader(500)
			w.Write(MakeJSONErrorResponse(MsgMethodNotAllowed))
		}
	}
}

// LogRequest логгирует каждый полученный запрос
func LogRequest(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		next.ServeHTTP(w, r)
		passed := time.Since(startTime)
		log.Printf("Request %s, time elapsed: %s\n", r.RequestURI, passed)
	}
}