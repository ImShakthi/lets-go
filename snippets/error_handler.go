package snippets

import (
	"fmt"
	"github.com/pkg/errors"
)

type ErrorHandler interface {
	Play()
}

type errorHandler struct {
}

func NewErrorHandler() ErrorHandler {
	return &errorHandler{}
}

func (e *errorHandler) Play() {
	err := throwError()
	if err != nil {
		fmt.Println(err)
	}
	err2 := wrapError()
	if err2 != nil {
		fmt.Println(err2)
	}

	err3 := causeError()
	if err3 != nil {
		fmt.Print(err3)
	}
}

func throwError() error {
	err := errors.New(" go error")
	return fmt.Errorf("oh this is a error: %+v", err)
}

func wrapError() error {
	err := errors.New("main error ")
	return errors.Wrap(err, "added to wrap")
}

func causeError() error {
	err := errors.New("cause error ")
	return errors.Cause(err)
}
