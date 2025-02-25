package types_experiments

import "fmt"

// https://go.dev/blog/alias-names
type stringAlias = string

type coordinatesForAliasBase struct {
	x, y, z int
}

type coordinatesAliased = coordinatesForAliasBase

func ReturnTheValueRepeatedValueForStringAlias(value stringAlias) string {
	return fmt.Sprintf("%v%v", value, value)
}

func ReturnCoordinatesUsingAlias(coords coordinatesAliased) (x, y, z int) {
	return coords.x, coords.y, coords.z
}
