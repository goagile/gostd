package cake

func NewPancakeHouseMenu() *PancakeHouseMenu {
	m := new(PancakeHouseMenu)
	m.AddItem("Regular Pancake Breakfast", "", false, 2.99)
	return m
}

type PancakeHouseMenu struct {
	MenuItems []MenuItem
}

func (m *PancakeHouseMenu) AddItem(Name, Description string, Vegeterian bool,
	Price float64) {
		mi := MenuItem{Name, Description, Vegeterian, Price}
		m.MenuItems = append(m.MenuItems, mi)
}

func (m *PancakeHouseMenu) Iterator() *PancakeIterator {
	it := new(PancakeIterator)
	it.Position = 0
	it.Items = m.MenuItems
	return it
}

type PancakeIterator struct {
  Position int
  Items []MenuItem
}

func (it *PancakeIterator) Next() MenuItem {
	mi := it.Items[it.Position]
	it.Position += 1
	return mi
}

func (it *PancakeIterator) HasNext() bool {
	return it.Position < len(it.Items)
}
