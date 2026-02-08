package main

import "fmt"

func main() {
	var firstName string
	var lastName string
	var choice int
	var email string
	var phone string
	var password string
	var age int

	fmt.Println("=== Facebook Account Creation ===")

	fmt.Print("Enter your First Name: ")
	fmt.Scan(&firstName)

	fmt.Print("Enter your Last Name: ")
	fmt.Scan(&lastName)

	fmt.Println("\nChoose signup method:")
	fmt.Println("1. Email")
	fmt.Println("2. Phone Number")
	fmt.Println("Enter choice (1 or 2): ")
	fmt.Scan(&choice)
	if choice == 1 {
		fmt.Print("Enter your Email: ")
		fmt.Scan(&email)
	} else if choice == 2 {
		fmt.Print("Enter your phone number: ")
		fmt.Scan(&phone)
	} else {
		fmt.Println("Invalid choice. Look carefully and choose.")
		return
	}

	fmt.Print("Creat a Password: ")
	fmt.Scan(&password)

	fmt.Print("Enter your Age: ")
	fmt.Scan(&age)

	fmt.Println("\n--- Account Created Successfully ---")
	fmt.Printf("Name: %s %s\n", firstName, lastName)
	fmt.Printf("Email: %s\n", email)
	fmt.Printf("Age: %d\n", age)

	fmt.Println("\nWelcome to Facebook 🎉")
}
