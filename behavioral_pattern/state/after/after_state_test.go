package after_state

import (
	"fmt"
	"testing"
)

func TestOrder(t *testing.T) {
	order := NewOrder()
	order.ChangeState(&B2BOrder{order})
	order.Send()
	fmt.Println(order.GetPrefix())
}
