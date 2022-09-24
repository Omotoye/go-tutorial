package main

import (
	"fmt"
	"strings"
)

// creating an alias for fmt.Println
var pl = fmt.Println

func main() {
	sV1 := "A word"
	// strings are arrays of bytes

	// replacing a charater in a string
	replacer := strings.NewReplacer("A", "Another")
	sV2 := replacer.Replace(sV1)
	pl(sV2)

	// getting the length of string / anything
	pl("Length :", len(sV2))

	// checking if a string contains a substring
	pl("Contains Another :", strings.Contains(sV2, "Another"))

	// getting the index of a string 
	pl("o index :", strings.Index(sV2, "o"))

	// replace matches
	pl("Replace :", strings.Replace(sV2, "o", "0", -1)) // -1 mean replace all matches
	// a specific number of matches to be replaced can also be set

	// removing white spaces
	sV3 := "\nSome Words\n"
	sV3 = strings.TrimSpace(sV3)

	// slit string with a dilimeter
	pl("Split :", strings.Split("a-b-c-d", "-"))

	// convert to lowercase
	pl("Lower :", strings.ToLower(sV2))

	// convert to Uppercase
	pl("Upper :", strings.ToUpper(sV2))

	// check the beginning of a string
	pl("Prefix :", strings.HasPrefix("tacocat", "taco"))

	// check the end of a string 
	pl("Suffix :", strings.HasSuffix("tacocat", "cat"))

}
