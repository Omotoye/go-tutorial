package main

import (
	"fmt"
)

// creating an alias for fmt.Println
var pl = fmt.Println

func sayHello() {
	pl("Hello")
}

func getSum(x int, y int) int {
	return x + y
}

func getQuotient(x, y float64) (ans float64, err error) {
	if y == 0 {
		return 0, fmt.Errorf("You cant divide by zero")
	} else {
		return x / y, nil
	}
}

// function that can take unspecified amount of values
func getSum2(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func getArraySum(arr []int) int {
	sum := 0
	for _, val := range arr {
		sum += val 
	}
	return sum 
}

// return tuples
func getTwo(x int) (int, int) {
	return x + 1, x + 2
}

// func parameters are pass by value in go
func changeVal(f3 int) int {
	f3 += 1
	return f3
}

func main() {
	// func funcName(parameters) returnType {BODY}
	sayHello()
	pl(getSum(5, 4))
	pl(getTwo(5))
	pl(getQuotient(5, 0))
	pl(getQuotient(5, 4))
	pl(getSum2(1, 2, 3, 4, 5))

	vArr := []int{1,2,3,4}

	// arrays are pass by value in go
	pl("Array Sum :", getArraySum(vArr))
	f3 := 5
	pl("f3 before func :", f3)
	changeVal(f3)
	pl("f3 after func :", f3)

}
