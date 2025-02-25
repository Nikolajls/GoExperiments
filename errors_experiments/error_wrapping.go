package errors_experiments

import "fmt"

// UnableToMakeCoffeeError is a pointer receiver error that wraps other errors
type UnableToMakeCoffeeError struct {
	Msg string
	Err error
}

func (e *UnableToMakeCoffeeError) Error() string {
	return fmt.Sprintf("%s", e.Msg)
}

func (e *UnableToMakeCoffeeError) Unwrap() error {
	return e.Err
}

func MakeCupOfCoffee(throwErrorDueToWater bool, throwErrorDueToPower bool, throwNoCoffeeError bool) (r string, err error) {
	_, err = BoilWater(throwErrorDueToWater, throwErrorDueToPower)
	if err != nil {
		return "", &UnableToMakeCoffeeError{
			Msg: "making a cup of boiling water failed",
			Err: err,
		}
	}

	coffeeAdded, err := AddCoffeeToWater(throwNoCoffeeError)
	if err != nil {
		return "", &UnableToMakeCoffeeError{
			Msg: "failed to add coffee",
			Err: err,
		}
	}

	return coffeeAdded, nil
}

// NoPowerError  Is a value receiver error that indicates no power for the boiling machine
type NoPowerError struct {
	Msg string
}

func (e NoPowerError) Error() string {
	s := fmt.Sprintf("%v#%v", e.Msg, e.Msg)
	return s
}

// NoWaterError is a pointer receiver error that indicates no water to boil
type NoWaterError struct {
	Msg string
}

func (e *NoWaterError) Error() string {
	s := fmt.Sprintf("%v", e.Msg)
	return s
}

// Is this enables the cal of Errors.Is to work with pointer receiver due to this line in wrap.go
// if x, ok := err.(interface{ Is(error) bool }); ok && x.Is(target) {
func (e *NoWaterError) Is(target error) bool {
	t, ok := target.(*NoWaterError)
	if !ok {
		return false
	}
	return e.Msg == t.Msg
}

func BoilWater(throwErrorDueToWater bool, throwErrorDueToPower bool) (water string, err error) {
	if throwErrorDueToWater {
		return "", &NoWaterError{
			Msg: "no water",
		}
	}

	if throwErrorDueToPower {
		return "", NoPowerError{
			Msg: "no power",
		}
	}
	return "WATER", nil
}

// NoCoffeeError is a pointer receiver error that indicates no coffee to add to water.
type NoCoffeeError struct {
	Msg string
}

func (e *NoCoffeeError) Error() string {
	s := fmt.Sprintf("%v_%v", e.Msg, e.Msg)
	return s
}

func AddCoffeeToWater(throwNoCoffeeError bool) (addedCoffee string, err error) {
	if throwNoCoffeeError {
		return "", &NoCoffeeError{
			Msg: "no coffee beans",
		}
	}
	return "KAFFE", nil
}
