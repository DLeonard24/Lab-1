package main

import "fmt"

func main() {
	R := 1.2
	D := 0
	I := 7
	Pop := 39740
	Final := 0

	fmt.Println("Enter Day: ")
	fmt.Scanln(&D)

	for i := 0; i < D; i++ {
		Final += R * I
		println(Pop)
		return Final
	}
}
