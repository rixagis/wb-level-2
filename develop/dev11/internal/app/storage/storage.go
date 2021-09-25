package storage

import (
	"errors"
	"log"
	"time"

	"github.com/rixagis/wb-level-2/develop/dev11/internal/app/models"
)


var ErrEventNotFound = errors.New("event not found")

// EventStorage представляет абстракцию над хранилищем данных (событий)
type EventStorage struct {
	nextID int					// следующий id для выдачи
	events map[int]models.Event
}

// NewEventStorage - конструктор EventStorage
func NewEventStorage() *EventStorage {
	s := EventStorage{}
	s.nextID = 0
	s.events = make(map[int]models.Event)
	return &s
}

// getNextID возвращает следующий по порядку id
func (s *EventStorage) getNextID() int {
	s.nextID++
	return s.nextID
}

// CreateEvent записывает event в хранилище, возвращает копию записанного event и ошибку
func (s *EventStorage) CreateEvent(event models.Event) (models.Event, error) {
	eventID := s.getNextID()
	event.EventID = eventID
	s.events[eventID] = event
	log.Println(event)
	return event, nil
}

// UpdateEvent обновляет event с указанным eventID данными из newEvent, возвращает копию записанного event и ошибку
func (s *EventStorage) UpdateEvent(eventID int, newEvent models.Event) (models.Event, error) {
	event, ok := s.events[eventID]
	if !ok {
		return models.Event{}, ErrEventNotFound
	}
	event.EventID = eventID
	s.events[eventID] = event
	return event, nil
}

// DeleteEvent удаляет event с заданным eventID из хранилища, возвращает ошибку
func (s *EventStorage) DeleteEvent(eventID int) error {
	_, ok := s.events[eventID]
	if !ok {
		return ErrEventNotFound
	}

	delete(s.events, eventID)
	return nil
}

// ReadEventsForDay возвращает слайс событий, назначенных на тот же день, что и date
func (s *EventStorage) ReadEventsForDay(date time.Time) []models.Event {
	result := []models.Event{}
	year, month, day := date.Date()
	for _, event := range s.events {
		y, m, d := event.Date.Date()
		if year == y && month == m && day == d {
			result = append(result, event)
		}
	}
	return result
}

// ReadEventsForWeek возвращает слайс событий, назначенных на ту же неделю, что и date
func (s *EventStorage) ReadEventsForWeek(date time.Time) []models.Event {
	result := []models.Event{}
	year, week := date.ISOWeek()
	for _, event := range s.events {
		y, w := event.Date.ISOWeek()
		if year == y && week == w {
			result = append(result, event)
		}
	}
	return result
}

// ReadEventsForMonth возвращает слайс событий, назначенных на тот же месяц, что и date
func (s *EventStorage) ReadEventsForMonth(date time.Time) []models.Event {
	result := []models.Event{}
	year, month, _ := date.Date()
	for _, event := range s.events {
		y, m, _ := event.Date.Date()
		if year == y && month == m {
			result = append(result, event)
		}
	}
	return result
}