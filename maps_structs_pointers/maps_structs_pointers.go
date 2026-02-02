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
	delete(m, "apple")

}
