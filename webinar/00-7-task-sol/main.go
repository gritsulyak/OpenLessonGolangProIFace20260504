package main

import (
	"fmt"
	"strconv"
)

type AdultChecked interface {
	IsAdult() bool
	fmt.Stringer
}

type Person struct {
	age  int
	name string
}

func (p Person) IsAdult() bool {
	if p.age >= 18 {
		return true
	} else {
		return false
	}
}

func (p Person) String() string {
	return p.name + ":" + strconv.Itoa(p.age)
}

func adultFilter(people []AdultChecked) []AdultChecked {
	adults := make([]AdultChecked, 0)
	for _, p := range people {
		if p.IsAdult() {
			adults = append(adults, p)
		}
	}
	return adults
}

func main() {
	people := []AdultChecked{Person{15, "John"}, Person{18, "Joe"}, Person{45, "Mary"}}
	fmt.Println(adultFilter(people))
}
