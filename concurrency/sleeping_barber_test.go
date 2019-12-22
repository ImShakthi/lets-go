package concurrency_test

import (
	"lets-go/concurrency"
	"testing"
)

func TestSleepingBarber(t *testing.T) {
	concurrency.SleepingBarber()
}
