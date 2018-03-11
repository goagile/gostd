package cake

const MAX_NUM_OF_ITEMS int = 6

func NewDinerMenu() *DinerMenu {
  m := new(DinerMenu)
  m.NumberOfItems = 0
  m.AddItem("Soup of the day", "", true, 3.05)
  return m
}

type DinerMenu struct {
  NumberOfItems int
  MenuItems [MAX_NUM_OF_ITEMS]MenuItem
}

func (m *DinerMenu) AddItem(Name, Description string, Vegeterian bool,
	Price float64) {
    if m.NumberOfItems >= MAX_NUM_OF_ITEMS {
      panic("Sorry, menu is full")
    }
		mi := MenuItem{Name, Description, Vegeterian, Price}
		m.MenuItems[m.NumberOfItems] = mi
    m.NumberOfItems += 1
}

func (m *DinerMenu) Iterator() *DinerIterator {
	it := new(DinerIterator)
	it.Position = 0
	it.Items = m.MenuItems
	return it
}

type DinerIterator struct {
  Position int
  Items [MAX_NUM_OF_ITEMS]MenuItem
}

func (it *DinerIterator) Next() MenuItem {
	mi := it.Items[it.Position]
	it.Position += 1
	return mi
}

func (it *DinerIterator) HasNext() bool {
	return it.Position < len(it.Items)
}
