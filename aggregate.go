package main

import (
	"fmt"
	"time"

	"github.com/satori/go.uuid"
)

func NewTrainFromHistory(events []interface{}) *Train {
	train := &Train{}
	for i := range events {
		train.apply(events[i])
	}
	return train
}

func AnnounceNewTrain(From string, FromTime time.Time, To string, ToTime time.Time) *Train {
	t := &Train{}
	t.Announce(uuid.NewV4().String(), From, FromTime, To, ToTime)
	return t
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

	changes []interface{}
	version int
}

func (t *Train) Announce(ID, From string, FromTime time.Time, To string, ToTime time.Time) {
	t.recordThat(TrainWasAnnounced{
		ID:       ID,
		From:     From,
		FromTime: FromTime,
		To:       To,
		ToTime:   ToTime,
	})
}

func (t *Train) Leaves(tm time.Time) {
	t.recordThat(TrainHasLeft{When: tm})
}

func (t *Train) Arrives(tm time.Time) {
	t.recordThat(TrainHasArrived{When: tm})
}

func (t *Train) Move(tm time.Time, p Position) {
	t.recordThat(TrainHasMoved{Where: p, When: tm})
}

func (t *Train) recordThat(event interface{}) {
	t.changes = append(t.changes, event)
	t.version++
	t.apply(event)
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
