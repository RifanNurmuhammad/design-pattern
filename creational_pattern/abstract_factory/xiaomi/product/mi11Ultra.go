package product

import (
	"github.com/rifannurmuhammad/design-pattern/creational_pattern/abstract_factory"
)

type Mi11Ultra struct{}

func (m Mi11Ultra) Name() string {
	return "Mi 11 Ultra"
}

func (m Mi11Ultra) Price() float64 {
	return 16999000
}
func (m Mi11Ultra) HasNFC() bool {
	return true
}
func (m Mi11Ultra) WithChargerAdapter() bool {
	return true
}
func (m Mi11Ultra) HandphoneType() abstract_factory.HandphoneType {
	return abstract_factory.Regular
}
