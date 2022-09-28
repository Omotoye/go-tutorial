package main

import (
	"fmt"
)

// creating an alias for fmt.Println
var pl = fmt.Println

func main() {
	// var myMap map [KeyType]valueType
	var heroes map[string]string
	heroes = make(map[string]string)

	// Another way of defining a map in go
	villians := make(map[string]string)

	// assigning values to keys
	heroes["Batman"] = "Bruce Wayne"
	heroes["Superman"] = "Clark Kent"
	heroes["The Flask"] = "Barry Allen"

	villians["Lex Luther"] = "Lex Luther"

	superPets := map[int]string{1: "Krypto", 2: "Bat Hound"}
	fmt.Printf("Batman is %v\n", heroes["Batman"])
	pl("Chip :", superPets[3])
	_, ok := superPets[3]
	pl("Is there a 3rd pet :", ok)

	// for loops for map
	for key, value := range heroes {
		fmt.Printf("%s is %s\n", key, value)
	}

	// removing an item in a map
	delete(heroes, "The Flash")

}
