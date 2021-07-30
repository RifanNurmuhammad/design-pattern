package facade

import (
	"fmt"
	"testing"
)

func TestTakeOrder(t *testing.T) {
	data := Order{PaymentType: "bank", Amount: 2000}
	fmt.Println(TakeOrder(data))
}
