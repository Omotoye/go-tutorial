package main

import (
	"fmt"
)

// creating an alias for fmt.Println
var pl = fmt.Println

func main() {
	// var name []datatype
	sl1 := make([]string, 6)

	sl1[0] = "Society"
	sl1[1] = "of"
	sl1[2] = "the"
	sl1[3] = "Simulated"
	sl1[4] = "Universe"
	pl("Slice size :", len(sl1))
	for i := 0; i < len(sl1); i++ {
		pl(sl1[i])
	}

	for _, x := range sl1 {
		pl(x)
	}

	sArr := [5]int{1, 2, 3, 4, 5}
	sl3 := sArr[0:2]
	pl("1st 3 :", sArr[:3])
	pl("Last 3 :", sArr[2:])
	sArr[0] = 10
	pl("sl3 :", sl3)

	sl3 = append(sl3, 12)
	pl("sl3 :", sl3)
	pl("sA :", sl3)

	sl4 := make([]string, 6)
	pl("sl4 :", sl4)
	pl("sl4[0] :", sl4[0])

}
