package decorator

// Decorator is a structural design pattern that lets you attach new behaviors to objects by placing
// these objects inside special wrapper objects that contain the behaviors.

// agregation to component
type LatteCofee struct {
	Drink Drink
}

func (l LatteCofee) GetPrice() int {
	return l.Drink.GetPrice() + 500
}

func (l LatteCofee) GetName() string {
	return l.Drink.GetName() + " with latte"
}

type BobaCofee struct {
	Drink Drink
}

func (b BobaCofee) GetPrice() int {
	return b.Drink.GetPrice() + 200
}

func (b BobaCofee) GetName() string {
	return b.Drink.GetName() + " with Boba"
}
