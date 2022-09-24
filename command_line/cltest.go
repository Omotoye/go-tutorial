package main

import (
	"fmt"
	"os"
	"strconv"
)


func main() {
	fmt.Println(os.Args[1:])
	args := os.Args[1:]
	var iArgs = []int{}
	for _, i := range args {
		val, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		iArgs = append(iArgs, val)
	}
}
