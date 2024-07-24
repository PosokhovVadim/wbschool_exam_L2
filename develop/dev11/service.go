package main

import "time"

// Пустая бизнес логика

type Service struct {
	// empty
}

func NewEventService() *Service {
	return &Service{}
}

func (e *Service) CreateEvent(userID int, title string, date time.Time) (int, error) {

	return 0, nil
}

func (e *Service) UpdateEvent(id int, userID int, title string, date time.Time) error {
	return nil
}

func (e *Service) DeleteEvent(id int) error {
	return nil
}

func (e *Service) GetEventsForDay(userID int, date time.Time) ([]Event, error) {
	return nil, nil
}

func (e *Service) GetEventsForWeek(userID int, date time.Time) ([]Event, error) {
	return nil, nil
}

func (e *Service) GetEventsForMonth(userID int, date time.Time) ([]Event, error) {
	return nil, nil
}
