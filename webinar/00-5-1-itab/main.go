package main

import (
	"fmt"
	"unsafe"
)

type Sayer interface {
	Say() string
}

type Dog struct{ Name string }

func (d Dog) Say() string { return "Woof, I'm " + d.Name }

// Зеркало структуры iface из runtime
type iface struct {
	tab  uintptr // указатель на itab
	data uintptr // указатель на данные
}

func main() {
	var s Sayer = Dog{Name: "Rex"}

	// Берём сырые байты переменной интерфейса
	raw := (*iface)(unsafe.Pointer(&s))

	fmt.Printf("tab  ptr: 0x%x\n", raw.tab)
	fmt.Printf("data ptr: 0x%x\n", raw.data)
	fmt.Printf("Оба != 0, значит интерфейс не nil: %v\n", raw.tab != 0 && raw.data != 0)

	// А теперь nil-интерфейс
	var s2 Sayer
	raw2 := (*iface)(unsafe.Pointer(&s2))
	fmt.Printf("\nnil iface — tab: 0x%x, data: 0x%x\n", raw2.tab, raw2.data)
}
