package main

import (
	"fmt"
	"lets-go/snippets/event_mgmt/models"
	"time"
)

func main() {
	event := models.Event{
		Type:     "opening",
		StartsAt: time.Date(2025, 1, 15, 9, 0, 0, 0, time.UTC),
		EndsAt:   time.Date(2025, 1, 15, 17, 0, 0, 0, time.UTC),
	}
	fmt.Println(event)
}
