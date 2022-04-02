package main

import (
	"fmt"
)

//ViewShopList perform option 1: View Entire Shopping List
func ViewShopList() {
	fmt.Println("Shopping List Contents:")
	PrintBreak()
	for key, element := range nameMap {
		fmt.Printf("Category: %v - Item: %v Quantity: %v Unit Cost: %v\n", category[element.catType], key, element.quantity, element.cost)

	}
	PressToContinue()
}

//GenerateShoppingListReport perform Option 2: Generate Shopping List Report
func GenerateShoppingListReport() {
	var choice int
	fmt.Println("Generate Report")
	PrintBreak()
	fmt.Println("1. Total cost of each category")
	fmt.Println("2. List of item by category")
	fmt.Println("3. Main Menu")
	fmt.Println()
	fmt.Println("Choose your Report:")

	_, _ = fmt.Scanln(&choice)

	switch choice {
	case 1:
		num := make(map[string]float64)
		fmt.Println("Total cost by category")
		PrintBreak()
		for i, catName := range category {
			for _, element := range nameMap {
				if i == element.catType {
					sum := float64(element.quantity) * element.cost
					num[catName] = num[catName] + sum
				}
			}
		}
		for key, total := range num {
			fmt.Printf("%v cost: %v\n", key, total)
		}
		PressToContinue()
	case 2:
		fmt.Println("List by category")
		PrintBreak()
		for i, cat := range category {
			for key, element := range nameMap {
				if i == element.catType {

					fmt.Printf("Category: %v - Item: %v Quantity: %v Unit Cost: %v\n", cat, key, element.quantity, element.cost)
				}

			}
		}
		PressToContinue()
	case 3:
		break
	}
}

// AddItem perform option 3: Add Items
func AddItem() {
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
		fmt.Println("Category name mis-spell or category's name that you have enter not found. " +
			"\nPlease register the new category before adding new item. " +
			"\nSystem won't auto add new category during add new item processes " +
			"\nin case of user mis-spell and create multiple similar categories by accident.\n" +
			"(e.g sport and sports categories)")

	}

	PressToContinue()
}

//ModifyItems perform option 4: Modify Items
func ModifyItems() {
	var userInput string
	var newName string
	var newCategory string
	var newQuantity int
	var newCost float64
	var isChangeName = false
	var isChangeCategory = false
	var isChangeQuantity = false
	var catNum int
	fmt.Println("Modify Item")
	PrintBreak()
	fmt.Println("Which item would you wish to modify?")
	_, _ = fmt.Scanln(&userInput)
	// search map base on userInput as key

	for key, element := range nameMap {
		if userInput == key {
			fmt.Printf("Current item name is %v - Category is %v - Quantity is %v - Unit Cost %v\n", key, category[element.catType], element.quantity, element.cost)
			fmt.Println("Enter new name. Enter for no change")
			_, _ = fmt.Scanln(&newName)
			fmt.Println("Enter new Category. Enter for no change")
			_, _ = fmt.Scanln(&newCategory)
			fmt.Println("Enter new Quantity. Enter for no change")
			_, _ = fmt.Scanln(&newQuantity)
			fmt.Println("Enter new Unit cost. Enter for no change")
			_, _ = fmt.Scanln(&newCost)
			if newName == "" {
				fmt.Println("No change to item name made")
			} else {
				nameMap[newName] = nameMap[key]
				delete(nameMap, key)
				isChangeName = true
			}
			if newCategory == "" {
				fmt.Println("No change to category made")
			} else {

				//key, category[element.catType], element.quantity, element.cost
				//nameMap["Sprite"] = itemInformation{2, 5, 2}
				for i, value := range category {
					if newCategory == value {
						catNum = i
					}
				}
				if isChangeName == true {
					nameMap[newName] = itemInformation{catNum, element.quantity, element.cost}
					isChangeCategory = true
				} else {
					nameMap[key] = itemInformation{catNum, element.quantity, element.cost}
					isChangeCategory = true
				}
			}
			if newQuantity == 0 {
				fmt.Println("No change to quantity made")
			} else {
				if isChangeName == true && isChangeCategory == false {
					nameMap[newName] = itemInformation{element.catType, newQuantity, element.cost}
					isChangeQuantity = true
				} else if isChangeName == false && isChangeCategory == true {
					nameMap[key] = itemInformation{catNum, newQuantity, element.cost}
					isChangeQuantity = true
				} else if isChangeName == true && isChangeCategory == true {
					nameMap[newName] = itemInformation{catNum, newQuantity, element.cost}
					isChangeQuantity = true
				} else {
					nameMap[key] = itemInformation{element.catType, newQuantity, element.cost}
					isChangeQuantity = true
				}
			}
			if newCost == 0.0 {
				fmt.Println("No change to unit cost made")
			} else {
				if isChangeName == true && isChangeCategory == false && isChangeQuantity == false {
					nameMap[newName] = itemInformation{element.catType, element.quantity, newCost}
				} else if isChangeName == true && isChangeCategory == true && isChangeQuantity == false {
					nameMap[newName] = itemInformation{catNum, element.quantity, newCost}
				} else if isChangeName == true && isChangeCategory == true && isChangeQuantity == true {
					nameMap[newName] = itemInformation{catNum, newQuantity, newCost}
				} else if isChangeName == false && isChangeCategory == true && isChangeQuantity == true {
					nameMap[key] = itemInformation{catNum, newQuantity, newCost}
				} else if isChangeName == false && isChangeCategory == false && isChangeQuantity == true {
					nameMap[key] = itemInformation{element.catType, newQuantity, newCost}
				} else if isChangeName == false && isChangeCategory == true && isChangeQuantity == false {
					nameMap[key] = itemInformation{catNum, element.quantity, newCost}
				} else if isChangeName == true && isChangeCategory == false && isChangeQuantity == true {
					nameMap[newName] = itemInformation{element.catType, newQuantity, newCost}
				} else {
					nameMap[key] = itemInformation{element.catType, element.quantity, newCost}
				}
			}
		}
	}

	PressToContinue()
}

//DeleteItem perform option 5: Delete Item

func DeleteItem() {
	var userInput string
	var isFound = false
	fmt.Println("Delete Item")
	PrintBreak()
	fmt.Println("Enter item name to delete:")
	_, _ = fmt.Scanln(&userInput)

	for key := range nameMap {
		if userInput == key {
			isFound = true
		}
	}
	if isFound == true {
		delete(nameMap, userInput)
		fmt.Println("Deleted", userInput)
	} else {
		fmt.Println("Item not found. Nothing to delete!")
	}
	PressToContinue()
}

//PrintData perform option 5: Print Current Data
func PrintData() {
	fmt.Println("Print Current Data")
	PrintBreak()
	for key, element := range nameMap {
		fmt.Println(key, element)
	}
	PressToContinue()
}

func AddNewCategoryName() {
	var matchAny = false
	var userInput string
	var indexPlace int
	fmt.Println("Add New Category Name")
	PrintBreak()
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
		PressToContinue()
	} else {

		category = append(category, userInput)
		fmt.Printf("New category: %v added at index %v\n", userInput, len(category)-1)
		PressToContinue()
	}
}
