package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"
)

type PlayerStore interface {
	GetPlayerScore(string) int
	RecordWin(name string)
}

type PlayerServer struct {
	mu sync.Mutex
	store PlayerStore
}

func (p *PlayerServer)ServeHTTP(w http.ResponseWriter, r *http.Request) {
	name := strings.TrimPrefix(r.URL.String(), "/players/")

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
	p.mu.Lock()
	p.store.RecordWin(name)
	p.mu.Unlock()

	w.WriteHeader(http.StatusAccepted)
}

func main() {
	//StubPlayerStore := &StubPlayerStore {
	//	stores:map[string]int{
	//		"Pepper":20,
	//		"Floyd":10,
	//	},
	//}

	store := NewInMemoryPlayerStore()
	server := &PlayerServer{store:store}

	err := http.ListenAndServe(":5000", server)
	if err != nil {
		log.Fatalf("can't start server : %s", err)
	}
}