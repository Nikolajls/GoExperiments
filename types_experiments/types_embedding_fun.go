package types_experiments

import (
	"math"
)

// Embedding in struct (and also possible in interfaces)
// Is the act of defining a base
type wingspan struct {
	sqft int
}

// Airfoil struct embeds wingspan struct and thus has the property sqft under the wingspan base.
type Airfoil struct {
	wingspan
	manufacturer ParachuteManufacturer
}

// Methods
func (w wingspan) CalculateWingload(exitWeightLbs int) float64 {
	return CalculateWingload(w.sqft, exitWeightLbs)
}

func CalculateWingload(sqft int, exitWeightLbs int) float64 {
	wingLoad := float64(exitWeightLbs) / float64(sqft)
	wingLoad = math.Round(wingLoad*100) / 100
	return wingLoad
}
