package main

import (
	"fmt"
	"reflect"
	"strconv"
)

// creating an alias for fmt.Println
var pl = fmt.Println

func main() {
	// int, float64, bool, string, rune
	// Default type 0, 0.0, false, ""
	pl(reflect.TypeOf(25))
	pl(reflect.TypeOf(3.14))
	pl(reflect.TypeOf(true))
	pl(reflect.TypeOf("Hello"))
	pl(reflect.TypeOf('😎'))

	/** Casting **/
	cV1 := 1.5

	// converting float64 to int
	cV2 := int(cV1)
	pl(cV2)

	// converting string to int
	cV3 := "50000000"
	cV4, err := strconv.Atoi(cV3)
	pl(cV4, err, reflect.TypeOf(cV4))

	// converting int to string
	cV5 := strconv.Itoa(cV4)
	pl(cV5)

	// converting string to float
	cV6 := "3.14"
	if cV7, err := strconv.ParseFloat(cV6, 64); err == nil {
		pl(cV7, reflect.TypeOf(cV7))
	}

	// converting from float to string 
	cV8 := fmt.Sprintf("%f", 3.14)
	pl(cV8, reflect.TypeOf(cV8))

}
