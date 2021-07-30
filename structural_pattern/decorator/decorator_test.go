package decorator

import (
	"fmt"
	"testing"
)

func TestInit(t *testing.T) {

	c := Coffee{}
	fmt.Println(c.GetName())
	fmt.Println(c.GetPrice())

	fmt.Println(" ======= ")

	latte := LatteCofee{
		Drink: c,
	}
	fmt.Println(latte.GetName())
	fmt.Println(latte.GetPrice())

	fmt.Println(" ======= ")

	boba := BobaCofee{
		Drink: latte,
	}
	fmt.Println(boba.GetName())
	fmt.Println(boba.GetPrice())

	fmt.Println(" ======= ")

	withBoba := BobaCofee{
		Drink: c,
	}
	fmt.Println(withBoba.GetName())
	fmt.Println(withBoba.GetPrice())

}
