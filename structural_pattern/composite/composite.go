package composite

// Composite is a structural design pattern that lets you compose objects into tree structures
// and then work with these structures as if they were individual objects.

type Product interface {
	GetPrice() int
	TotalPriceProduct() int
}

type PaketWorkingSpace struct {
	ProductItem []Product
}

func (pws PaketWorkingSpace) GetPrice() int {
	return 50
}

func (pws PaketWorkingSpace) TotalPriceProduct() int {

	var totalPriceItem int
	for _, product := range pws.ProductItem {
		totalPriceItem = totalPriceItem + product.TotalPriceProduct()
	}
	return pws.GetPrice() + totalPriceItem
}

type MejaKerja struct {
	ProductItem []Product
}

func (m MejaKerja) GetPrice() int {
	return 20
}

func (m MejaKerja) TotalPriceProduct() int {
	var totalPriceItem int
	for _, product := range m.ProductItem {
		totalPriceItem = totalPriceItem + product.GetPrice()
	}
	return m.GetPrice() + totalPriceItem
}

type Monitor struct {
	ProductItem []Product
}

func (m Monitor) GetPrice() int {
	return 5
}

func (m Monitor) TotalPriceProduct() int {
	return 0
}
