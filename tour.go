package main

import (
	"fmt"
	"time"
)

func add(x, y int) int {
	return x + y
}

func addAndPrint(x, y int) string {
	return fmt.Sprintf("The sum of %d and %d is %d", x, y, add(x, y))
}

func main() {
	var x, y int = 42, 13
	fmt.Println("Welcome to the playground!")
	fmt.Println("The time is", time.Now())
	fmt.Println(addAndPrint(x, y))
}
