package command

import "fmt"

type ShippingNextDayParams struct {
	ExtraBubbleWraps bool
	Box              int
}

type ShippingRegularParams struct {
	Kargo bool
}

type Courier interface {
	SendNextDayDelivery(params ShippingNextDayParams)
	SendDeliveryRegular(params ShippingRegularParams)
}

type JNE struct{}

func (j JNE) SendNextDayDelivery(params ShippingNextDayParams) {
	fmt.Println("send delivery next day by JNE", params)
}

func (j JNE) SendDeliveryRegular(params ShippingRegularParams) {
	fmt.Println("send delivery regular by JNE", params)
}
