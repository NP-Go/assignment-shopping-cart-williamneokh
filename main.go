package main

import (
	"assignment-shopping-cart-williamneokh/pkg/config"
	"assignment-shopping-cart-williamneokh/pkg/data"
	"assignment-shopping-cart-williamneokh/pkg/handler"
	"fmt"
)

func main() {
	config.LoadMap()
	data.PreLoadData() //Pre-load datasheet or comment out to remove datasheet
	for {
		var choice int
		fmt.Println("Shopping List Application")
		handler.PrintBreak()
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
			handler.ViewShopList()
		case 2:
			handler.GenerateShoppingListReport()
		case 3:
			handler.AddItem()
		case 4:
			handler.ModifyItems()
		case 5:
			handler.DeleteItem()
		case 6:
			handler.PrintData()
		case 7:
			handler.AddNewCategoryName()

		}
	}
}
