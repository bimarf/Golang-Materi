package main

import "fmt"

func main() {
	var firstName string = "john"

	var lastName string
	lastName = "wick"

	// printf tidak menghasilkan baris baru
	fmt.Printf("halo john wick!\n")
	fmt.Printf("halo %s %s!\n", firstName, lastName)
	fmt.Println("halo", firstName, lastName+"!")
}