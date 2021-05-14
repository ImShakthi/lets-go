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
