package errors_experiments

import (
	"errors"
	"fmt"
)

func RunErrorsAsTests() {
	runErrorsAs(giveMeMyError(0))
	runErrorsAs(giveMeMyError(1))
	runErrorsAs(giveMeMyError(2))
	runErrorsAs(giveMeMyError(3))
}

func runErrorsAs(theError error) {
	fmt.Println("===========================================================================")
	if theError == nil {
		fmt.Printf("Test of error:\nType:%T\n", theError)
		return
	}

	fmt.Printf("Test of error:\nType:%T\nValue:%v\nError:%v\nTest results:\n", theError, theError, theError.Error())

	a1, b1 := errorAsMyWonderfulValueErrorByValueType(theError)
	fmt.Printf("IsErrorMyWonderfulValueErrorByValueType=%v - %v\n", a1, b1)

	a2, b2 := errorAsMyWonderfulValueErrorByPointerType(theError)
	fmt.Printf("IsErrorMyWonderfulValueErrorByPointerType=%v - %v\n", a2, b2)

	a3, b3 := errorAsMyWonderfulPointerErrorByPointerType(theError)
	fmt.Printf("IsErrorMyWonderfulPointerErrorByPointerType=%v - %v\n", a3, b3)

	a4, b4 := errorAsMyWonderfulPointerErrorByPointerToPointerType(theError)
	fmt.Printf("IsErrorMyWonderfulPointerErrorByPointerToPointerType=%v - %v\n", a4, b4)
}

//Handling of MyWonderfulValueError which is a value type error

func errorAsMyWonderfulValueErrorByValueType(theError error) (bool, string) {
	var targetV MyWonderfulValueError = MyWonderfulValueError{}
	if theError == nil {
		return false, ""
	}

	if errors.As(theError, &targetV) {
		return true, targetV.Error()
	}
	return false, ""
}

func errorAsMyWonderfulValueErrorByPointerType(theError error) (bool, string) {
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
func errorAsMyWonderfulPointerErrorByPointerType(theError error) (success bool, returning string) {
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

func errorAsMyWonderfulPointerErrorByPointerToPointerType(theError error) (bool, string) {
	if theError == nil {
		return false, ""
	}
	var targetP *MyWonderfulPointerError = &MyWonderfulPointerError{} //Give me the pointer of the struct

	if errors.As(theError, &targetP) { // the target is **PtrError in this case ("a pointer to a pointer type")
		return true, targetP.Error()
	}
	return false, ""
}
