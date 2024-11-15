// You can edit this code!
// Click here and start typing.
package main

import (
	"GoExperiments/errors_experiments"
	"errors"
	"fmt"
)

func main() {

	fmt.Println("My Testing")
	testError(giveMeMyError(0))
	testError(giveMeMyError(1))
	testError(giveMeMyError(2))
	testError(giveMeMyError(3))
}

func testError(theError error) {
	fmt.Println("===========================================================================")
	if theError == nil {
		fmt.Printf("Test of error:\nType:%T\n", theError)
		return
	}

	fmt.Printf("Test of error:\nType:%T\nValue:%v\nError:%v\nTest results:\n", theError, theError, theError.Error())

	a1, b1 := errors_experiments.IsErrorMyWonderfulValueErrorByValueType(theError)
	fmt.Printf("IsErrorMyWonderfulValueErrorByValueType=%v - %v\n", a1, b1)

	a2, b2 := errors_experiments.IsErrorMyWonderfulValueErrorByPointerType(theError)
	fmt.Printf("IsErrorMyWonderfulValueErrorByPointerType=%v - %v\n", a2, b2)

	a3, b3 := errors_experiments.IsErrorMyWonderfulPointerErrorByPointerType(theError)
	fmt.Printf("IsErrorMyWonderfulPointerErrorByPointerType=%v - %v\n", a3, b3)

	a4, b4 := errors_experiments.IsErrorMyWonderfulPointerErrorByPointerToPointerType(theError)
	fmt.Printf("IsErrorMyWonderfulPointerErrorByPointerToPointerType=%v - %v\n", a4, b4)
}

func giveMeMyError(typeOfErrorToGive int) error {
	if typeOfErrorToGive == 0 {
		return errors_experiments.MyWonderfulValueError{Msg: "I use a value receiver"}
	} else if typeOfErrorToGive == 1 {
		return &errors_experiments.MyWonderfulPointerError{Msg: "I use a pointer receiver"}
	} else if typeOfErrorToGive == 2 {
		return errors.New("my random error")
	}
	return nil
}
