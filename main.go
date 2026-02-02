package main

import "fmt"

func main() {
	age := 10

	fmt.Println(age)

	if age >= 18 {
		fmt.Println("Adult")
	} else {
		fmt.Println("Minor")
	}

	if x := 6; x > 5 {
		fmt.Println("Big")
	}

}
