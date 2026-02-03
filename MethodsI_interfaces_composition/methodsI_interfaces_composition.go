package main

import "fmt"

type User struct {
	Name string
}

func (u User) Greet() string {
	return "Hello, " + u.Name
}

// Value Received: The method is read-only, The struct is small and immutable
func (u User) Rename(name string) {
	u.Name = name
}

// Pointer received: to modify the struct, The struct is large, want consistent behavior
func (u *User) RenamePointer(name string) {
	u.Name = name
}

func main() {
	u := User{Name: "Swapno"}
	fmt.Println(u.Greet())

	u.Rename("Rokon")
	fmt.Println(u.Greet())

	u.RenamePointer("shek")
	fmt.Println(u.Greet())
}
