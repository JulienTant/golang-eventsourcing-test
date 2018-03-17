package main

import (
	"fmt"
	"time"
)

func NewTrainFromHistory(events []interface{}) *Train {
	train := &Train{}
	for i := range events {
		train.apply(events[i])
	}
	return train
}

// Aggregate
type Train struct {
	ID        string
	From      string
	FromTime  time.Time
	FromDelay time.Duration
	To        string
	ToTime    time.Time
	ToDelay   time.Duration
	Position  Position
}

func (t *Train) apply(event interface{}) {
	switch e := event.(type) {

	case TrainWasAnnounced:
		t.ID = e.ID
		t.From = e.From
		t.FromTime = e.FromTime
		t.To = e.To
		t.ToTime = e.ToTime

	case TrainHasLeft:
		t.FromDelay = e.When.Sub(t.FromTime)

	case TrainHasMoved:
		t.Position = e.Where

	case TrainHasArrived:
		t.ToDelay = e.When.Sub(t.ToTime)
	}
}

func (t Train) String() string {
	format := `Train: %s
	From: %s - %s (%s delay)
	To: %s - %s (%s delay)
	Current position: %s
`
	return fmt.Sprintf(format, t.ID, t.From, t.FromTime, t.FromDelay, t.To, t.ToTime, t.ToDelay, t.Position)
}
