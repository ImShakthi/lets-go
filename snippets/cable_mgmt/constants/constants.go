package constants

import "fmt"

var (
	ErrRequestedChannelNotAvailable    = fmt.Errorf("requested channel not available")
	ErrRequestedPackNotAvailable       = fmt.Errorf("requested pack not available")
	ErrInsufficientBalance             = fmt.Errorf("insufficent balance")
	ErrAlreadySubscribed               = fmt.Errorf("already subscribed to the channel")
	ErrPackAlreadySubscribed           = fmt.Errorf("already subscribed to the pack")
	ErrRequestedChannelIsNotSubscribed = fmt.Errorf("requested channel is not subscribed")
)
