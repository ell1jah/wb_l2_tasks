package handler

import (
	"d11/database"
	"d11/model"
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

type Handler struct {
	db *database.DataBase
}

func NewHandler(db *database.DataBase) *Handler {
	return &Handler{db: db}
}

func (h *Handler) InitRoutes() *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc("/create_event", Middlewarelog(h.Create_event))
	router.HandleFunc("/update_event", Middlewarelog(h.Update_event))
	router.HandleFunc("/delete_event", Middlewarelog(h.Delete_event))
	router.HandleFunc("/events_for_day", Middlewarelog(h.Events_for_day))
	router.HandleFunc("/events_for_week", Middlewarelog(h.Events_for_week))
	router.HandleFunc("/events_for_month", Middlewarelog(h.Events_for_month))
	return router
}

func (h *Handler) Create_event(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		errorresponse(w, "incorrect method")
		return
	}
	var event model.Event
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		errorresponse(w, err.Error())
		return
	}
	if err := h.db.CreateEvent(event); err != nil {
		errorresponse(w, err.Error())
		return
	}
	successresponse(w, "event created")
	return
}

func (h *Handler) Update_event(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		errorresponse(w, "incorrect method")
		return
	}
	var event model.Event
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		errorresponse(w, err.Error())
		return
	}
	if err := h.db.UpdateEvent(event); err != nil {
		errorresponse(w, err.Error())
		return
	}
	successresponse(w, "event updated")
	return
}
func (h *Handler) Delete_event(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		errorresponse(w, "incorrect method")
		return
	}
	var event model.Event
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		errorresponse(w, err.Error())
		return
	}
	if err := h.db.DeleteEvent(event); err != nil {
		errorresponse(w, err.Error())
		return
	}
	successresponse(w, "event deleted")
	return
}
func (h *Handler) Events_for_day(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		errorresponse(w, "incorrect method")
		return
	}
	user_id, err := strconv.Atoi(r.URL.Query().Get("user_id"))
	if err != nil {
		errorresponse(w, err.Error())
	}
	date, err := time.Parse("2006-01-02", r.URL.Query().Get("date"))
	if err != nil {
		errorresponse(w, err.Error())
	}
	events, err := h.db.Events_for_day(user_id, date)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	successresponse(w, events)
	return
}
func (h *Handler) Events_for_week(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		errorresponse(w, "incorrect method")
		return
	}
	user_id, err := strconv.Atoi(r.URL.Query().Get("user_id"))
	if err != nil {
		errorresponse(w, err.Error())
	}
	date, err := time.Parse("2006-01-02", r.URL.Query().Get("date"))
	if err != nil {
		errorresponse(w, err.Error())
	}
	events, err := h.db.Events_for_week(user_id, date)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	successresponse(w, events)
	return
}
func (h *Handler) Events_for_month(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		errorresponse(w, "incorrect method")
		return
	}
	user_id, err := strconv.Atoi(r.URL.Query().Get("user_id"))
	if err != nil {
		errorresponse(w, err.Error())
	}
	date, err := time.Parse("2006-01-02", r.URL.Query().Get("date"))
	if err != nil {
		errorresponse(w, err.Error())
	}
	events, err := h.db.Events_for_month(user_id, date)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	successresponse(w, events)
	return
}
