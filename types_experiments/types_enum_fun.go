package types_experiments

// ParachuteManufacturer shows how to do an enum in Golang with the iota keyword being the key to increasing each int value for the const declaration.
type ParachuteManufacturer int

const (
	PerformanceDesigns ParachuteManufacturer = iota
	NZAeroSports
	FluidWings
	Jyro
	Unknown
)

// Despite the ParachuteManufacturer being defined as an int type it values cannot be used to access the map that is defined with the enumType ParachuteManufacturer
var manufacturersAbbs = map[ParachuteManufacturer]string{
	PerformanceDesigns: "PD",
	NZAeroSports:       "NZ",
	FluidWings:         "FW",
	Jyro:               "JY",
}

func (w Parachute) GetManufacturerAbbreviation() string {
	if abb, foundInMap := manufacturersAbbs[w.manufacturer]; foundInMap {
		return abb
	}

	return "NoAbbreviation"
}

func (w Parachute) GetManufacturerPriority() int {
	switch w.manufacturer {
	case PerformanceDesigns:
		return 0
	case NZAeroSports, FluidWings:
		return 1
	case Jyro:
		return 2
	default:
		return 3
	}
}

type ParachuteLevel int

const (
	Beginner ParachuteLevel = iota
	Novice
	Intermediary
	Advanced
	HighPerformance
	Competition
)
