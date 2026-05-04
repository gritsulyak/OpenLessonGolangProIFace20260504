package main

import "fmt"

func PrintAll(vals []interface{}) {
	for _, val := range vals {
		fmt.Println(val)
	}
}

func main() {
	names := []string{"stanley", "david", "oscar"}

	//  Создаем срез нужного типа
	vals := make([]interface{}, len(names))
	for i, v := range names {
		vals[i] = v
	}

	PrintAll(vals)
	PrintAllT(names)
}

func PrintAllT[T any](vals []T) {
	for _, val := range vals {
		fmt.Println(val)
	}
}
