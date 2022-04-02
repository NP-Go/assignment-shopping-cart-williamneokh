package main

func PreLoadData() {
	nameMap = make(map[string]itemInformation)
	nameMap["Fork"] = itemInformation{0, 4, 3}
	nameMap["Plates"] = itemInformation{0, 4, 3}
	nameMap["Cups"] = itemInformation{0, 5, 3}
	nameMap["Bread"] = itemInformation{1, 2, 2}
	nameMap["Cake"] = itemInformation{1, 3, 1}
	nameMap["Coke"] = itemInformation{2, 5, 2}
	nameMap["Sprite"] = itemInformation{2, 5, 2}
	category = []string{"Household", "Food", "Drink"}
}
