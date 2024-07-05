package main

import "fmt"

func main() {
	var firstName string = "john"

	var lastName string
	lastName = "wick"

	// karakter %s digunakan untuk menunjukkan tipe data string
	fmt.Printf("halo %s %s!\n", firstName, lastName)
}