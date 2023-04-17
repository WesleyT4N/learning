package poker

import "time"

type Game interface {
	Start(numberOfPlayers int)
	Finish(winner string)
}

type TexasHoldEm struct {
	alerter BlindAlerter
	store   PlayerStore
}

func NewTexasHoldEm(alerter BlindAlerter, store PlayerStore) *TexasHoldEm {
	return &TexasHoldEm{
		alerter,
		store,
	}
}

func (g *TexasHoldEm) Start(numberOfPlayers int) {
	blindIncrement := time.Duration(5+numberOfPlayers) * time.Minute

	blinds := []int{100, 200, 300, 400, 500, 600, 800, 1000, 2000, 4000, 8000}
	blindTime := 0 * time.Second

	for _, blind := range blinds {
		g.alerter.ScheduleAlertAt(blindTime, blind)
		blindTime = blindTime + blindIncrement
	}
}

func (g *TexasHoldEm) Finish(winner string) {
	g.store.RecordWin(winner)
}
