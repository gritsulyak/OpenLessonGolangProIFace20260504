package main

import "fmt"

func main() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	p, ko := i.(fmt.Stringer)
	fmt.Println(p, ko)

	r, ok := i.(fmt.Stringer)
	fmt.Println(r, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)
}
