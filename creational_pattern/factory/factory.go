package factory

import "fmt"

// Factory Method is a creational design pattern that provides an interface for creating objects in a superclass,
// but allows subclasses to alter the type of objects that will be created.

// The food declare the interfaces, whic is common to all objects that can be produced by the resto
type Food interface {
	// any food have a function order()
	Order()
}

// Food category as product
type JunkFood struct{}

// class implement order function
// only call Order will implement the same interface 'food'
// duck typing golang example, when call function Order() get to know Food interface
func (f *JunkFood) Order() {
	fmt.Println("your order is junk food")
}

// Food category
type HealthyFood struct{}

// class implement order function
func (f *HealthyFood) Order() {
	fmt.Println("your order is healthy food")
}

// ============================================ //

// interface for resto
type Resto interface {
	// function cook will produce food
	// inheritence
	Cook() Food
}

// factory
type MCD struct{}

func (r *MCD) Cook() Food {
	return &JunkFood{}
}

type Lemonilo struct{}

func (r *Lemonilo) Cook() Food {
	return &HealthyFood{}
}

func Init() {
	var resto Resto
	resto = &Lemonilo{}

	food := resto.Cook()
	food.Order()
}
