package menu

import "fmt"

func MainMenu() {
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
		viewEntireShoppingList()
	case 2:
		generateReport()
	case 3:
		addItem()
	case 4:
		modifyItem()
	case 5:
		deleteItem()
	case 6:
		printCurrentData()
	case 7:
		newCategory()
	default:
		MainMenu()

	}
}
