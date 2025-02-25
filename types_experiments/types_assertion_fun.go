package types_experiments

import (
	"fmt"
)

func CalculateWingLoadSafeTypeAssertion(arg interface{}, exitWeightLbs int) (float64, error) {
	if argType, ok := arg.(wingspan); ok {
		return argType.CalculateWingload(exitWeightLbs), nil
	}

	if v, ok := arg.(Airfoil); ok {
		return v.CalculateWingload(exitWeightLbs), nil
	}

	//// If the type has an Sqft() method we allow it to be called as well.
	if v, ok := arg.(interface{ Sqft() int }); ok {
		return CalculateWingload(v.Sqft(), exitWeightLbs), nil
	}

	if v, ok := arg.(int); ok {
		return CalculateWingload(v, exitWeightLbs), nil
	}

	return 0, fmt.Errorf("arg is not wingspan")
}

func CalculateWingLoadSwitchTypeAssertion(arg interface{}, exitWeightLbs int) (float64, error) {
	switch v := arg.(type) {
	case wingspan:
		return v.CalculateWingload(exitWeightLbs), nil
	case Airfoil:
		return v.CalculateWingload(exitWeightLbs), nil
	case interface{ Sqft() int }: // If the type has an Sqft() method we allow it to be called as well.
		return CalculateWingload(v.Sqft(), exitWeightLbs), nil
	case int:
		return CalculateWingload(v, exitWeightLbs), nil
	default:
		return 0, fmt.Errorf("arg is not wingspan")
	}
}

// CalculateWingLoadNotSafeTypeAssertion does a type assertion on the interface{} argument to a wingspan struct.
// It does it without an ok boolean check meaning if it is anything else it will panic
func CalculateWingLoadNotSafeTypeAssertion(arg interface{}, exitWeightLbs int) float64 {
	res := arg.(wingspan).CalculateWingload(exitWeightLbs)
	return res
}
