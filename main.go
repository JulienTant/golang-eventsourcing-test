package main

import (
	"fmt"
	"math/rand"
	"time"

	log "gopkg.in/inconshreveable/log15.v2"
)

func main() {
	log.Info("generating events...")
	history := []interface{}{
		TrainWasAnnounced{
			ID:       "some-random-id",
			From:     "Paris",
			FromTime: time.Date(2018, 03, 12, 12, 0, 0, 0, time.Local),
			To:       "Marseille",
			ToTime:   time.Date(2018, 03, 12, 15, 30, 0, 0, time.Local),
		},
		TrainHasLeft{When: time.Date(2018, 03, 12, 12, 4, 0, 0, time.Local)},
	}

	for i := 0; i < 999997; i++ {
		history = append(history, TrainHasMoved{
			Where: Position{rand.Float64(), rand.Float64()},
			When:  time.Now(), // not coherent but who cares
		})
	}

	history = append(history, TrainHasArrived{time.Date(2018, 03, 12, 15, 39, 0, 0, time.Local)})
	log.Info(fmt.Sprintf("%d events generated", len(history)))

	log.Info("Applying events...")
	start := time.Now()
	t := NewTrainFromHistory(history)
	log.Info("Events applied in " + time.Now().Sub(start).String())
	fmt.Print(t.String())
}
