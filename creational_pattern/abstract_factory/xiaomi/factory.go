package xiaomi

import (
	"github.com/rifannurmuhammad/design-pattern/creational_pattern/abstract_factory"
	"github.com/rifannurmuhammad/design-pattern/creational_pattern/abstract_factory/xiaomi/product"
)

type Xiaomi struct{}

func (x *Xiaomi) CreateHeadset() abstract_factory.Headset {
	return &product.MiTrueWireless{}
}

func (x *Xiaomi) CreateHandphone() abstract_factory.Handphone {
	return &product.Mi11Ultra{}
}
