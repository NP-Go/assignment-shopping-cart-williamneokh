package handler

import (
	"assignment-shopping-cart-williamneokh/internal/config"
	"fmt"
)

//ViewShopList perform option 1: View Entire Shopping List
func ViewShopList() {
	fmt.Println("Shopping List Contents:")
	PrintBreak()
	for key, element := range config.NameMap {
		fmt.Printf("Category: %v - Item: %v Quantity: %v Unit Cost: %v\n", config.Category[element.CatType], key, element.Quantity, element.Cost)

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
		for i, catName := range config.Category {
			for _, element := range config.NameMap {
				if i == element.CatType {
					sum := float64(element.Quantity) * element.Cost
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
		for i, cat := range config.Category {
			for key, element := range config.NameMap {
				if i == element.CatType {

					fmt.Printf("Category: %v - Item: %v Quantity: %v Unit Cost: %v\n", cat, key, element.Quantity, element.Cost)
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

	for i, cat := range config.Category {
		if catName == cat {
			config.NameMap[itemName] = config.ItemInformation{i, quantityNum, itemCost}
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
	var foundItemName = false
	var foundCategory = false
	var catNum int
	fmt.Println("Modify Item")
	PrintBreak()
	fmt.Println("Which item would you wish to modify?")
	_, _ = fmt.Scanln(&userInput)
	//check userInput is valid
	for itemName := range config.NameMap {
		if itemName == userInput {
			foundItemName = true
		}
	}
	if foundItemName == false {
		fmt.Println("Item not found!")
	}

	for key, element := range config.NameMap {
		if userInput == key {
			fmt.Printf("Current item name is %v - Category is %v - Quantity is %v - Unit Cost %v\n", key, config.Category[element.CatType], element.Quantity, element.Cost)
			fmt.Println("Enter new name. Enter for no change")
			_, _ = fmt.Scanln(&newName)
			fmt.Println("Enter new Category. Enter for no change")
			_, _ = fmt.Scanln(&newCategory)
			//check if newCategory is valid
			if newCategory != "" {
				for _, catName := range config.Category {
					if catName == newCategory {
						foundCategory = true
					}
				}
				if foundCategory == false {
					fmt.Println("Category not found!")
					break
				}
			}

			fmt.Println("Enter new Quantity. Enter for no change")
			_, _ = fmt.Scanln(&newQuantity)
			fmt.Println("Enter new Unit cost. Enter for no change")
			_, _ = fmt.Scanln(&newCost)
			if newName == "" {
				fmt.Println("No change to item name made")
			} else {
				config.NameMap[newName] = config.NameMap[key]
				delete(config.NameMap, key)
				isChangeName = true
			}
			if newCategory == "" {
				fmt.Println("No change to category made")
			} else {

				for i, value := range config.Category {
					if newCategory == value {
						catNum = i
					}
				}
				if isChangeName == true {
					config.NameMap[newName] = config.ItemInformation{catNum, element.Quantity, element.Cost}
					isChangeCategory = true
				} else {
					config.NameMap[key] = config.ItemInformation{catNum, element.Quantity, element.Cost}
					isChangeCategory = true
				}
			}
			if newQuantity == 0 {
				fmt.Println("No change to quantity made")
			} else {
				if isChangeName == true && isChangeCategory == false {
					config.NameMap[newName] = config.ItemInformation{element.CatType, newQuantity, element.Cost}
					isChangeQuantity = true
				} else if isChangeName == false && isChangeCategory == true {
					config.NameMap[key] = config.ItemInformation{catNum, newQuantity, element.Cost}
					isChangeQuantity = true
				} else if isChangeName == true && isChangeCategory == true {
					config.NameMap[newName] = config.ItemInformation{catNum, newQuantity, element.Cost}
					isChangeQuantity = true
				} else {
					config.NameMap[key] = config.ItemInformation{element.CatType, newQuantity, element.Cost}
					isChangeQuantity = true
				}
			}
			if newCost == 0.0 {
				fmt.Println("No change to unit cost made")
			} else {
				if isChangeName == true && isChangeCategory == false && isChangeQuantity == false {
					config.NameMap[newName] = config.ItemInformation{element.CatType, element.Quantity, newCost}
				} else if isChangeName == true && isChangeCategory == true && isChangeQuantity == false {
					config.NameMap[newName] = config.ItemInformation{catNum, element.Quantity, newCost}
				} else if isChangeName == true && isChangeCategory == true && isChangeQuantity == true {
					config.NameMap[newName] = config.ItemInformation{catNum, newQuantity, newCost}
				} else if isChangeName == false && isChangeCategory == true && isChangeQuantity == true {
					config.NameMap[key] = config.ItemInformation{catNum, newQuantity, newCost}
				} else if isChangeName == false && isChangeCategory == false && isChangeQuantity == true {
					config.NameMap[key] = config.ItemInformation{element.CatType, newQuantity, newCost}
				} else if isChangeName == false && isChangeCategory == true && isChangeQuantity == false {
					config.NameMap[key] = config.ItemInformation{catNum, element.Quantity, newCost}
				} else if isChangeName == true && isChangeCategory == false && isChangeQuantity == true {
					config.NameMap[newName] = config.ItemInformation{element.CatType, newQuantity, newCost}
				} else {
					config.NameMap[key] = config.ItemInformation{element.CatType, element.Quantity, newCost}
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

	for key := range config.NameMap {
		if userInput == key {
			isFound = true
		}
	}
	if isFound == true {
		delete(config.NameMap, userInput)
		fmt.Println("Deleted", userInput)
	} else {
		fmt.Println("Item not found. Nothing to delete!")
	}
	PressToContinue()
}

//PrintData perform option 6: Print Current Data
func PrintData() {
	fmt.Println("Print Current Data")
	PrintBreak()
	for key, element := range config.NameMap {
		fmt.Println(key, element)
	}
	PressToContinue()
}

//AddNewCategoryName perform option 7: Add New Category Name
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
		PressToContinue()
	} else {
		for _, element := range config.NameMap {
			if userInput == config.Category[element.CatType] {
				matchAny = true
				indexPlace = element.CatType
			}
		}
		if matchAny == true {
			fmt.Printf("Category: %v already exit at index %v !\n", userInput, indexPlace)
			PressToContinue()
		} else {
			config.Category = append(config.Category, userInput)
			fmt.Printf("New category: %v added at index %v\n", userInput, len(config.Category)-1)
			PressToContinue()
		}
	}
}
