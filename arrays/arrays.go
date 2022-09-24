package main

import (
	"fmt"
)

// creating an alias for fmt.Println
var pl = fmt.Println

func main() {
	var arr1 [5]int
	arr1[0] = 1
	arr2 := [5]int{1, 2, 3, 4, 5}
	pl("Index 0 :", arr2[0])
	pl("Arr Length :", len(arr2))

	// iterating through arrays
	for i := 0; i < len(arr2); i++ {
		pl(arr2[i])
	}

	// range based for loop
	for i, v := range arr2 {
		fmt.Printf("%d : %d\n", i, v)
	}

	// creating multidimensional array
	arr3 := [2][2]int{
		{1, 2},
		{3, 4},
	}

	for i := 0; i < len(arr3); i++ {
		for j := 0; j < len(arr3[i]); j++ {
			pl(arr3[i][j])
		}
	}

	// slicing a string 
	aStr1 := "abcde"
	rArr := []rune(aStr1)
	for _, v  := range rArr {
		fmt.Printf("Rune Array : %d\n", v)
	}

	// converting a byte array into a string
	byteArr := []byte{'a', 'b', 'c'}
	bStr := string(byteArr[:])
	pl("I'm a string :", bStr)
}
