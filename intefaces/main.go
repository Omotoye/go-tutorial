package main

import (
	"fmt"
)

// creating an alias for fmt.Println
var pl = fmt.Println

type Animal interface {
	AngrySound()
	HappySound()
}

type Cat string

func (c Cat) Attack() {
	pl("cat Attacks its Prey")
}

func (c Cat) Name() string {
	return string(c)
}

func (c Cat) AngrySound() {
	pl("Cat says Hisssss")
}

func (c Cat) HappySound() {
	pl("Cat says Purrrr")
}

func main() {
	var kitty Animal
	kitty = Cat("Kitty")
	kitty.AngrySound()
	var kitty2 Cat = kitty.(Cat)
	kitty2.Attack()
	pl("Cats Name :", kitty2.Name())
}
