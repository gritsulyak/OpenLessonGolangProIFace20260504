package main

import (
	"fmt"
)

type S struct{} // Пустая структура — не содержит полей, занимает 0 байт

func main() {
	var s1 S
	fmt.Printf("%#v, %T\n", s1, s1)

	var s2 any = 123
	fmt.Printf("%#v, %T\n", s2, s2)
	s2 = S{}
	fmt.Printf("%#v, %T\n", s2, s2)
}
