package poker_test

import (
	"bytes"
	"fmt"
	poker "learning_go/21-http-server"
	"strings"
	"testing"
	"time"
)

func TestCLI(t *testing.T) {
	t.Run("starts game with 3 players and finishes game with 'Chris'", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		in := strings.NewReader("3\nChris wins\n")
		game := &GameSpy{}

		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		assertMessageSentToUser(t, stdout, poker.PlayerPrompt)
		assertGameStartedWith(t, game, 3)
		assertGameFinishedWith(t, game, "Chris")
	})
	t.Run("starts game with 7 players and finiishes game with Cleo", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		in := strings.NewReader("7\nCleo wins\n")
		game := &GameSpy{}

		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		assertMessageSentToUser(t, stdout, poker.PlayerPrompt)
		assertGameStartedWith(t, game, 7)
		assertGameFinishedWith(t, game, "Cleo")
	})
	t.Run("it prints an error when a non numeric value is enetered and does not start the game", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		in := strings.NewReader("Pies\n")
		game := &GameSpy{}

		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		assertGameNotStarted(t, game)
		assertMessageSentToUser(t, stdout, poker.PlayerPrompt, poker.BadPlayerInputErrorMsg)
	})
	t.Run("prints an error if user types invalid win string", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		in := strings.NewReader("7\nkekw\n")
		game := &GameSpy{}
		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		assertGameStartedWith(t, game, 7)
		assertMessageSentToUser(t, stdout, poker.PlayerPrompt, poker.BadWinnerInputErrorMessage)
	})
}

type ScheduledAlert struct {
	at     time.Duration
	amount int
}

func (s ScheduledAlert) String() string {
	return fmt.Sprintf("%d chips at %v", s.amount, s.at)
}

type SpyBlindAlerter struct {
	alerts []ScheduledAlert
}

func (s *SpyBlindAlerter) ScheduleAlertAt(duration time.Duration, amount int) {
	s.alerts = append(s.alerts, ScheduledAlert{duration, amount})
}

func assertScheduledAlert(t testing.TB, got, want ScheduledAlert) {
	t.Helper()
	if got.amount != want.amount {
		t.Errorf("got amount %d, want %d", got.amount, want.amount)
	}

	if got.at != want.at {
		t.Errorf("got scheduled time of %v, want %v", got.at, want.at)
	}
}

func assertMessageSentToUser(t testing.TB, stdout *bytes.Buffer, messages ...string) {
	t.Helper()
	want := strings.Join(messages, "")
	got := stdout.String()

	if got != want {
		t.Errorf("got %q sent to stdout but expected %+v", got, messages)
	}
}

func assertGameStartedWith(t testing.TB, game *GameSpy, want int) {
	t.Helper()
	if game.StartedWith != want {
		t.Errorf("wanted Start called with %d, got %d", want, game.StartedWith)
	}
}

func assertGameFinishedWith(t testing.TB, game *GameSpy, want string) {
	t.Helper()
	if game.FinishedWith != want {
		t.Errorf("wanted Finish called with %s, got %s", want, game.FinishedWith)
	}
}

func assertGameNotStarted(t testing.TB, game *GameSpy) {
	t.Helper()
	if game.StartCalled {
		t.Errorf("game should not have started")
	}
}
