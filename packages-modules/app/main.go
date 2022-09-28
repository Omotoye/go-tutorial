package main

import "fmt"
import	stuff "example/project/mypackage"


func main() {
	fmt.Println("Hello", stuff.Name)
	intArr := []int{2,3,5,7,11}
	strArr := stuff.IntArrToStrArr(intArr)
	fmt.Println(strArr)
	fmt.Println(reflect.TypeOf(strArr))

	date := stuff.Date{}
	err := date.SetMonth(12)
	if err != nil {
		log.Fatal(err)
	}
	
	err = date.SetDay(21)
	if err != nil {
		log.Fatal(err)
	}
	
	err = date.SetYear(12)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("1st Day : %d/%d/%d\n", date.Month(), date.Day(), day.Year)
}