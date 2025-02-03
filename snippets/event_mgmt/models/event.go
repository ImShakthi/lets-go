package models

import "time"

type Event struct {
	Type     string
	StartsAt time.Time
	EndsAt   time.Time
}
