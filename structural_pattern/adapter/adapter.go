package adapter

import (
	"fmt"

	"github.com/shopspring/decimal"
)

// Adapter is a structural design pattern that allows objects with incompatible interfaces to collaborate.

type USD struct {
	Value decimal.Decimal
}

func (u USD) CairkanUangDolar(usd USD) {
	fmt.Println("uang yang diambil: $" + usd.Value.String())
}

type IDR struct {
	Uang float64
}

type Rupiah struct{}

func (r Rupiah) Cairkan(i IDR) {
	s := fmt.Sprintf("%f", i.Uang)
	fmt.Println("uang yang diambil: Rp. " + s)
}

type ConverterUSDToIDR struct {
	Adaptee Rupiah
}

func (c ConverterUSDToIDR) Cairkan(u USD) {
	usd, _ := u.Value.Float64()

	usd = usd * 14500
	idr := IDR{}
	idr.Uang = usd
	c.Adaptee.Cairkan(idr)
}

func Init() {
	u, _ := decimal.NewFromString("23.2")
	usd := USD{Value: u}

	idr := Rupiah{}

	adapter := ConverterUSDToIDR{Adaptee: idr}
	adapter.Cairkan(usd)
}

type ConverterIDRToUSD struct {
	Adaptee USD
}

func (c ConverterIDRToUSD) CairkanUangDolar(i IDR) {
	i.Uang = i.Uang / float64(14500)
	v := decimal.NewFromFloat(i.Uang)
	usd := USD{}
	usd.Value = v
	c.Adaptee.CairkanUangDolar(usd)
}

func InitIDRtoUSD() {

	idr := IDR{Uang: float64(290000)}

	usd := USD{}
	adapter := ConverterIDRToUSD{Adaptee: usd}
	adapter.CairkanUangDolar(idr)
}
