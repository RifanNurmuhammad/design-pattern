package command

// Command is a behavioral design pattern that turns a request into a stand-alone object that contains all information about the request.
// This transformation lets you pass requests as a method arguments, delay or queue a request’s execution, and support undoable operations.

type Interface interface {
	Send()
}

type ShippingNextDayCommand struct {
	Courier Courier
	Params  ShippingNextDayParams
}

func (s ShippingNextDayCommand) Send() {
	s.Courier.SendNextDayDelivery(s.Params)
}

type ShippingRegularCommand struct {
	Courier Courier
	Params  ShippingRegularParams
}

func (s ShippingRegularCommand) Send() {
	s.Courier.SendDeliveryRegular(s.Params)
}
