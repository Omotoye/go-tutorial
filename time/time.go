package main

import (
	"fmt"
	"time"
)

// creating an alias for fmt.Println
var pl = fmt.Println

func main() {
	now := time.Now()
	pl(now.Year(), now.Month(), now.Day())
	pl(now.Hour(), now.Minute(), now.Second())
}
