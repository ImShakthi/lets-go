package hello

import (
	"testing"
)

func Test_Hello(t *testing.T) {
	want := "Hello, world."
	if got := Say(); got != want {
		t.Errorf("Say() want = '%s'  got = '%s' ", want, got)
	}
}

func Test_Proverb(t *testing.T) {
	want := "Concurrency is not parallelism."
	if got := Proverb(); got != want {
		t.Errorf("Proverb() want = '%s'  got = '%s' ", want, got)
	}
}
