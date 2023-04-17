package poker_test

import (
	"fmt"
	poker "learning_go/21-http-server"
	"testing"
	"time"
)

type GameSpy struct {
	StartedWith  int
	FinishedWith string
	StartCalled  bool
}

func (g *GameSpy) Start(numberOfPlayers int) {
	g.StartedWith = numberOfPlayers
	g.StartCalled = true
}

func (g *GameSpy) Finish(winner string) {
	g.FinishedWith = winner
}

func TestTexasHoldEm_Start(t *testing.T) {
	t.Run("schedule alerts on game start for 5 players", func(t *testing.T) {
		playerStore := &poker.StubPlayerStore{}
		blindAlerter := &SpyBlindAlerter{}
		game := poker.NewTexasHoldEm(blindAlerter, playerStore)

		game.Start(5)
		cases := []ScheduledAlert{
			{0 * time.Second, 100},
			{10 * time.Minute, 200},
			{20 * time.Minute, 300},
			{30 * time.Minute, 400},
			{40 * time.Minute, 500},
			{50 * time.Minute, 600},
			{60 * time.Minute, 800},
			{70 * time.Minute, 1000},
			{80 * time.Minute, 2000},
			{90 * time.Minute, 4000},
			{100 * time.Minute, 8000},
		}

		checkSchedulingCases(cases, t, blindAlerter)
	})

	t.Run("schedules alerts on game start for 7 players", func(t *testing.T) {
		playerStore := &poker.StubPlayerStore{}
		blindAlerter := &SpyBlindAlerter{}
		game := poker.NewTexasHoldEm(blindAlerter, playerStore)

		game.Start(7)
		cases := []ScheduledAlert{
			{0 * time.Second, 100},
			{12 * time.Minute, 200},
			{24 * time.Minute, 300},
			{36 * time.Minute, 400},
		}

		checkSchedulingCases(cases, t, blindAlerter)
	})
}

func checkSchedulingCases(cases []ScheduledAlert, t *testing.T, blindAlerter *SpyBlindAlerter) {
	t.Helper()
	for i, want := range cases {
		t.Run(fmt.Sprint(want), func(t *testing.T) {
			if len(blindAlerter.alerts) <= i {
				t.Fatal("expected a blind alert to be scheduled")
			}

			got := blindAlerter.alerts[i]
			assertScheduledAlert(t, got, want)
		})
	}
}

func TestTexasHoldEm_Finish(t *testing.T) {
	store := &poker.StubPlayerStore{}
	dummySpyAlerter := &SpyBlindAlerter{}
	game := poker.NewTexasHoldEm(dummySpyAlerter, store)
	winner := "Ruth"

	game.Finish(winner)
	poker.AssertPlayerWin(t, store, winner)
}
