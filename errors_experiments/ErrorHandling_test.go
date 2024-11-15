// You can edit this code!
// Click here and start typing.
package errors_experiments

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsErrorMyWonderfulValueErrorByValueType(t *testing.T) {
	var tests = []struct {
		theError             error
		expectedResult       bool
		expectedResultString string
	}{
		{MyWonderfulValueError{Msg: "I use a value receiver"}, true, "I use a value receiver"},
		{&MyWonderfulValueError{Msg: "I use a value receiver but am a pointer"}, false, ""},
		{&MyWonderfulPointerError{Msg: "I use a pointer receiver"}, false, ""},
		{errors.New("my random error"), false, ""},
		{nil, false, ""},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("%v", tt.theError)
		t.Run(testName, func(t *testing.T) {
			result, resultString := IsErrorMyWonderfulValueErrorByValueType(tt.theError)
			assert.Equal(t, tt.expectedResult, result, testName)
			assert.Equal(t, tt.expectedResultString, resultString, testName)

		})
	}
}

func TestIsErrorMyWonderfulValueErrorByPointerType(t *testing.T) {
	var tests = []struct {
		theError             error
		expectedResult       bool
		expectedResultString string
	}{
		//Despite the error being a value receiver using a pointer to one, and the impl testing with a pointer to a pointer it works.

		{&MyWonderfulValueError{Msg: "I use a value receiver but am a pointer"}, true, "I use a value receiver but am a pointer"},
		{MyWonderfulValueError{Msg: "I use a value receiver"}, false, ""},
		{&MyWonderfulPointerError{Msg: "I use a pointer receiver"}, false, ""},
		{errors.New("my random error"), false, ""},
		{nil, false, ""},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("%v", tt.theError)
		t.Run(testName, func(t *testing.T) {
			result, resultString := IsErrorMyWonderfulValueErrorByPointerType(tt.theError)
			assert.Equal(t, tt.expectedResult, result, testName)
			assert.Equal(t, tt.expectedResultString, resultString, testName)
		})
	}
}

func TestIsErrorMyWonderfulPointerErrorByPointerToPointerType(t *testing.T) {
	var tests = []struct {
		theError             error
		expectedResult       bool
		expectedResultString string
	}{

		{&MyWonderfulPointerError{Msg: "I use a pointer receiver"}, true, "I use a pointer receiver"},
		{&MyWonderfulValueError{Msg: "I use a value receiver but am a pointer"}, false, ""},
		{MyWonderfulValueError{Msg: "I use a value receiver"}, false, ""},
		{errors.New("my random error"), false, ""},
		{nil, false, ""},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("%v", tt.theError)
		t.Run(testName, func(t *testing.T) {
			result, resultString := IsErrorMyWonderfulPointerErrorByPointerToPointerType(tt.theError)
			assert.Equal(t, tt.expectedResult, result, testName)
			assert.Equal(t, tt.expectedResultString, resultString, testName)
		})
	}
}

func TestIsErrorMyWonderfulPointerErrorByPointerType(t *testing.T) {
	var tests = []struct {
		theError             error
		expectedResult       bool
		expectedResultString string
	}{

		{&MyWonderfulPointerError{Msg: "I use a pointer receiver"}, true, "I use a pointer receiver"},
		{&MyWonderfulValueError{Msg: "I use a value receiver but am a pointer"}, false, ""},
		{MyWonderfulValueError{Msg: "I use a value receiver"}, false, ""},
		{errors.New("my random error"), false, ""},
		{nil, false, ""},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("%v", tt.theError)
		t.Run(testName, func(t *testing.T) {
			result, resultString := IsErrorMyWonderfulPointerErrorByPointerType(tt.theError)
			assert.Equal(t, tt.expectedResult, result, testName)
			assert.Equal(t, tt.expectedResultString, resultString, testName)
		})
	}
}
