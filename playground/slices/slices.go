package main

type User struct {
	Name string
	Age  int
}

func main() {
	users := []User{{Name: "John", Age: 20}, {Name: "Jane", Age: 22}}

	for _, user := range users {
		println("Name:", user.Name, "Age:", user.Age)
	}
}