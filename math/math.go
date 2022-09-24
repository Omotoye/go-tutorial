package main

import (
	"fmt"
	"math/rand"
	"time"
)

import "math"

// creating an alias for fmt.Println
var pl = fmt.Println

func main() {
	pl("5 + 4 =", 5+4)
	pl("5 - 4 =", 5-4)
	pl("5 * 4 =", 5*4)
	pl("5 / 4 =", 5/4)
	pl("5 % 4 =", 5%4)

	// incrementing a variable
	mInt := 1
	mInt = mInt + 1
	mInt += 1
	mInt++
	pl(mInt)

	// Creating random value
	seedSecs := time.Now().Unix()
	rand.Seed(seedSecs)

	// random value from 0 - 50
	randNum := rand.Intn(50) + 1
	pl("Random :", randNum)

	// Using some of the math functions
	pl("Abs(-10) =", math.Abs(-10))
	pl("Pow(4, 2) =", math.Pow(4, 2))
	pl("Sqrt(16) =", math.Sqrt(16))
	pl("Cbrt(8) =", math.Cbrt(8))
	pl("Ceil(4.4) =", math.Ceil(4.4))
	pl("Floor(4.4) =", math.Floor(4.4))
	pl("Round(4.4) =", math.Round(4.4))
	pl("Log2(8) =", math.Log2(8))
	pl("Log10(100) =", math.Log10(100))
	pl("Log(7.389) =", math.Log(math.Exp(2)))
	pl("Max(5, 4) =", math.Max(5, 4))
	pl("Min(5, 4) =", math.Min(5, 4))

	// convet 90 degrees to radians
	r90 := 90 * math.Pi / 180
	d90 := r90 * (180 / math.Pi)
	fmt.Printf("%.2f radians = %.4f degree\n", r90, d90)
	pl("Sin(90) =", math.Sin(r90))
}
