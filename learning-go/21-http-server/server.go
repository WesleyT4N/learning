package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"sync"
)

type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
	GetLeague() []Player
}

type PlayerServer struct {
	store PlayerStore
	mu    sync.Mutex
	http.Handler
}

type Player struct {
	Name string
	Wins int
}

const jsonContentType = "application/json"

func NewPlayerServer(store PlayerStore) *PlayerServer {
	p := new(PlayerServer)
	p.store = store
	router := http.NewServeMux()

	router.Handle("/league", http.HandlerFunc(p.leagueHandler))
	router.Handle("/players/", http.HandlerFunc(p.playerHandler))

	p.Handler = router
	return p
}

func GetPlayerScore(name string) string {
	if name == "Pepper" {
		return "20"
	}
	if name == "Floyd" {
		return "10"
	}
	return ""
}

func (p *PlayerServer) leagueHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", jsonContentType)
	json.NewEncoder(w).Encode(p.store.GetLeague())
}

func (p *PlayerServer) playerHandler(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	switch r.Method {
	case http.MethodPost:
		p.procoessWin(w, player)
	case http.MethodGet:
		p.showScore(w, player)
	}
}

func (p *PlayerServer) showScore(w http.ResponseWriter, player string) {
	score := p.store.GetPlayerScore(player)

	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, score)
}

func (p *PlayerServer) procoessWin(w http.ResponseWriter, player string) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.store.RecordWin(player)
	w.WriteHeader(http.StatusAccepted)
}
