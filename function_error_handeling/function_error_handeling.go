package main

import (
	"errors"
	"fmt"
	"os"
)

func add(a int, b int) int {
	return a + b
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Can not divide by zero")
	}
	return a / b, nil
}

func openFile(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("file : ", file, "err : ", err)
		return
	}
}
func main() {

	result := add(2, 3)
	fmt.Println(result)

	res, err := divide(2, 5)
	fmt.Println("res : ", res, "err : ", err)
	openFile("data.txt")
}
