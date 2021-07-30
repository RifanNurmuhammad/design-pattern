package bridge

type Color interface {
	Name() string
}

type BlueOcean struct{}

func (b BlueOcean) Name() string {
	return "BlueOcean"
}
