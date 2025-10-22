package main

import (
	"fmt"
)

// Map of valid users
var users = map[string]string{
	"adams":  "424hh1",
	"davani": "3773h5",
	"godwin": "373hsn",
	"john":   "jhdfhj",
}

func main() {
	var username, userpassword string

	fmt.Println("Enter your username: ")
	fmt.Scan(&username)

	fmt.Println("Enter your userpassword: ")
	fmt.Scan(&userpassword)
	// check if username exist and userpassword matches
	if storedUserpassword, exists := users[username]; exists && storedUserpassword == userpassword {
		fmt.Printf("\n Welcome, %v! login successful\n", username)
	} else {
		fmt.Println("Y\n Welcome,%v your Userpassword, %v! is incorrect\n")
	}
}
