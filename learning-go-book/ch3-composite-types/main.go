package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	fmt.Println("Hello, World!")
	p := person{name: "Alice", age: 30}
	fmt.Println(p)
	fmt.Println(p.name)
	fmt.Println(p.age)
}
