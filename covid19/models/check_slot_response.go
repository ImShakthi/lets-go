package models

import "fmt"

type CheckSlotResponse struct {
	Centers []Centers `json:"centers,omitempty"`
}

func (csr CheckSlotResponse) GetSlotAvailHospital() {
	slotsAvailable := false
	for _, center := range csr.Centers {
		if center.Sessions[0].IsSlotAvailable() {
			fmt.Printf("%v", center.PrintInfo())
			slotsAvailable = true
		}
	}
	fmt.Println("\nSlots available=", slotsAvailable)
}

func (csr CheckSlotResponse) GetResults() Result {
	slotInfo := make([]string, 0)
	slotsOpenedDates := make(map[string]string)
	for _, center := range csr.Centers {
		for _, session := range center.Sessions {
			slotsOpenedDates[session.Date] = session.Date
			if session.IsSlotAvailable() {
				info := fmt.Sprintf("Name::%s >> %s", center.Name, session.PrintInfo())
				slotInfo = append(slotInfo, info)
			}
		}
	}
	fmt.Println("No of days slot opened:: #", len(slotsOpenedDates))

	dates := make([]string, 0)
	for _, date := range slotsOpenedDates {
		dates = append(dates, date)
	}
	return Result{
		SlotDates: dates,
		SlotInfo:  slotInfo,
	}
}
