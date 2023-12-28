package database

import (
	"d11/model"
	"fmt"
	"sync"
	"time"
)

type DataBase struct {
	mx    sync.RWMutex
	event map[int][]model.Event
}

func NewDataBase() *DataBase {
	return &DataBase{event: make(map[int][]model.Event)}
}

func (d *DataBase) CreateEvent(event model.Event) error {
	d.mx.Lock()
	defer d.mx.Unlock()
	if _, ok := d.event[event.Id]; ok {
		if _, ok := d.CheckEvent(event); ok {
			return fmt.Errorf("duplicate event")
		}
	}
	d.event[event.Id] = append(d.event[event.Id], event)
	return nil
}

func (d *DataBase) DeleteEvent(event model.Event) error {
	d.mx.Lock()
	defer d.mx.Unlock()
	if _, ok := d.event[event.Id]; ok {
		if idx, ok := d.CheckEvent(event); ok {
			d.event[event.Id] = append(d.event[event.Id][:idx], d.event[event.Id][idx+1:]...)
			return nil
		}
	}
	return fmt.Errorf("event not found")
}

func (d *DataBase) UpdateEvent(event model.Event) error {
	d.mx.Lock()
	defer d.mx.Unlock()
	events, ok := d.event[event.Id]
	if ok {
		if idx, ok := d.CheckEvent(event); ok {
			events[idx] = event
			return nil
		}
	}
	return fmt.Errorf("event not found")
}

func (d *DataBase) CheckEvent(event model.Event) (int, bool) {
	for i, elem := range d.event[event.Id] {
		if elem.EventId == event.EventId {
			return i, true
		}
	}
	return 0, false
}

func (d *DataBase) Events_for_day(user_id int, date time.Time) ([]model.Event, error) {
	d.mx.RLock()
	defer d.mx.RUnlock()
	var events []model.Event
	if _, ok := d.event[user_id]; !ok {
		return nil, fmt.Errorf("user not found")
	}
	for _, event := range d.event[user_id] {
		if event.Date.Year() == date.Year() && event.Date.Month() == date.Month() && event.Date.Day() == date.Day() {
			events = append(events, event)
		}
	}
	return events, nil
}

func (d *DataBase) Events_for_week(user_id int, date time.Time) ([]model.Event, error) {
	d.mx.RLock()
	defer d.mx.RUnlock()
	var events []model.Event
	if _, ok := d.event[user_id]; !ok {
		return nil, fmt.Errorf("user not found")
	}
	end := date.AddDate(0, 0, 7)
	for _, event := range d.event[user_id] {
		if event.Date.After(date) && event.Date.Before(end) {
			events = append(events, event)
		}
	}
	return events, nil
}

func (d *DataBase) Events_for_month(user_id int, date time.Time) ([]model.Event, error) {
	d.mx.RLock()
	defer d.mx.RUnlock()
	var events []model.Event
	if _, ok := d.event[user_id]; !ok {
		return nil, fmt.Errorf("user not found")
	}
	for _, event := range d.event[user_id] {
		if event.Date.Year() == date.Year() && event.Date.Month() == date.Month() {
			events = append(events, event)
		}
	}
	return events, nil
}
