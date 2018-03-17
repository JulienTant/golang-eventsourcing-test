package main

import "time"

// Events
type TrainWasAnnounced struct {
	ID       string
	From     string
	FromTime time.Time
	To       string
	ToTime   time.Time
}
type TrainHasLeft struct {
	When time.Time
}
type TrainHasMoved struct {
	Where Position
	When  time.Time
}
type TrainHasArrived struct {
	When time.Time
}
