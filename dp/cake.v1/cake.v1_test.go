package cake

import (
  "testing"
  "fmt"
)

func TestMenuItemString(t *testing.T) {
  want := `Name: "Pancake Breakfast"`

  item := MenuItem{
    Name: "Pancake Breakfast",
    Description: "Eggs and toast",
    Vegetarian: true,
    Price: 2.99,
  }

  got := fmt.Sprintf("%v", item)

  if got != want {
    t.Errorf("Error in String \n\tgot: %v\n\twant: %v", got, want)
  }
}

func TestGetMenuItems(t *testing.T) {
  item0 := MenuItem{
    Name: "Regular Pancake Breakfast",
    Description: "",
    Vegetarian: false,
    Price: 2.99,
  }

  menu := new(PancakeHouseMenu)
  menu.AddItem("Regular Pancake Breakfast", "", false, 2.99)

  items := menu.MenuItems()
  if len(items) != 1 {
    t.Errorf("Error in len MenuItems")
  }

  if items[0] != item0 {
    t.Errorf("Error in MenuItems \n\tgot: %v\n\twant: %v", items[0], item0)
  }
}

// func TestAddMenuItem(t *testing.T) {
//   want := "Regular Pancake Breakfast"
//
//   menu := new(PancakeHouseMenu)
//   menu.AddItem("Regular Pancake Breakfast", "", false, 2.99)
//
//   if got := menu.GetItemString(0); got != want {
//     t.Errorf("Error in String \n\tgot: %v\n\twant: %v", got, want)
//   }
// }
