package main

import "fmt"

func main() {
	// Array: Fixed size, Size is part of the type, Rarely used directly
	var a [3]int
	var b [4]int // different type
	fmt.Println(a, b)
}
