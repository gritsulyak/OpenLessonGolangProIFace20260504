package main

import "fmt"

type Meower interface {
	meow() bool
}

type Cat struct{}

func (c Cat) meow() bool {
	fmt.Println("meowmrr")
	return true
}

func main() {
	var pus Meower = Cat{}
	pus.meow()
	var pus2 = pus
	fmt.Println(&pus)
	fmt.Println(&pus2)
}
