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
	nameMap["Fork"] = itemInformation{0, 4, 3}
	nameMap["Plates"] = itemInformation{0, 4, 3}
	nameMap["Cups"] = itemInformation{0, 5, 3}
	nameMap["Bread"] = itemInformation{1, 2, 2}
	nameMap["Cake"] = itemInformation{1, 3, 1}
	nameMap["Coke"] = itemInformation{2, 5, 2}
	nameMap["Sprite"] = itemInformation{2, 5, 2}
	category := []string{"Household", "Food", "Drink"}
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
			fmt.Println("Shopping List Contents:")
			for key, element := range nameMap {
				fmt.Printf("Category: %v - Item: %v Quantity: %v Unit Cost: %v\n", category[element.catType], key, element.quantity, element.cost)
			}
			pressToContinue()
		case 2:
			//generateReport()
			pressToContinue()
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
			pressToContinue()
		case 4:
			//modifyItem()
			pressToContinue()
		case 5:
			//deleteItem()
			pressToContinue()
		case 6:
			for key, element := range nameMap {
				fmt.Println(key, element)
				pressToContinue()
			}
		case 7:
			var userInput string
			fmt.Println("Add New Category Name")
			fmt.Println("What is the New Category Name to add?")
			_, _ = fmt.Scanln(&userInput)

			if userInput == "" {
				fmt.Println("No Input Found!")

			}
			for _, element := range nameMap {
				if userInput == category[element.catType] {
					fmt.Printf("Category: %v already exit at index %v !\n", userInput, element.catType)
					pressToContinue()
				} else {
					//category = append(category, userInput)
					//for i, cat := range category {
					//	if userInput == cat {
					//		fmt.Printf("New category: %v added at index %v\n", cat, i)
					//		pressToContinue()
					//	}

					//}
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
func pressToContinue() {
	fmt.Println("\nPress any key to continue....")
	_, _ = fmt.Scanln()
}
