package types_experiments

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReturnTheValueRepeatedValue_WithString(t *testing.T) {
	val := "hello"
	got := ReturnTheValueRepeatedValueForStringAlias(val)
	assert.EqualValues(t, got, "hellohello")
}

func TestReturnTheValueRepeatedValue_WithStringAlias(t *testing.T) {
	val := stringAlias("hello")
	got := ReturnTheValueRepeatedValueForStringAlias(val)
	assert.EqualValues(t, got, "hellohello")
}

func TestReturnCoordinates_WithCoordinatesAliasBase(t *testing.T) {
	//Arrange
	x, y, z := 1, 2, 3
	coords := coordinatesForAliasBase{x, y, z}

	//Act
	gotX, gotY, gotZ := ReturnCoordinatesUsingAlias(coords)

	//Assert
	assert.Equalf(t, x, gotX, "ReturnCoordinatesUsingAlias(%v)", coords)
	assert.Equalf(t, y, gotY, "ReturnCoordinatesUsingAlias(%v)", coords)
	assert.Equalf(t, z, gotZ, "ReturnCoordinatesUsingAlias(%v)", coords)
}

func TestReturnCoordinates_WithCoordinatesAlias(t *testing.T) {
	//Arrange
	x, y, z := 1, 2, 3
	coords := coordinatesAliased{x, y, z}

	//Act
	gotX, gotY, gotZ := ReturnCoordinatesUsingAlias(coords)

	//Assert
	assert.Equalf(t, x, gotX, "ReturnCoordinatesUsingAlias(%v)", coords)
	assert.Equalf(t, y, gotY, "ReturnCoordinatesUsingAlias(%v)", coords)
	assert.Equalf(t, z, gotZ, "ReturnCoordinatesUsingAlias(%v)", coords)
}
