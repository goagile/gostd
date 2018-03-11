package cake

import "sort"

func NewIceCreamMenu() *IceCreamMenu {
  m := new(IceCreamMenu)
  m.MenuItems = make(map[string]MenuItem)
  m.AddItem("Agggggrh snow", "", false, 2.00)
  return m
}

type IceCreamMenu struct {
  MenuItems map[string]MenuItem
}

func (m *IceCreamMenu) AddItem(Name, Description string, Vegetarian bool,
  Price float64) {
  m.MenuItems[Name] = MenuItem{Name, Description, Vegetarian, Price}
}

func (m *IceCreamMenu) Iterator() Iterator {
  it := new(IceCreamIterator)
  it.Position = 0
  it.Items = m.MenuItems
  it.SortKeys()
  return it
}

type IceCreamIterator struct{
  Position int
  Items map[string]MenuItem
  Keys []string
}

func (it *IceCreamIterator) SortKeys() {
  var keys []string
  for key := range it.Items {
    keys = append(keys, key)
  }
  sort.Strings(keys)
  it.Keys = keys
}

func (it *IceCreamIterator) Next() MenuItem {
  key := it.Keys[it.Position]
  mi, ok := it.Items[key]
  if !ok {
    panic("Error with index")
  }
  it.Position += 1
  return mi
}

func (it *IceCreamIterator) HasNext() bool {
  return it.Position < len(it.Keys)
}
