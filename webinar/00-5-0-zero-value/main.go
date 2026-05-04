package main

import (
	"log"
	"os"
)

func ReadFile(fname string) error {
	var err *os.PathError // nil

	if err == nil {
		log.Println("err is nil")
	}

	// Do some work without err...
	return err
}

func main() {
	if err := ReadFile(""); err != nil {
		log.Printf("ERR: (%T, %v)", err, err)
	} else {
		log.Println("OK")
	}
}
