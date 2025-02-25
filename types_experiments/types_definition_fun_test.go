package types_experiments

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//This test does not compile, and that is because type definition does not allow to substitute the types, that is what type aliasing allows to do.
//func TestReturnTheValueRepeatedValueForStringDefinition_WithString(t *testing.T) {
//	val := "hello"
//	got := ReturnTheValueRepeatedValueForStringDefinition(val)
//	assert.EqualValues(t, got, "hellohello")
//}

func TestReturnTheValueRepeatedValueForStringDefinition_WithStringDefinition(t *testing.T) {
	val := stringDefinition("hello")
	got := ReturnTheValueRepeatedValueForStringDefinition(val)
	assert.EqualValues(t, got, "hellohello")
}

//This test wont compile due to the definition not being the same type as the base it uses
//func TestReturnCoordinates_WithCoordinatesDefinitionBase(t *testing.T) {
//	//Arrange
//	x, y, z := 1, 2, 3
//	coords := coordinatesForDefinitionBase{x, y, z}
//
//	//Act
//	gotX, gotY, gotZ := ReturnCoordinatesUsingDefinition(coords)
//
//	//Assert
//	assert.Equalf(t, x, gotX, "ReturnCoordinatesUsingAlias(%v)", coords)
//	assert.Equalf(t, y, gotY, "ReturnCoordinatesUsingAlias(%v)", coords)
//	assert.Equalf(t, z, gotZ, "ReturnCoordinatesUsingAlias(%v)", coords)
//}

func TestReturnCoordinates_WithCoordinatesDefinition(t *testing.T) {
	//Arrange
	x, y, z := 1, 2, 3
	coords := coordinatesDefinition{x, y, z}

	//Act
	gotX, gotY, gotZ := ReturnCoordinatesUsingDefinition(coords)

	//Assert
	assert.Equalf(t, x, gotX, "ReturnCoordinatesUsingDefinition(%v)", coords)
	assert.Equalf(t, y, gotY, "ReturnCoordinatesUsingDefinition(%v)", coords)
	assert.Equalf(t, z, gotZ, "ReturnCoordinatesUsingDefinition(%v)", coords)
}
