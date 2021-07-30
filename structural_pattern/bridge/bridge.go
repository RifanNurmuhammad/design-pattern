package bridge

import "fmt"

// Bridge is a structural design pattern
// that lets you split a large class or a set of closely related classes
// into two separate hierarchies—abstraction and implementation—which can be developed independently of each other.

type Electronic interface {
	Produce()
}

type Handphone struct {
	OS    OS
	Color Color
}

func (h Handphone) Produce() {
	h.OS.Install()

	if h.Color != nil {
		fmt.Println("Handphone with color: " + h.Color.Name())
	} else {
		fmt.Println("Handphone with no color")
	}
}

// Clone using prototype pattern
func (h Handphone) Clone() Handphone {
	return Handphone{
		OS:    h.OS,
		Color: h.Color,
	}
}

func Init() {
	h := Handphone{
		OS: Android{},
	}
	h.Produce()
	// Install OS Android
	// Handphone with no color

	h2 := Handphone{
		OS:    IOS{},
		Color: BlueOcean{},
	}
	h2.Produce()
	// Install OS IOS
	// Handphone with color: BlueOcean

	// combile clone using prototype pattern and bridge
	hp := h.Clone()
	hp.Color = BlueOcean{}
	hp.Produce()
	// Install OS Android
	// Handphone with color: BlueOcean
}
