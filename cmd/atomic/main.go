package main

import (
	"fmt"
	"github.com/iqsnider/Radbio-Tools/internal/atomic_physics"
)

func main() {
	var i int

	fmt.Print("Type a number: ")
	fmt.Scan(&i)
	fmt.Println(atomic_physics.ElectronicConfiguration(i))

}
