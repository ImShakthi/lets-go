package concurrency_test

import (
	"lets-go/concurrency"
	"testing"
)

func TestChannelFibonacci(t *testing.T) {
	concurrency.ChannelFibonacci()
}

func TestPingPongMessenger(t *testing.T) {
	concurrency.PingPongMessenger()
}

func TestChanneledPrinter(t *testing.T) {
	concurrency.ChanneledPrinter()
}

func TestImplementSelect(t *testing.T) {
	concurrency.ImplementSelect()
}

func TestImplementChannel(t *testing.T) {
	concurrency.ImplementChannel()
}

func TestTikTok(t *testing.T) {
	concurrency.TikTok()
}

func TestNamePrinter(t *testing.T) {
	concurrency.NamePrinter()
}
