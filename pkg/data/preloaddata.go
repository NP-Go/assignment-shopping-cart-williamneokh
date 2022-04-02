package data

import "assignment-shopping-cart-williamneokh/pkg/config"

func PreLoadData() {

	config.NameMap["Fork"] = config.ItemInformation{0, 4, 3}
	config.NameMap["Plates"] = config.ItemInformation{0, 4, 3}
	config.NameMap["Cups"] = config.ItemInformation{0, 5, 3}
	config.NameMap["Bread"] = config.ItemInformation{1, 2, 2}
	config.NameMap["Cake"] = config.ItemInformation{1, 3, 1}
	config.NameMap["Coke"] = config.ItemInformation{2, 5, 2}
	config.NameMap["Sprite"] = config.ItemInformation{2, 5, 2}
	config.Category = []string{"Household", "Food", "Drink"}
}
