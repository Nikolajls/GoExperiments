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

//Handling of MyWonderfulValueError which is a value type error

func IsErrorMyWonderfulValueErrorByValueType(theError error) (bool, string) {
	var targetV MyWonderfulValueError = MyWonderfulValueError{}
	if theError == nil {
		return false, ""
	}

	if errors.As(theError, &targetV) {
		return true, targetV.Error()
	}
	return false, ""
}

func IsErrorMyWonderfulValueErrorByPointerType(theError error) (bool, string) {
	if theError == nil {
		return false, ""
	}
	var targetValueErrorWithPointer *MyWonderfulValueError = &MyWonderfulValueError{}
	if errors.As(theError, &targetValueErrorWithPointer) {
		return true, targetValueErrorWithPointer.Error()
	}
	return false, ""
}

//Handling of MyWonderfulPointerError which is a pointer error

// IsErrorMyWonderfulPointerErrorByPointerType
// This can give run time errors that can be recovered. The error type is a pointer receiver and the type checking towards is a pointer to the type.
func IsErrorMyWonderfulPointerErrorByPointerType(theError error) (success bool, returning string) {
	defer func() {
		if r := recover(); r != nil {
			success = false
			returning = "PANIC RECOVERED" //fmt.Sprintf("%v", r)
		}
	}()

	//Doing this will give a runtime error resulting in a panic due to the pointer being to the struct.
	//And not a pointer to a pointer which is needed for pointer receiver errors
	//var targetP MyWonderfulPointerError = MyWonderfulPointerError{}
	//if errors.As(theError, &targetP) {
	//	return true, targetP.Error()
	//}

	//This works and is the way it needs to be which is implemented in the other method below
	var targetP *MyWonderfulPointerError = &MyWonderfulPointerError{}
	if errors.As(theError, &targetP) {
		return true, targetP.Error()
	}

	return false, ""
}

func IsErrorMyWonderfulPointerErrorByPointerToPointerType(theError error) (bool, string) {
	if theError == nil {
		return false, ""
	}
	var targetP *MyWonderfulPointerError = &MyWonderfulPointerError{} //Give me the pointer of the struct

	if errors.As(theError, &targetP) { // the target is **PtrError in this case ("a pointer to a pointer type")
		return true, targetP.Error()
	}
	return false, ""
}
