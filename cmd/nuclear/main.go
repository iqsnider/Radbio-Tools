package main

import (
	"fmt"
	"github.com/iqsnider/Radbio-Tools/internal/nuclear_physics"
)

func main() {
	isotope := nuclear_physics.NewIsotope("I", 125, 53, 59.407)
	fmt.Println(isotope.TimeToActivity(0.0001))

}
