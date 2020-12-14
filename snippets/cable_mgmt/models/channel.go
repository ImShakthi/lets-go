package models

import (
	"cable_mgmt/constants"
	"fmt"
	"strings"
)

type Channels struct {
	Channels []Channel
}

type Channel struct {
	Name             string
	Packs            []Pack
	AllowedLocations []string
	SelectedPack     PackType
}

func (c *Channels) GetPackNames(channelName string) []PackType {
	for _, channel := range c.Channels {
		if strings.ToUpper(channel.Name) == channelName {
			var packNames []PackType
			for _, pack := range channel.Packs {
				packNames = append(packNames, pack.Type)
			}
			return packNames
		}
	}
	return []PackType{}
}

func (c *Channels) Get(name string, packType PackType) (Channel, error) {
	for _, channel := range c.Channels {
		if strings.ToUpper(channel.Name) == strings.ToUpper(name) {
			if channel.IsPackAvailable(packType) {
				return channel, nil
			}
			return Channel{}, constants.ErrRequestedPackNotAvailable
		}
	}
	return Channel{}, constants.ErrRequestedChannelNotAvailable
}

func (c *Channel) SubscribeBy(packType PackType, balance float64) (Pack, error) {
	for _, pack := range c.Packs {
		if pack.Type == packType {
			if pack.Price <= balance {
				return pack, nil
			}
			return Pack{}, constants.ErrInsufficientBalance
		}
	}
	return Pack{}, constants.ErrRequestedPackNotAvailable
}

func (c *Channel) UpgradeBy(packType PackType, balance float64) (Pack, error) {
	fmt.Printf("\n>>>>%v , %v \n", packType, c.SelectedPack)
	if packType == c.SelectedPack {
		return Pack{}, constants.ErrPackAlreadySubscribed
	}

	for _, pack := range c.Packs {
		if pack.Type == packType {
			if c.FeeForUpgrade(packType) <= balance {
				return pack, nil
			}
			return Pack{}, constants.ErrInsufficientBalance
		}
	}
	return Pack{}, constants.ErrRequestedPackNotAvailable
}

func (c *Channel) GenerateChannelWith(packType PackType) Channel {
	return Channel{
		Name:             c.Name,
		Packs:            c.Packs,
		AllowedLocations: c.AllowedLocations,
		SelectedPack:     packType,
	}
}

func (c *Channel) IsPackAvailable(packType PackType) bool {
	if (Pack{}) == c.getPackBy(packType) {
		return false
	}
	return true
}

func (c *Channel) AllowedForAll() bool {
	return c.AllowedFor("ALL")
}

func (c *Channel) AllowedFor(loc string) bool {
	for _, l := range c.AllowedLocations {
		location := strings.ToUpper(l)
		if loc == location || "ALL" == location {
			return true
		}
	}
	return false
}

func (c *Channel) getPackBy(packType PackType) Pack {
	for _, pack := range c.Packs {
		if pack.Type == packType {
			return pack
		}
	}
	return Pack{}
}

func (c *Channel) FeeForUpgrade(upgradeTo PackType) float64 {
	selectedPack := c.getPackBy(c.SelectedPack)
	upgradeToPack := c.getPackBy(upgradeTo)

	return upgradeToPack.Price - selectedPack.Price
}
