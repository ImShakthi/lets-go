package service

import "testing"

func TestEventManager_GetAllEvents(t *testing.T) {
	eventManager := NewEventManager()

	actual, actualErr := eventManager.GetAllEvents()

	if actualErr != nil {
		t.Error(actualErr)
	}

	if len(actual) != 1 {
		t.Error("expected 1 event")
	}
	if actual[0].Type != "opening" {
		t.Error("expected event type is opening")
	}
}
