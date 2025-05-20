package sandbox

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStudentService_CreateStudent(t *testing.T) {
	studentService := NewStudentService()

	expected := Student{"sakthi", 16}
	actual := studentService.CreateStudent("sakthi", 16)

	if actual.Name != "sakthi" || actual.Age != 16 {
		t.Errorf("error in creating student")
	}

	assert.Equal(t, expected, actual)
}
