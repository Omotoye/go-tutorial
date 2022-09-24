package main

import (
	"fmt"
	"bufio"
	"os"
	"log"
)

// creating an alias for fmt.Println
var pl = fmt.Println

func main() {
	pl("Hello Go")
	pl("What is your name?")

	// Initialize object for reading from the stdin
	reader := bufio.NewReader(os.Stdin)

	// use the bufio object to read from stdin unti "\n" or enter
	// err is for taking in error message that might occur and handle 
	// it. you can use _ to ignore the err
	name, err := reader.ReadString('\n')

	// error handling for stdin
	if err == nil {
		pl("Hello", name)
	} else {
		log.Fatal(err)
	}
}
