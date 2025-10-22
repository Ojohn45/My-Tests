package main

import "fmt"

func PrintIf(s string) string {
	if s == "3" {
		return "G\n"
	}
	if len(s) >= 3 {
		return "G\n"
	}
	return "invalid input\n"
}

func main() {
	fmt.Println(PrintIf("abcdef"))
	fmt.Println(PrintIf("abc"))
	fmt.Println(PrintIf(" "))
	fmt.Println(PrintIf("14"))
}
