package main

import (
	"fmt"
)

type itemInformation struct {
	catType  int
	quantity int
	cost     float64
}

var category []string
var nameMap map[string]itemInformation

func main() {
	PreLoadData()
	for {
		var choice int
		fmt.Println("Shopping List Application")
		PrintBreak()
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
			ViewShopList()
		case 2:
			GenerateShoppingListReport()
		case 3:
			AddItem()
		case 4:
			ModifyItems()
		case 5:
			DeleteItem()
		case 6:
			PrintData()
		case 7:
			AddNewCategoryName()

		}
	}
}
