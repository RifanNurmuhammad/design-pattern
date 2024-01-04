package main

import "fmt"

type Car struct {
	HasLeatherInterior bool
	HasAWD             bool
	HasPowerFrontSeat  bool
	Specs              string
}

type CarOption func(*Car)

func WithLeatherInterior() CarOption {
	return func(c *Car) {
		c.HasLeatherInterior = true
	}
}

func WithPowerFrontSeat() CarOption {
	return func(c *Car) {
		c.HasPowerFrontSeat = true
	}
}

func WithAWD() CarOption {
	return func(c *Car) {
		c.HasAWD = true
	}
}

func WithSpecs(s string) CarOption {
	return func(c *Car) {
		c.Specs = s
	}
}

const (
	defaultSpec      = "lifestyle"
	defaultAwdOption = false
)

func NewCar(opts ...CarOption) *Car {
	car := &Car{
		HasAWD: defaultAwdOption,
		Specs:  defaultSpec,
	}

	for _, opt := range opts {
		opt(car)
	}
	return car
}

func main() {
	car := NewCar(WithLeatherInterior(), WithSpecs("prestige"), WithPowerFrontSeat())
	fmt.Println(car)
}
