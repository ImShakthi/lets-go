package examples_test

import (
	"lets-go/concurrency/examples"
	"testing"
)

func TestChannelFibonacci(t *testing.T) {
	examples.ChannelFibonacci()
}

func TestPingPongMessenger(t *testing.T) {
	examples.PingPongMessenger()
}

func TestChanneledPrinter(t *testing.T) {
	examples.ChanneledPrinter()
}

func TestImplementSelect(t *testing.T) {
	examples.ImplementSelect()
}

func TestImplementChannel(t *testing.T) {
	examples.ImplementChannel()
}

func TestTikTok(t *testing.T) {
	examples.TikTok()
}

func TestNamePrinter(t *testing.T) {
	examples.NamePrinter()
}
