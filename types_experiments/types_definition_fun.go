package types_experiments

import "fmt"

// https://go.dev/blog/alias-names
type stringDefinition string

type coordinatesForDefinitionBase struct {
	x, y, z int
}

type coordinatesDefinition coordinatesForDefinitionBase

func ReturnCoordinatesUsingDefinition(coords coordinatesDefinition) (x, y, z int) {
	return coords.x, coords.y, coords.z
}

func ReturnTheValueRepeatedValueForStringDefinition(value stringDefinition) string {
	return fmt.Sprintf("%v%v", value, value)
}
