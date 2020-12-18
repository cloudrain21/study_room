package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Player struct {
	Name string
	Wins int
}

type PlayerStore interface {
	GetPlayerScore(string) int
	RecordWin(name string)
	GetLeague() []Player
}

type PlayerServer struct {
	store PlayerStore
	router *http.ServeMux
}

func NewPlayerServer(store PlayerStore) *PlayerServer {
	p := new(PlayerServer)

	p.store = store
	p.router = http.NewServeMux()
	p.router.Handle("/league", http.HandlerFunc(p.leagueHandler))
	p.router.Handle("/players/", http.HandlerFunc(p.playerHandler))

	return p
}

func (p *PlayerServer)ServeHTTP(w http.ResponseWriter, r *http.Request) {
	p.router.ServeHTTP(w,r)
}

func (p *PlayerServer)leagueHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(p.store.GetLeague())

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (p *PlayerServer)playerHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.String()[len("/players/"):]
	//name := strings.TrimPrefix(r.URL.String(), "/players/")

	switch r.Method {
	case http.MethodGet:
		p.showScore(w,name)
	case http.MethodPost:
		p.processWin(w,name)
	}
}

func (p *PlayerServer)showScore(w http.ResponseWriter, name string) {
	score := p.store.GetPlayerScore(name)
	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}
	fmt.Fprintf(w, "%d", score)
}

func (p *PlayerServer)processWin(w http.ResponseWriter, name string) {
	p.store.RecordWin(name)

	w.WriteHeader(http.StatusAccepted)
}

func main() {
	store := NewInMemoryPlayerStore()
	server := NewPlayerServer(store)

	err := http.ListenAndServe(":5000", server)
	if err != nil {
		log.Fatalf("can't start server : %s", err)
	}
}
