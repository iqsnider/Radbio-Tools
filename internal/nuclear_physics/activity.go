package nuclear_physics

import (
	"math"
)

type Isotope struct {
	Element        string
	Atomic_number  int
	Proton_number  int
	Neutron_number int
	Halflife       *float64
	Decayconstant  *float64
}

func NewIsotope(element string, atomic_number, proton_number int, optional ...float64) *Isotope {
	i := &Isotope{
		Element:        element,
		Atomic_number:  atomic_number,
		Proton_number:  proton_number,
		Neutron_number: atomic_number - proton_number,
	}

	// set halflife if provided
	if len(optional) > 0 {
		i.Halflife = &optional[0]
	}

	// set decayconstant if provided
	if len(optional) > 1 {
		i.Decayconstant = &optional[1]
	}

	return i
}

// calculate time to a given activity percentage
func (iso *Isotope) TimeToActivity(target_activity_percent float64) float64 {
	lambda := *iso.Halflife
	t := lambda * math.Log(target_activity_percent/100) / math.Log(0.5)

	return t
}
