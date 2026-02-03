package main

import "fmt"

type User struct {
	ID    int
	Name  string
	Email string
}

type BigStruct struct {
	ID    int
	Name  string
	Email string
}

// 1. To Modify Data
func update(u *User) {
	u.Name = "Bob"
}

// 2. To Avoid Copying Large Structs
func process(u *BigStruct) {}
func main() {

	// Reference type, Fast lookup (hash table), Unordered, Zero value is nil
	m := map[string]int{
		"apple":  3,
		"banana": 5,
	}

	v, ok := m["apple"]
	if ok {
		fmt.Println(v)
	}

	// Delete
	// delete(m, "apple")

	u := User{
		ID:   1,
		Name: "Alice",
	}
	fmt.Println(u)

	// When converting this struct to JSON, call this field id, not ID.
	// type User struct {
	// 	ID   int    `json:"id"`
	// 	Name string `json:"name"`
	// }
	x := 10
	p := &x
	fmt.Println(p)  // address
	fmt.Println(*p) // Value

}
