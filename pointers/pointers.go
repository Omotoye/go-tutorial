package main

import (
	"fmt"
)

// creating an alias for fmt.Println
var pl = fmt.Println

// func parameters are pass by value in go
func changeVal(f3 int) int {
	f3 += 1
	return f3
}

func changeVal2(myPtr *int) {
	*myPtr = 12
}

func dblArrVals(arr *[4]int) {
	for x := 0; x < 4; x++ {
		arr[x] *= 2
	}
}

func getAverage(nums ...float64) float64 {
	var sum float64 = 0.0
	var numSize float64 = float64(len(nums))
	for _, val := range nums {
		sum += val
	}
	return (sum / numSize)
}

func main() {
	f4 := 10
	var f4Ptr *int = &f4
	pl("f4 Address :", f4Ptr)
	pl("f4 Value :", *f4Ptr)
	*f4Ptr = 11
	pl("f4 Value :", *f4Ptr)

	f3 := 5
	pl("f3 before func :", f3)
	changeVal(f3)
	pl("f3 after func :", f3)

	pl("Using pointers")
	pl("f3 before func :", f3)
	changeVal2(&f3)
	pl("f3 after func :", f3)

	// passing array pointers
	pArr := [4]int{1, 2, 3, 4}
	dblArrVals(&pArr)
	pl(pArr)

	iSlice := []float64{11, 13, 17}
	fmt.Printf("Average : %.3f\n", getAverage(iSlice...))
}
