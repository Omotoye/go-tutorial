package main

import (
	"fmt"
)

// creating an alias for fmt.Println
var pl = fmt.Println

func main() {
	// Conditonal Operators : > < >= <= == !=
	// Logical Operators : && || !
	iAge := 8
	if (iAge >= 1) && (iAge <= 18) {
		pl("Important Birthday")
	} else if (iAge == 21) || (iAge == 50) {
		pl("Second Level Important Birthday")
	} else if iAge >= 65 {
		pl("Third Level Important Birthday")
	} else {
		pl("Not an Important Birthday")
	}
	pl("!true =", !true)
}
