package app

import (
    "time"

    "github.com/MarcosBrindis/hexagonal-api/internal/domain"
)

type EventService interface {
    DetectMovement() domain.Event
    DetectGasLeak() domain.Event
    DetectWindowOpen() domain.Event
}

type eventService struct{}

func NewEventService() EventService {
    return &eventService{}
}

func (s *eventService) DetectMovement() domain.Event {
    return domain.Event{Type: "Movement detected", Timestamp: time.Now()}
}

func (s *eventService) DetectGasLeak() domain.Event {
    return domain.Event{Type: "Gas leak detected", Timestamp: time.Now()}
}

func (s *eventService) DetectWindowOpen() domain.Event {
    return domain.Event{Type: "Window opened", Timestamp: time.Now()}
}