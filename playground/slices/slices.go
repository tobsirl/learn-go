package main

type User struct {
	Name string
	Age  int
}

func main() {
	var users []User
	users = append(users, User{Name: "Alice", Age: 30})
	users = append(users, User{Name: "Bob", Age: 25})

	for _, user := range users {
		println("Name:", user.Name, "Age:", user.Age)
	}
}