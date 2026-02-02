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

}
