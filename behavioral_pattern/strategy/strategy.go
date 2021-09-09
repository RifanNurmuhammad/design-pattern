package strategy

import "time"

type Interface interface {
	CalculateCost(distance float64) float64
	CalculateETA(startAt time.Time, trip TripDetail) time.Time
}

type Location struct {
	CityName  string
	Lat, Lang float64
}

type TripDetail struct {
	Origin, Destination Location
}

func (t TripDetail) Distance() float64 {
	return 10
}
