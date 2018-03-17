package main

import "strconv"

type Position struct {
	Latitude  float64
	Longitude float64
}

func (p Position) String() string {
	return strconv.FormatFloat(p.Latitude, 'f', -1, 64) + "," + strconv.FormatFloat(p.Longitude, 'f', -1, 64)
}
