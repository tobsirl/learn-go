package main

func main() {
	// learn about pointers in Go
	var x int32 = 10
	var y bool = true

	pointerX := &x
	pointerY := &y
	var pointerZ *string

	println("Value of x:", x)
	println("Address of x:", pointerX)
	println("Value at address of x:", *pointerX)

	println("Value of y:", y)
	println("Address of y:", pointerY)
	println("Value at address of y:", *pointerY)

	println("Value of pointerZ (should be nil):", pointerZ)

	
}