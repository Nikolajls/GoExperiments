package types_experiments

import "math"

// Embedding in struct (and also possible in interfaces)
// Is the act of defining a base
type wingspan struct {
	sqft int
}

// Parachute struct embeds wingspan struct and thus has the property sqft under the wingspan base.
type Parachute struct {
	wingspan
	manufacturer ParachuteManufacturer
	Level        ParachuteLevel
	Elliptical   bool
}

type Skydiver struct {
	nakedWeightKg        int
	gearWeightKg         int
	jumpCount            int
	highPerformanceJumps int
}

// Methods
func (w wingspan) CalculateWingload(exitWeightKg int) (wingload float64, sqftLoadedWithGrams int) {
	return CalculateWingload(w.sqft, exitWeightKg)
}

// CalculateWingload calculates the load of a wing by getting the relation of weight being lifted by the wing as well as how many grams each sqft is loaded by.
func CalculateWingload(sqft int, exitWeightKg int) (wingload float64, sqftLoadedWithGrams int) {
	exitWeightGram := (exitWeightKg) * 1000
	exitWeightLbs := float64(exitWeightGram) / 453.59237
	sqftFloat := float64(sqft)

	loadPrSqft := int(float64(exitWeightGram) / sqftFloat)

	wingLoad := exitWeightLbs / sqftFloat
	wingLoad = math.Round(wingLoad*100) / 100

	return wingLoad, loadPrSqft
}
