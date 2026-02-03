package main

import "fmt"

type User struct {
	Name string
}

func (u User) Greet() string {
	return "Hello, " + u.Name
}

// Value Received
func (u User) Rename(name string) {
	u.Name = name
}

func main() {
	u := User{Name: "Swapno"}
	fmt.Println(u.Greet())

	u.Rename("Rokon")
	fmt.Println(u.Greet())
}
