package strategy

import (
	"fmt"
	"time"
)

func NewContext(s Interface) *Context {
	return &Context{strategy: s}
}

type Context struct {
	strategy Interface
}

func (c *Context) SetStrategy(s Interface) {
	c.strategy = s
}

func (c Context) CalculateETA(startAt time.Time, td TripDetail) time.Time {
	return c.strategy.CalculateETA(startAt, td)
}

func Exec() {
	road := Road{}
	flight := Flight{wp: &WPImpl{}}

	now := time.Now()
	fmt.Println(now)

	ctx := NewContext(road)
	t1 := ctx.CalculateETA(now, TripDetail{
		Origin: Location{
			CityName: "Sukabumi",
		},
		Destination: Location{
			CityName: "Cengkareng",
		},
	})

	fmt.Println(t1)

	ctx.SetStrategy(flight)
	t2 := ctx.CalculateETA(t1, TripDetail{
		Origin: Location{
			CityName: "Cengkareng",
		},
		Destination: Location{
			CityName: "Surabaya",
		},
	})

	fmt.Println(t2)
}
