package abstract_factory

// Abstract Factory is a creational design pattern that lets you produce families of related objects
// without specifying their concrete classes.

type Headset interface {
	Name() string
	Price() float64
	IsBluetoothType() bool
	Color() string
}

type HandphoneType string

const (
	Gaming  HandphoneType = "gaming"
	Regular HandphoneType = "regular"
)

type Handphone interface {
	Name() string
	Price() float64
	HasNFC() bool
	WithChargerAdapter() bool
	HandphoneType() HandphoneType
}

type ElectronicFactory interface {
	CreateHeadset() Headset
	CreateHandphone() Handphone
}
