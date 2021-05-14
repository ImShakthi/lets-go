package models

import "fmt"

type Slots []string

type Sessions []Session

type Session struct {
	SessionId         string `json:"session_id"`
	Date              string `json:"date"`
	AvailableCapacity int    `json:"available_capacity"`
	MinAgeLimit       int    `json:"min_age_limit"`
	Vaccine           string `json:"vaccine"`
	Slots             Slots  `json:"slots"`
}

func (s Sessions) PrintInfo() string {
	result := ""
	for _, session := range s {
		result = result + session.PrintInfo()
	}
	return result
}

func (s Session) PrintInfo() string {
	return fmt.Sprintf("Date=%v, Available=%v, MinAge=%v, Vaccine=%v",
		s.Date, s.AvailableCapacity, s.MinAgeLimit, s.Vaccine)
}

func (s Session) IsSlotAvailable() bool {
	return s.AvailableCapacity > 0
}

func (s Slots) Parse() {
	for _, slot := range s {
		fmt.Printf("%v", slot)
	}
}
