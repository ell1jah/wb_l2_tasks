package model

import "time"

type Event struct {
	Id      int       `json:"user_id"`
	EventId int       `json:"event_id"`
	Title   string    `json:"title"`
	Date    time.Time `json:"date"`
}
