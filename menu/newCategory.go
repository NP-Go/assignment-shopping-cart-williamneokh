package menu

import "fmt"

type categories []string

//main menu option 7
func newCategory() {
	var userInput string
	fmt.Println("Add New Category Name")
	fmt.Println("What is the New Category Name to add?")
	_, _ = fmt.Scanln(&userInput)
	category := categories{}
	// if input is nil print no input Found
	if userInput == "" {

		fmt.Println("No Input Found!")
		printBreak()
		returnMainOrExit()

	} else {
		for _, catName := range category {
			if catName == userInput {
				fmt.Printf("This %v has been taken!", userInput)
				printBreak()
				returnMainOrExit()
			}

		}
		category = append(category, userInput)
		fmt.Println(userInput, "has been added!")
		printBreak()
		returnMainOrExit()
	}
}

func createCat(s string) string {
	return s
}
