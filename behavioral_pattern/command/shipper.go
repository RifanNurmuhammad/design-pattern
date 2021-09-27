package command

type Shipper struct {
	Commands []Interface
}

func (s *Shipper) AddShipping(i Interface) {
	s.Commands = append(s.Commands, i)
}

func (s Shipper) BulkShipping() {
	for _, i := range s.Commands {
		i.Send()
	}
}
