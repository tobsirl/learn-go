package main

import "fmt"

func main() {
	// learn about pointers in Go
	// var x int32 = 10
	// var y bool = true

	// pointerX := &x
	// pointerY := &y
	// var pointerZ *string

	// println("Value of x:", x)
	// println("Address of x:", pointerX)
	// println("Value at address of x:", *pointerX)

	// println("Value of y:", y)
	// println("Address of y:", pointerY)
	// println("Value at address of y:", *pointerY)

	// println("Value of pointerZ (should be nil):", pointerZ)

	// x := "hello"
	// pointerX := &x
	// println("Before change, x:", x)
	// *pointerX = "world"
	// println("After change, x:", x)

	// dereferencing
	// x := 10
	// pointerToX := &x
	// fmt.Println(pointerToX)
	// fmt.Println(*pointerToX)

	// z := 20 + *pointerToX
	// fmt.Println(z)
	name := "Alice"
	changeName(&name)
	fmt.Println("Memory address of name:", &name)

	m := &name
	fmt.Println("Value of m (address of name):", m)
	fmt.Println("Value at address m (value of name):", *m)
}

func changeName(name *string) {
	*name = "Bob"
}
