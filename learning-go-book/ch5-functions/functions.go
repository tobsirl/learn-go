package main

import "fmt"

type MyFuncOpts struct {
	FirstName string
	LastName string
	Age int
}

func MyFunc(opts MyFuncOpts) error  {
	fmt.Println("Name:", opts.FirstName + " " + opts.LastName)
	fmt.Println("Age:", opts.Age)
	return nil
}

func main() {
	MyFunc(MyFuncOpts{
		FirstName: "John",
		LastName: "Doe",
		Age: 30,
	})

	MyFunc(MyFuncOpts{
		FirstName: "Jane",
		LastName: "Smith",
		Age: 25,
	})
}