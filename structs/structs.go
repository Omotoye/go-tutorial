package main

import (
	"fmt"
)

// creating an alias for fmt.Println
var pl = fmt.Println

type customer struct {
	name    string
	address string
	bal     float64
}

type rectangle struct {
	length, height float64
}

func getCustInfo(c customer) {
	fmt.Printf("%s owes us %.2f\n", c.name, c.bal)
}

func newCustAdd(c *customer, address string) {
	c.address = address
}

// this is a function that is part of the struct rectangle
func (r rectangle) Area() float64 {
	return r.length * r.length
}

/*
 Go does not support inheritance, but it supports
 composition how this is done if by embeding a
 struct inside another struct
*/

type contact struct {
	fName string
	lName string
	phone string
}

type business struct {
	name string
	address string
	contact  
}

func (b business) info (){
	fmt.Printf("Contact at %s is %s %s", b.name, b.contact.fName, b.contact.lName)
}

func main() {
	var tS customer
	tS.name = "Tom Smith"
	tS.address = "5 main st"
	tS.bal = 234.56
	getCustInfo(tS)
	newCustAdd(&tS, "123 South st")
	pl("Address :", tS.address)
	sS := customer{"Sally Smith", "123 Main", 0.0}
	pl("Name :", sS.name)

	rect1 := rectangle{10.0, 15.0}
	pl("Rect Area :", rect1.Area())
	con1 := contact {
		"James", 
		"Wang", 
		"555-1212",
	}
	bus1 := business {
		"ABC Plumbing", 
		"234 North St", 
		con1,
	}
	bus1.info()
}
