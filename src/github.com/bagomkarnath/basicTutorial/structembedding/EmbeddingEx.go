package main

import "fmt"

type Animal struct {
	name   string
	origin string
}

type Bird struct {
	Animal
	canFly string
	speed  string
}

func main() {
	b := Bird{}
	b.name = "humming"
	b.origin = "sa"
	b.speed = "23kmph"
	b.canFly = "yes"
	fmt.Println(b)

	c := Bird{
		Animal: Animal{name: "Emu", origin: "Aus"},
		speed:  "45kmph",
		canFly: "no",
	}
	fmt.Println(c)
}
