package atomic_physics

import (
	"strconv"
	"strings"
)

func ElectronicConfiguration(protonNum int) string {
	if protonNum <= 0 || protonNum > 118 {
		return "Invalid atomic number"
	}

	// filling order based on Aufbau principle
	type orbital struct {
		name         string
		maxElectrons int
	}

	orbitals := []orbital{
		{"1s", 2}, {"2s", 2}, {"2p", 6}, {"3s", 2}, {"3p", 6},
		{"4s", 2}, {"3d", 10}, {"4p", 6}, {"5s", 2}, {"4d", 10},
		{"5p", 6}, {"6s", 2}, {"4f", 14}, {"5d", 10}, {"6p", 6},
		{"7s", 2}, {"5f", 14}, {"6d", 10}, {"7p", 6},
	}

	var config strings.Builder
	remainingElectrons := protonNum

	for _, orb := range orbitals {
		if remainingElectrons <= 0 {
			break
		}

		electronsInOrbital := remainingElectrons
		if electronsInOrbital > orb.maxElectrons {
			electronsInOrbital = orb.maxElectrons
		}

		if config.Len() > 0 {
			config.WriteString(" ")
		}
		config.WriteString(orb.name + strconv.Itoa(electronsInOrbital))
		remainingElectrons -= electronsInOrbital
	}

	return config.String()
}
