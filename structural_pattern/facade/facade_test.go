package facade

import (
	"fmt"
	"testing"
)

func TestTakeOrder(t *testing.T) {
	data := Order{PaymentType: "bank", Amount: 2000}
	fmt.Println(TakeOrder(data))
}

func TestTakeOrderEwallet(t *testing.T) {
	data := Order{PaymentType: "mobileWallet", Amount: 2000}
	fmt.Println(TakeOrder(data))
}

func TestTakeOrderCOD(t *testing.T) {
	data := Order{PaymentType: "cod", Amount: 2000}
	fmt.Println(TakeOrder(data))
}
