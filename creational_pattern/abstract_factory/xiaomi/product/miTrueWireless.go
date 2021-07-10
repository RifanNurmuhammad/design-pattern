package product

type MiTrueWireless struct{}

func (m MiTrueWireless) Name() string {
	return "Mi True Wireless"
}

func (m MiTrueWireless) Price() float64 {
	return 399000
}

func (m MiTrueWireless) IsBluetoothType() bool {
	return true
}

func (m MiTrueWireless) Color() string {
	return "white"
}
