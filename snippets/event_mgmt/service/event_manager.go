package service

import (
	"lets-go/snippets/event_mgmt/models"
	"time"
)

type EventManager interface {
	GetAllEvents() ([]models.Event, error)
}

type eventManager struct{}

func NewEventManager() EventManager {
	return &eventManager{}
}

func (e *eventManager) GetAllEvents() ([]models.Event, error) {

	return []models.Event{
		{
			Type:     "opening",
			StartsAt: time.Date(2025, 1, 15, 9, 0, 0, 0, time.UTC),
			EndsAt:   time.Date(2025, 1, 15, 17, 0, 0, 0, time.UTC),
		},
	}, nil
}
