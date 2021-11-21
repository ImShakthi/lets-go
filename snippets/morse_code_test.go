package snippets

import (
	"testing"
)

func TestEncodeToMorseCode(t *testing.T) {
	actual := encodeToMorseCode("SAKTHIVEL")

	expected := "._ _... _._. _.."
	if actual != expected {
		t.Errorf("expected: %s VS actual: %s", expected, actual)
	}
}

func TestDecodeMorseCode(t *testing.T) {
	actual := decodeMorseCode("._ _... _._. _..")

	expected := "ABCD"
	if actual != expected {
		t.Errorf("expected: %s VS actual: %s", expected, actual)
	}
}
