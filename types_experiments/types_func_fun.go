package types_experiments

type IsParachuteAllowedForJumper func(jumper Skydiver, parachute Parachute) bool

func VerifyJumperCanJump(ruleset IsParachuteAllowedForJumper, jumper Skydiver, parachute Parachute) bool {
	canJump := ruleset(jumper, parachute)
	return canJump
}

func NotAllowedToJump(jumper Skydiver, parachute Parachute) bool {
	return false
}
func AllowedToJump(jumper Skydiver, parachute Parachute) bool {
	return true
}

func DFURuleset(jumper Skydiver, parachute Parachute) bool {
	if parachute.Level == Competition {
		// Requires 200 jumps with high performance which are only allowed from 600 jumps onwards
		return jumper.highPerformanceJumps >= 200 && jumper.jumpCount >= 800
	}

	if parachute.Level == HighPerformance {
		// Requires at least 600 jumps to start
		return jumper.jumpCount >= 600
	}

	if jumper.jumpCount >= 400 {
		// If not a Competition or high performance any load and size is allwoed
		return true
	}

	// At this point we have less than 400 jumps and if the size is below 120 its not allowed
	if parachute.sqft < 120 {
		return false
	}

	maxLoad := 500
	if jumper.jumpCount >= 200 && !parachute.Elliptical {
		// If we are above 200 jumps and parachute is not elliptical we can increase load at bit.
		maxLoad = 650
	}
	exitWeightKilo := jumper.nakedWeightKg + jumper.gearWeightKg

	_, loadPrSqft := parachute.CalculateWingload(exitWeightKilo)

	isBelowLimit := loadPrSqft < maxLoad
	return isBelowLimit
}
