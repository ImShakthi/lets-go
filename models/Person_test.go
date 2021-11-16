package models_test

import (
	"fmt"
	"lets-go/models"
	"testing"
	"time"
)

func TestPerson(t *testing.T) {
	t.Helper()

	p := models.NewPerson().
		WithName("Sakthivel").
		WithDob(time.Now()).
		WithAddr("Coimbatore").
		WithAge(27)

	fmt.Printf("%+v", p)

}
