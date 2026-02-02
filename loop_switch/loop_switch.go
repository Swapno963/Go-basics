package main

import "fmt"

func main() {
	x := 1

	// for (Go has only one loop)
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// While-style:
	for x < 10 {
		x++
	}

	// Infinite loop:
	for {
		break
	}

	// switch
	day := "Sun"

	switch day {
	case "Mon":
		fmt.Println("Work")
	case "Sun":
		fmt.Println("Rest")
	default:
		fmt.Println("Unknown")
	}

	fmt.Println(x)

}
