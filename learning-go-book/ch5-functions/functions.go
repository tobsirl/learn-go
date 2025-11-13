package main

import (
	"encoding/json"
	"fmt"
)

type MyFuncOpts struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
}

func MyFunc(opts MyFuncOpts) error {
	fmt.Println("Name:", opts.FirstName+" "+opts.LastName)
	fmt.Println("Age:", opts.Age)
	return nil
}

// ToJSON converts the struct to JSON string
func (m MyFuncOpts) ToJSON() (string, error) {
	jsonBytes, err := json.Marshal(m)
	if err != nil {
		return "", err
	}
	return string(jsonBytes), nil
}

// FromJSON creates a MyFuncOpts from JSON string
func FromJSON(jsonStr string) (MyFuncOpts, error) {
	var opts MyFuncOpts
	err := json.Unmarshal([]byte(jsonStr), &opts)
	return opts, err
}

func main() {
	// Example 1: Regular struct usage
	// person1 := MyFuncOpts{
	// 	FirstName: "John",
	// 	LastName: "Doe",
	// 	Age: 30,
	// }
	// MyFunc(person1)

	// // Example 2: Convert struct to JSON
	// jsonStr, err := person1.ToJSON()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("JSON representation:", jsonStr)

	// // Example 3: Create struct from JSON string
	// jsonData := `{"first_name":"Jane","last_name":"Smith","age":25}`
	// person2, err := FromJSON(jsonData)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("Created from JSON:")
	// MyFunc(person2)

	// // Example 4: Pretty print JSON
	// prettyJSON, err := json.MarshalIndent(person1, "", "  ")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("Pretty JSON:")
	// fmt.Println(string(prettyJSON))

	// Example 5: Anonymous function
	// f := func(j int) {
	// 	fmt.Println("printing", j, "from inside of an anonymous function")
	// }
	// for i := 0; i < 5; i++ {
	// 	f(i)
	// }

	// Inline Anonymous function
	for i := 0; i < 5; i++ {
		func(j int) {
			fmt.Println("printing", j, "from inside of an inline anonymous function")
		}(i)
	}
	func(j int) {
		fmt.Println("printing", j, "from inside of an inline anonymous function")
	}(11)
}
