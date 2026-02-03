package main

import "fmt"

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

	type User struct {
		ID    int
		Name  string
		Email string
	}

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
