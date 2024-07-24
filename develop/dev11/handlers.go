package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"
)

type EventService interface {
	CreateEvent(userID int, title string, date time.Time) (int, error)
	UpdateEvent(id int, userID int, title string, date time.Time) error
	DeleteEvent(id int) error
	GetEventsForDay(userID int, date time.Time) ([]Event, error)
	GetEventsForWeek(userID int, date time.Time) ([]Event, error)
	GetEventsForMonth(userID int, date time.Time) ([]Event, error)
}

type Server struct {
	eventService EventService
}

func NewServer(eventService EventService) *Server {
	return &Server{eventService: eventService}
}

func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(os.Stdout, "request: %v %v\n", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func (s *Server) createEvent(w http.ResponseWriter, r *http.Request) {
	params, err := GetParams(r)

	if err != nil {
		WriteJSON(w, http.StatusBadRequest, err)
		return
	}

	err = ValidateParams(params, []string{"user_id", "title", "date"})
	if err != nil {
		WriteJSON(w, http.StatusBadRequest, err)
		return
	}

	userID, _ := strconv.Atoi(params["user_id"])
	date, _ := time.Parse("2006-01-02", params["date"])
	id, err := s.eventService.CreateEvent(userID, params["title"], date)

	if err != nil {
		WriteJSON(w, http.StatusServiceUnavailable, err)
		return
	}

	WriteJSON(w, http.StatusOK, id)
}

func (s *Server) updateEvent(w http.ResponseWriter, r *http.Request) {
	params, err := GetParams(r)

	if err != nil {
		WriteJSON(w, http.StatusBadRequest, err)
		return
	}

	err = ValidateParams(params, []string{"user_id", "id"})
	if err != nil {
		WriteJSON(w, http.StatusBadRequest, err)
		return
	}

	userID, _ := strconv.Atoi(params["user_id"])
	id, _ := strconv.Atoi(params["id"])
	date, _ := time.Parse("2006-01-02", params["date"])

	err = s.eventService.UpdateEvent(id, userID, params["title"], date)

	if err != nil {
		WriteJSON(w, http.StatusServiceUnavailable, err)
		return
	}

	WriteJSON(w, http.StatusOK, nil)
}

func (s *Server) deleteEvent(w http.ResponseWriter, r *http.Request) {
	params, err := GetParams(r)

	if err != nil {
		WriteJSON(w, http.StatusBadRequest, err)
		return
	}

	id, err := ValidateInt(params["id"])
	if err != nil {
		WriteJSON(w, http.StatusBadRequest, err)
		return
	}

	err = s.eventService.DeleteEvent(id)
	if err != nil {
		WriteJSON(w, http.StatusServiceUnavailable, err)
		return
	}

	WriteJSON(w, http.StatusOK, nil)
}

func (s *Server) eventsForDay(w http.ResponseWriter, r *http.Request) {
	params, err := GetParams(r)

	if err != nil {
		WriteJSON(w, http.StatusBadRequest, err)
		return
	}

	userID, _ := strconv.Atoi(params["user_id"])
	date, err := ValidateDate(params["date"])
	if err != nil {
		WriteJSON(w, http.StatusBadRequest, err)
		return
	}

	events, err := s.eventService.GetEventsForDay(userID, date)
	if err != nil {
		WriteJSON(w, http.StatusServiceUnavailable, err)
		return
	}

	WriteJSON(w, http.StatusOK, events)

}

func (s *Server) eventsForWeek(w http.ResponseWriter, r *http.Request) {
	params, err := GetParams(r)

	if err != nil {
		WriteJSON(w, http.StatusBadRequest, err)
		return
	}

	userID, _ := strconv.Atoi(params["user_id"])
	date, err := ValidateDate(params["date"])
	if err != nil {
		WriteJSON(w, http.StatusBadRequest, err)
		return
	}

	events, err := s.eventService.GetEventsForWeek(userID, date)
	if err != nil {
		WriteJSON(w, http.StatusServiceUnavailable, err)
		return
	}

	WriteJSON(w, http.StatusOK, events)

}

func (s *Server) eventsForMonth(w http.ResponseWriter, r *http.Request) {
	params, err := GetParams(r)

	if err != nil {
		WriteJSON(w, http.StatusBadRequest, err)
		return
	}

	userID, _ := strconv.Atoi(params["user_id"])
	date, err := ValidateDate(params["date"])
	if err != nil {
		WriteJSON(w, http.StatusBadRequest, err)
		return
	}

	events, err := s.eventService.GetEventsForMonth(userID, date)
	if err != nil {
		WriteJSON(w, http.StatusServiceUnavailable, err)
		return
	}

	WriteJSON(w, http.StatusOK, events)
}
