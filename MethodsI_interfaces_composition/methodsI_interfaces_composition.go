package main

import "fmt"

type User struct {
	Name string
}

func (u User) Greet() string {
	return "Hello, " + u.Name
}

func main() {
	u := User{Name: "Swapno"}
	fmt.Println(u.Greet())
}
