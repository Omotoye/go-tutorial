package main

import (
	"fmt"
)

// creating an alias for fmt.Println
var pl = fmt.Println

type MyConstraint interface {
	int | float64
}

// find more about constraint in the link below
// pkg.go.dev/golang.org/x/exp/constraints

func getSumGen[T MyConstraint](x T, y T) T {
	return x + y
}

func main() {
	pl("5 + 4 =", getSumGen(5, 4))
	pl("5.6 + 4.7 =", getSumGen(5.6, 4.7))
}
