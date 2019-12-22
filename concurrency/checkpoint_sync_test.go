package concurrency_test

import (
	"lets-go/concurrency"
	"testing"
)

func TestCheckpointSync(t *testing.T) {
	concurrency.CheckpointSync()
}
