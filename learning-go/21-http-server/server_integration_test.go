package main

import (
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	store := NewInMemoryPlayerStore()
	server := PlayerServer{store, sync.Mutex{}}
	player := "Pepper"

	wantedCount := 50
	var wg sync.WaitGroup
	wg.Add(wantedCount)

	for i := 0; i < wantedCount; i++ {
		go func() {
			server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
			wg.Done()
		}()
	}
	wg.Wait()

	response := httptest.NewRecorder()
	server.ServeHTTP(response, newGetScoreRequest(player))
	assertStatus(t, response.Code, http.StatusOK)

	assertResponseBody(t, response.Body.String(), "50")
}
