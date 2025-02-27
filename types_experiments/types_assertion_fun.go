package types_experiments

import (
	"fmt"
)

func CalculateWingLoadSafeTypeAssertion(arg interface{}, exitWeightKg int) (float64, error) {
	if argType, ok := arg.(wingspan); ok {
		load, _ := argType.CalculateWingload(exitWeightKg)
		return load, nil
	}

	if v, ok := arg.(Parachute); ok {
		load, _ := v.CalculateWingload(exitWeightKg)
		return load, nil
	}

	//// If the type has an Sqft() method we allow it to be called as well.
	if v, ok := arg.(interface{ Sqft() int }); ok {
		load, _ := CalculateWingload(v.Sqft(), exitWeightKg)
		return load, nil
	}

	if v, ok := arg.(int); ok {
		load, _ := CalculateWingload(v, exitWeightKg)
		return load, nil
	}

	return 0, fmt.Errorf("arg is not wingspan")
}

func CalculateWingLoadSwitchTypeAssertion(arg interface{}, exitWeightKg int) (float64, error) {
	switch v := arg.(type) {
	case wingspan:
		load, _ := v.CalculateWingload(exitWeightKg)
		return load, nil
	case Parachute:
		load, _ := v.CalculateWingload(exitWeightKg)
		return load, nil
	case interface{ Sqft() int }: // If the type has an Sqft() method we allow it to be called as well.
		load, _ := CalculateWingload(v.Sqft(), exitWeightKg)
		return load, nil
	case int:
		load, _ := CalculateWingload(v, exitWeightKg)
		return load, nil
	default:
		return 0, fmt.Errorf("arg is not wingspan")
	}

}

// CalculateWingLoadNotSafeTypeAssertion does a type assertion on the interface{} argument to a wingspan struct.
// It does it without an ok boolean check meaning if it is anything else it will panic
func CalculateWingLoadNotSafeTypeAssertion(arg interface{}, exitWeightKg int) float64 {
	res, _ := arg.(wingspan).CalculateWingload(exitWeightKg)
	return res
}
