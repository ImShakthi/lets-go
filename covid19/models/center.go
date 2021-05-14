package models

import "fmt"

type Centers struct {
	CenterId     int     `json:"center_id"`
	Name         string  `json:"name"`
	Address      string    `json:"address"`
	StateName    string    `json:"state_name"`
	DistrictName string    `json:"district_name"`
	BlockName    string    `json:"block_name"`
	Pincode      int       `json:"pincode"`
	Lat          int       `json:"lat"`
	Long         int       `json:"long"`
	From         string    `json:"from"`
	To           string    `json:"to"`
	FeeType      string    `json:"fee_type"`
	Sessions     Sessions `json:"sessions,omitempty"`
}

func (c Centers) PrintInfo() string {
	return fmt.Sprintf("\nName=%v \n %v", c.Name, c.Sessions.PrintInfo())
}
