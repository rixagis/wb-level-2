package models

import (
	"log"
	"time"
)

// Event - структура события
type Event struct {
	EventID     int        `json:"event_id"`
	UserID      int        `json:"user_id"`
	Date        Time       `json:"date"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
}

// Time - кастомный тип времени, нужен для правильного перевода из json в Time
type Time struct {
    time.Time
}

const TimeFormat = "02-01-2006"

// UnmarshalJSON переводит строку байт в структуру Time по заданному формату, вызывается при Decode
func (t *Time) UnmarshalJSON(b []byte) error {
    
	newTime, err := time.Parse("\""+TimeFormat+"\"", string(b))
	log.Printf("%q\n",string(b))

	if err != nil {
		*t = Time{time.Now()}
		return err
	}


	*t = Time{newTime}
    return nil
}
