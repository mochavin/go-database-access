//go:build playground
// +build playground

package main

import (
	"fmt"
)

func example() {
	fmt.Println("Start")

	defer fmt.Println("Deferred at the beginning")
	defer fmt.Println("Deferred in the middle")
	defer fmt.Println("Deferred in the last")

	fmt.Println("Middle")
	fmt.Println("Middle")
	fmt.Println("Middle")
	fmt.Println("Middle")

	fmt.Println("End")
}

func main() {
	example()
}
