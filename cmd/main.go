package main

import (
	"fmt"
)

type itemInformation struct {
	catType  int
	quantity int
	cost     int
}

func main() {
	nameMap := make(map[string]itemInformation)
	var category []string
	for {
		var choice int
		fmt.Println("Shopping List Application")
		printBreak()
		fmt.Println("1. View Entire Shopping list")
		fmt.Println("2. Generate Shopping List Report")
		fmt.Println("3. Add Items")
		fmt.Println("4. Modify Items")
		fmt.Println("5. Delete Item")
		fmt.Println("6. Print Current Data")
		fmt.Println("7. Add New Category Name")
		fmt.Println("Select your choice:")
		_, _ = fmt.Scanln(&choice)

		switch choice {
		case 1:
			//viewEntireShoppingList()
		case 2:
			//generateReport()
		case 3:
			//addItem()
			var itemName string
			var catName string
			var quantityNum int
			var itemCost int

			fmt.Println("What is the name of your item?")
			_, _ = fmt.Scanln(&itemName)
			fmt.Println("What category does it belong to?")
			_, _ = fmt.Scanln(&catName)
			fmt.Println("How many units are there?")
			_, _ = fmt.Scanln(&quantityNum)
			fmt.Println("How much does it cost per unit?")
			_, _ = fmt.Scanln(&itemCost)

			for i, cat := range category {
				if catName == cat {
					nameMap[itemName] = itemInformation{i, quantityNum, itemCost}
				} else {
					fmt.Println("Please first register the Category, before adding new item.")
				}
			}

			for key, element := range nameMap {
				fmt.Println(key, element)
			}
		case 4:
			//modifyItem()
		case 5:
			//deleteItem()
		case 6:
			for key, element := range nameMap {
				fmt.Println(key, element)
			}
		case 7:
			var userInput string
			fmt.Println("Add New Category Name")
			fmt.Println("What is the New Category Name to add?")
			_, _ = fmt.Scanln(&userInput)

			if userInput == "" {
				fmt.Println("No Input Found!")

			} else {

				category = append(category, userInput)
				for i, cat := range category {
					if userInput == cat {
						fmt.Printf("New category: %v added at index %v\n", cat, i)
					}

				}
			}

		default:
			//MainMenu()

		}
	}
}
func printBreak() {
	fmt.Println("=========================")
}
