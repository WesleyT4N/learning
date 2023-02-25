package poker

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	database, cleanDatabase := createTempFile(t, `[]`)
	defer cleanDatabase()

	store, err := NewFileSystemPlayerStore(database)
	if err != nil {
		t.Fatalf(err.Error())
	}

	server := NewPlayerServer(store)
	player := "Pepper"

	wantedCount := 10
	var wg sync.WaitGroup
	wg.Add(wantedCount)

	for i := 0; i < wantedCount; i++ {
		go func() {
			server.ServeHTTP(httptest.NewRecorder(), NewPostWinRequest(player))
			wg.Done()
		}()
	}
	wg.Wait()

	t.Run("get score", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, NewGetScoreRequest(player))
		AssertStatus(t, response.Code, http.StatusOK)

		AssertResponseBody(t, response.Body.String(), fmt.Sprint(wantedCount))
	})

	t.Run("get league", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, NewLeagueRequest())
		AssertStatus(t, response.Code, http.StatusOK)

		got := GetLeagueFromResponse(t, response.Body)
		want := League{
			{"Pepper", 10},
		}

		AssertLeague(t, got, want)
	})
}
