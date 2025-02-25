package errors_experiments

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMakeCupOfCoffee(t *testing.T) {
	type args struct {
		throwErrorDueToWater bool
		throwErrorDueToPower bool
		throwNoCoffeeError   bool
	}
	tests := []struct {
		name    string
		args    args
		wantR   string
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "success",
			args: args{
				throwErrorDueToWater: false,
				throwErrorDueToPower: false,
				throwNoCoffeeError:   false,
			},
			wantR:   "KAFFE",
			wantErr: assert.NoError,
		},
		{
			name: "Fails due to missing water",
			args: args{
				throwErrorDueToWater: true,
				throwErrorDueToPower: false,
				throwNoCoffeeError:   false,
			},
			wantR: "",
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				assert.Error(t, err, "Expected  error but got nil")
				return true
			},
		},
		{
			name: "Fails due to missing power",
			args: args{
				throwErrorDueToWater: false,
				throwErrorDueToPower: true,
				throwNoCoffeeError:   false,
			},
			wantR: "",
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				assert.Error(t, err, "Expected  error but got nil")
				return true
			},
		},
		{
			name: "Fails due to missing coffee",
			args: args{
				throwErrorDueToWater: false,
				throwErrorDueToPower: false,
				throwNoCoffeeError:   true,
			},
			wantR: "",
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				assert.Error(t, err, "Expected  error but got nil")

				//The error returned from the method, but it wraps the errors from further down the chain.
				var a *UnableToMakeCoffeeError
				validatedErrorTopAs := errors.As(err, &a)
				assert.Equal(t, true, validatedErrorTopAs)
				assert.EqualError(t, a, "failed to add coffee")

				//The error is cascaded up and wrapped.
				var ae *NoCoffeeError
				validatedErrorAs := errors.As(err, &ae)
				assert.Equal(t, true, validatedErrorAs)
				assert.EqualError(t, ae, "no coffee beans_no coffee beans")

				return true
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			gotR, err := MakeCupOfCoffee(tt.args.throwErrorDueToWater, tt.args.throwErrorDueToPower, tt.args.throwNoCoffeeError)
			tt.wantErr(t, err, fmt.Sprintf("MakeCupOfCoffee(%v, %v, %v)", tt.args.throwErrorDueToWater, tt.args.throwErrorDueToPower, tt.args.throwNoCoffeeError))

			assert.Equalf(t, tt.wantR, gotR, "MakeCupOfCoffee(%v, %v, %v)", tt.args.throwErrorDueToWater, tt.args.throwErrorDueToPower, tt.args.throwNoCoffeeError)
		})
	}
}

func TestBoilWater(t *testing.T) {
	type args struct {
		throwErrorDueToWater bool
		throwErrorDueToPower bool
	}
	tests := []struct {
		name      string
		args      args
		wantWater string
		wantErr   assert.ErrorAssertionFunc
	}{
		{
			name: "Success",
			args: args{
				throwErrorDueToWater: false,
				throwErrorDueToPower: false,
			},
			wantWater: "WATER",
			wantErr:   assert.NoError,
		}, {
			name: "Failed due to water",
			args: args{
				throwErrorDueToWater: true,
				throwErrorDueToPower: false,
			},
			wantWater: "",
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				assert.Error(t, err, "Expected an error but got nil")

				//Validate the pointer type with type assertion
				myErr, errIsValidated := err.(*NoWaterError)
				assert.Equal(t, true, errIsValidated)
				assert.EqualError(t, myErr, "no water")

				errIsValidated = errors.Is(err, &NoWaterError{
					Msg: "no water"})
				assert.Equal(t, true, errIsValidated)
				assert.EqualError(t, err, "no water")

				//Validate the pointer type error with ErrorAs
				var ae *NoWaterError
				validatedErrorAs := errors.As(err, &ae)
				assert.Equal(t, true, validatedErrorAs)
				assert.EqualError(t, ae, "no water")

				return true
			},
		}, {
			name: "Failed due to no power value error",
			args: args{
				throwErrorDueToWater: false,
				throwErrorDueToPower: true,
			},
			wantWater: "",
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				assert.Error(t, err, "Expected an error but got nil")

				//Validate the Value type error with Errors.Is
				errIsValidated := errors.Is(err, NoPowerError{
					Msg: "no power"})
				assert.Equal(t, true, errIsValidated)
				assert.EqualError(t, err, "no power#no power")

				//Validate the value type error with ErrorAs
				var targetValueError NoPowerError = NoPowerError{}
				errAsValidated := errors.As(err, &targetValueError)
				assert.Equal(t, true, errAsValidated)
				assert.EqualError(t, targetValueError, "no power#no power")

				return true
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			gotWater, err := BoilWater(tt.args.throwErrorDueToWater, tt.args.throwErrorDueToPower)

			tt.wantErr(t, err, fmt.Sprintf("BoilWater(%v, %v)", tt.args.throwErrorDueToWater, tt.args.throwErrorDueToPower))

			assert.Equalf(t, tt.wantWater, gotWater, "BoilWater(%v, %v)", tt.args.throwErrorDueToWater, tt.args.throwErrorDueToPower)
		})
	}
}

func TestAddCoffeeToWater(t *testing.T) {
	type args struct {
		throwNoCoffeeError bool
	}
	tests := []struct {
		name            string
		args            args
		wantAddedCoffee string
		wantErr         assert.ErrorAssertionFunc
	}{
		{
			name:            "Success",
			args:            args{throwNoCoffeeError: false},
			wantAddedCoffee: "KAFFE",
			wantErr:         assert.NoError,
		}, {
			name:            "Fails",
			args:            args{throwNoCoffeeError: true},
			wantAddedCoffee: "",
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				assert.Error(t, err, "Expected an error but got nil")

				//Assert the specific error type with As
				var ae *NoCoffeeError
				isOutOfCoffee := errors.As(err, &ae)
				assert.Equal(t, true, isOutOfCoffee)
				assert.EqualError(t, ae, "no coffee beans_no coffee beans")

				return false
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotAddedCoffee, err := AddCoffeeToWater(tt.args.throwNoCoffeeError)
			tt.wantErr(t, err, fmt.Sprintf("AddCoffeeToWater(%v)", tt.args.throwNoCoffeeError))

			assert.Equalf(t, tt.wantAddedCoffee, gotAddedCoffee, "AddCoffeeToWater(%v)", tt.args.throwNoCoffeeError)
		})
	}
}
