package errors_experiments

import (
	"errors"
	"fmt"
)

type MyWonderfulValueError struct {
	Msg string
}

func (e MyWonderfulValueError) Error() string {
	s := fmt.Sprintf("%v", e.Msg)
	return s
}

type MyWonderfulPointerError struct {
	Msg string
}

func (e *MyWonderfulPointerError) Error() string {
	s := fmt.Sprintf("%v", e.Msg)
	return s
}

func giveMeMyError(typeOfErrorToGive int) error {
	if typeOfErrorToGive == 0 {
		return MyWonderfulValueError{Msg: "I use a value receiver"}
	} else if typeOfErrorToGive == 1 {
		return &MyWonderfulPointerError{Msg: "I use a pointer receiver"}
	} else if typeOfErrorToGive == 2 {
		return errors.New("my random error")
	}
	return nil
}
