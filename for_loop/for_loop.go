package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

// creating an alias for fmt.Println
var pl = fmt.Println

func main() {
	// for initialization; condition; postStatement {BODY}
	for x := 1; x <= 5; x++ {
		pl(x)
	}
	pl()
	fX := 0
	for fX < 5 {
		pl(fX)
		fX++
	}

	// looping through an array 
	aNums := []int{1,2,3}
	pl()
	for _, num := range aNums {
		pl(num)
	}

	// creating a gueesing game
	seedSecs := time.Now().Unix()
	rand.Seed(seedSecs)
	randNum := rand.Intn(50) + 1

	for true {
		fmt.Print("Guess a number between 0 and 50 :")
		pl("Random Number is :", randNum)
		reader := bufio.NewReader(os.Stdin)
		guess, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		guess = strings.TrimSpace(guess)
		iGuess, err := strconv.Atoi(guess)
		if err != nil {
			log.Fatal(err)
		}
		if iGuess > randNum {
			pl("Pick a Lower Value")
		} else if iGuess < randNum {
			pl("Pick a Higher Value")
		} else {
			pl("You Guessed it")
			break
		}
	}
}
