package menu

import "fmt"

//main menu option 7
func newCategory() {
	var newCatName string
	fmt.Println("Add New Category Name")
	fmt.Println("What is the New Category Name to add?")
	_, _ = fmt.Scanln(&newCatName)
	// if input is nil print no input Found
	if newCatName == "" {

		fmt.Println("No Input Found!")
		printBreak()
		returnMainOrExit()

	}
}
