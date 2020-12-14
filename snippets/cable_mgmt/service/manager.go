package service

import (
	"cable_mgmt/models"
	"fmt"
)

type Manager struct {
	Customer models.Customer
	Channels models.Channels
}

func (m *Manager) Manage() {
	var operation string
	fmt.Println("Do you want to subscribe or upgrade? <s/u> ")
	_, _ = fmt.Scanln(&operation)

	var channelName string
	fmt.Println("Enter Channel Type: ")
	_, _ = fmt.Scanln(&channelName)

	var packType models.PackType
	fmt.Printf("Please choose subscription type: (%+v)\n", m.Channels.GetPackNames(channelName))
	_, _ = fmt.Scanln(&packType)

	channel, err := m.Channels.Get(channelName, packType)
	if err != nil {
		fmt.Println(fmt.Errorf("requested channel=%v and package=%v not available!\n", channelName, packType))
		return
	}

	fmt.Printf("operation=%v channelName=%v packType=%v", operation, channelName, packType)

	fmt.Println("-----------------------------------------------------------------------")

	switch operation {
	case "s":
	case "S":
		fmt.Println("-----------------SUBSCRIBE--------------------")
		err = m.Customer.SubscribeTo(channel, packType)
		if err != nil {
			fmt.Println("sorry couldn't subscribe to requested channelName")
			fmt.Printf("Available account balance: %v INR", m.Customer.Balance)
			return
		}
		fmt.Printf("Channel %v (%v) subscribed successfully", channelName, packType)
		fmt.Printf("Available account balance: %v INR", m.Customer.Balance)
		break
	case "u":
	case "U":
		fmt.Println("-----------------UPGRADE--------------------")
		err = m.Customer.UpgradeTo(channel, packType)
		if err != nil {
			fmt.Println("sorry couldn't upgrade to requested channelName")
			fmt.Printf("Available account balance: %v INR", m.Customer.Balance)
			return
		}
		fmt.Printf("Channel %v (%v) upgraded successfully", channelName, packType)
		fmt.Printf("Available account balance: %v INR", m.Customer.Balance)
		break

	default:
		fmt.Println(fmt.Errorf("invalid operation selected"))
	}

}
