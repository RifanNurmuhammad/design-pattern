package after_state

// State is a behavioral design pattern that lets an object alter its behavior when its internal state changes.
// It appears as if the object changed its class.
type State interface {
	Send()
}

func NewOrder() *Order {
	o := &Order{}
	o.ChangeState(&B2COrder{o})
	return o
}

type Order struct {
	State  State
	Prefix string
}

func (o Order) Send() {
	o.State.Send()
}

func (o *Order) ChangeState(state State) {
	o.State = state
}

func (o *Order) SetPrefix(prefix string) {
	o.Prefix = prefix
}

func (o *Order) GetPrefix() string {
	return o.Prefix
}

type B2COrder struct {
	Order *Order
}

func (c *B2COrder) Send() {
	c.Order.ChangeState(&B2COrder{Order: c.Order})
	c.Order.SetPrefix("C001")
}

type B2BOrder struct {
	Order *Order
}

func (b *B2BOrder) Send() {
	b.Order.ChangeState(&B2BOrder{Order: b.Order})
	b.Order.SetPrefix("B001")
}
