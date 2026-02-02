package main

import "fmt"

func main() {
	// Array: Fixed size, Size is part of the type, Rarely used directly
	var a [3]int
	var b [4]int // different type

	// Arrays shine when:
	//  The size is truly fixed and small
	//  You want value semantics (copy on pass)
	//  You’re modeling something physical and bounded, like a 3D vector or a fixed-size hash block

	fmt.Println(a, b)

	// Slice: len vs cap
	// When you pass a slice to a function, you copy the header, not the data. Two slice headers can point to the same backing array. That’s why mutations can leak across boundaries like gossip in an open office.
	// len: how many elements you can safely access
	// cap: how many elements you can grow into before reallocation

	nums := []int{1, 2, 3}
	fmt.Println(len(nums)) // visible elements
	fmt.Println(cap(nums)) // total available space

}
