package models

import (
	"cable_mgmt/constants"
	"strings"
)

type Customer struct {
	Name               string
	SubscribedChannels []Channel
	Balance            float64
}

func (c *Customer) SubscribeTo(channel Channel, selectedPackType PackType) error {
	_, err := c.getSubscribedChannel(channel.Name)
	if err == nil {
		return constants.ErrAlreadySubscribed
	}

	selectedPack, err := channel.SubscribeBy(selectedPackType, c.Balance)
	if err != nil {
		return err
	}

	c.Balance = c.Balance - selectedPack.Price
	ch := channel.GenerateChannelWith(selectedPackType)
	c.SubscribedChannels = append(c.SubscribedChannels, ch)
	return nil
}

func (c *Customer) UpgradeTo(channel Channel, selectedPackType PackType) error {
	subscribedChannel, err := c.getSubscribedChannel(channel.Name)
	if err != nil {
		return constants.ErrRequestedChannelIsNotSubscribed
	}

	_, err = subscribedChannel.UpgradeBy(selectedPackType, c.Balance)
	if err != nil {
		return err
	}

	c.Balance = c.Balance - subscribedChannel.FeeForUpgrade(selectedPackType)
	ch := subscribedChannel.GenerateChannelWith(selectedPackType)
	for i, sc := range c.SubscribedChannels {
		if sc.Name == channel.Name {
			c.SubscribedChannels[i] = ch
			break
		}
	}
	return nil
}

func (c *Customer) getSubscribedChannel(name string) (Channel, error) {
	for _, channel := range c.SubscribedChannels {
		if strings.ToUpper(channel.Name) == strings.ToUpper(name) {
			return channel, nil
		}
	}
	return Channel{}, constants.ErrRequestedChannelNotAvailable
}
