package before_state

type ChannelOrderState string

const (
	B2C ChannelOrderState = "B2C"
	B2B ChannelOrderState = "B2B"
	B2G ChannelOrderState = "B2G"
)

type Order struct {
	State  ChannelOrderState
	Prefix string
}

func (o Order) Send() {
	if o.State == B2B {
		o.Prefix = "B001"
	} else if o.State == B2C {
		o.Prefix = "C001"
	} else {
		o.Prefix = "G001"
	}
}
