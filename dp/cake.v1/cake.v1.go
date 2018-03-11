package cake

import (
	"fmt"
)

type MenuItem struct {
		Name string
		Description string
		Vegetarian bool
		Price float64
}

func (mi MenuItem) String() string {
	return fmt.Sprintf("Name: %q", mi.Name)
}

type PancakeHouseMenu struct {
	menuItems []MenuItem
}

func (m *PancakeHouseMenu) MenuItems() []MenuItem {
	return m.menuItems
}

func (m *PancakeHouseMenu) AddItem(
	Name, Description string, Vegeterian bool, Price float64) {
		mi := MenuItem{Name, Description, Vegeterian, Price}
		m.menuItems = append(m.menuItems, mi)
}

func (m *PancakeHouseMenu) GetItemString(index int) string {
	// if (len(m.menuItems) - 1) < index {
	// 	return ""
	// }
	return fmt.Sprintf("%v", m.menuItems[index])
}
