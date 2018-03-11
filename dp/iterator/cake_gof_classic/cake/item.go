package cake

import "fmt"

type MenuItem struct {
		Name string
		Description string
		Vegetarian bool
		Price float64
}

func (mi MenuItem) String() string {
	return fmt.Sprintf("$%v\t%q", mi.Price, mi.Name)
}
