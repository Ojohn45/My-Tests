package main

import "fmt"

func main() {
	var name string
	var age int

	fmt.Println("Enter your name: ")
	fmt.Scan(&name)

	fmt.Println("Enter your age: ")
	fmt.Scan(&age)

	fmt.Printf("Hello, %s! You Are, %d Years old, Thanks.", name, age)
}
