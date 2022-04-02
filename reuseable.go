package main

import "fmt"

func PrintBreak() {
	fmt.Println("=========================")
}
func PressToContinue() {
	fmt.Println("\nPress any key to continue....")
	_, _ = fmt.Scanln()
}
