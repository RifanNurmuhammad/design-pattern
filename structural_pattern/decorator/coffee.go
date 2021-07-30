package decorator

// component interface
type Drink interface {
	GetPrice() int
	GetName() string
}

// concrete component
type Coffee struct{}

func (c Coffee) GetPrice() int {
	return 20000
}

func (c Coffee) GetName() string {
	return "Coffee"
}
