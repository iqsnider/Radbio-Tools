package main

import (
	"fmt"
	"github.com/iqsnider/Radbio-Tools/internal/nuclear_physics"
)

func main() {
	isotope := nuclear_physics.NewIsotope("Mo", 99, 42, 65.924)
	fmt.Println(isotope.TimeToActivity(0.0001))

}
