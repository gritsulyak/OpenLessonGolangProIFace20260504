package main

import (
	"fmt"
)

type Hound interface {
	destroy()
	bark(int)
}

type Retriever interface {
	Hound
	bark(int) // duplicate method - int
}

func main() {
	fmt.Println("hello")
}
