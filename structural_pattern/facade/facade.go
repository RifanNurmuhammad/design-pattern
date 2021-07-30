package facade

// Facade is a structural design pattern that provides a simplified interface
// to a library, a framework, or any other complex set of classes.

type Order struct {
	PaymentType string
	Amount      int
}

type Response struct {
	Status  string
	Message string
	Data    Order
}

func TakeOrder(o Order) Response {
	switch o.PaymentType {
	case "mobileWallet":
		o.Amount = mobileWalletCalc(o.Amount)
	case "bank":
		o.Amount = bankCalc(o.Amount)
	case "cod":
		o.Amount = codCalc(o.Amount)
	}
	return Response{Status: "success", Message: "order already take", Data: o}
}

func mobileWalletCalc(amount int) int {
	return amount + 1000
}

func bankCalc(amount int) int {
	return amount + 6500
}

func codCalc(amount int) int {
	return amount + 0
}
