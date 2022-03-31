package menu

import "fmt"

//reusable misc func here

func printBreak() {
	fmt.Println("=========================")
}

func returnMainOrExit() {
	var option int
	fmt.Println("1. Main Menu")
	fmt.Println("2. Quit")
	_, _ = fmt.Scanln(&option)
	switch option {
	case 1:
		MainMenu()
	case 2:
		break
	}
}
