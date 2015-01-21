package main

import "fmt"

func main() {
	//prints with newline
	fmt.Println("Hello World!")
	fmt.Println("Hello　世界!")

	//prints without newline
	fmt.Print("Hello World!\n")

	something := "World"
	//printf
	fmt.Printf("%s\n", something)
}
