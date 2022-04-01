package main

import (
	"fmt"
)

type itemInformation struct {
	catType  int
	quantity int
	cost     float64
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
			fmt.Println("Generate Report")
			fmt.Println("1. Total cost of each category")
			fmt.Println("2. List of item by category")
			fmt.Println("3. Main Menu")
			fmt.Println()
			fmt.Println("Choose your Report:")

			_, _ = fmt.Scanln(&choice)

			switch choice {
			case 1:
				// find total cost of each category
				var totalSum float64
				totalSumMap := make(map[string]float64)
				for i, catName := range category {
					for _, element := range nameMap {
						if i == element.catType {
							sum := float64(element.quantity) * element.cost
							totalSum = totalSum + sum

							totalSumMap[catName] = totalSum

							//fmt.Printf("Category: %v - Item: %v Quantity: %v Unit Cost: %v\n", cat, key, element.quantity, element.cost)
						}
					}
				}
				for key, total := range totalSumMap {
					fmt.Printf("%v cost: %v\n", key, total)
				}
				pressToContinue()
			case 2:
				// list by category
				//var totalCost float64

				for i, cat := range category {
					for key, element := range nameMap {
						if i == element.catType {

							fmt.Printf("Category: %v - Item: %v Quantity: %v Unit Cost: %v\n", cat, key, element.quantity, element.cost)
						}

					}
				}
				pressToContinue()
			case 3:
				continue
			}

		case 3:
			//addItem()

			var itemName string
			var catName string
			var quantityNum int
			var itemCost float64
			var matchAny = false

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
					matchAny = true
				}

			}
			if matchAny == false {
				fmt.Println("Please register new category before adding new item.")
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
			}
			pressToContinue()
		case 7:
			var matchAny = false
			var userInput string
			var indexPlace int
			fmt.Println("Add New Category Name")
			fmt.Println("What is the New Category Name to add?")
			_, _ = fmt.Scanln(&userInput)

			if userInput == "" {
				fmt.Println("No Input Found!")

			}
			for _, element := range nameMap {
				if userInput == category[element.catType] {
					matchAny = true
					indexPlace = element.catType
				}
			}
			if matchAny == true {
				fmt.Printf("Category: %v already exit at index %v !\n", userInput, indexPlace)
				pressToContinue()
			} else {

				category = append(category, userInput)
				fmt.Printf("New category: %v added at index %v\n", userInput, len(category)-1)
				pressToContinue()
			}

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
