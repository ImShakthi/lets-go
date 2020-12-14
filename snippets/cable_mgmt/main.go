package main

import (
	"cable_mgmt/models"
	"cable_mgmt/service"
	"fmt"
)

func main() {
	fmt.Println("Cable management")
	availableChannels := getChannels()
	fmt.Printf("-> %+v", availableChannels)

	customer := models.Customer{
		Name:    "John Doe",
		Balance: 1000,
	}
	fmt.Println("Customer Details:")
	fmt.Printf("%+v", customer)

	manager := service.Manager{
		Customer: customer,
		Channels: availableChannels,
	}

	toContinue := "y"
	for toContinue == "y" {
		manager.Manage()
		fmt.Println("Continue? <y/n> ")
		_, _ = fmt.Scanln(&toContinue)
	}
}

func getChannels() models.Channels {
	return models.Channels{
		Channels: []models.Channel{
			{
				Name: "ABC",
				Packs: []models.Pack{
					{
						Type:  models.PT_Base,
						Price: 120,
					},
					{
						Type:  models.PT_SD,
						Price: 130,
					},
					{
						Type:  models.PT_HD,
						Price: 150,
					},
					{
						Type:  models.PT_4K,
						Price: 200,
					},
				},
				AllowedLocations: []string{"USA", "Canada"},
			},
			{
				Name: "PQR",
				Packs: []models.Pack{
					{
						Type:  models.PT_Base,
						Price: 120,
					},
					{
						Type:  models.PT_SD,
						Price: 130,
					},
				},
				AllowedLocations: []string{"ALL"},
			},
		},
	}
}
