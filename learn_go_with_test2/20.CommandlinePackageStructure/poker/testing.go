package poker

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type StubPlayerStore struct {
	stores  map[string]int
	winCalls []string
	league League
}

func (s *StubPlayerStore)GetPlayerScore(name string) int {
	return s.stores[name]
}

func (s *StubPlayerStore)RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

func (s *StubPlayerStore)GetLeague() League {
	return s.league
}


func AssertPlayerWin(t *testing.T, store *StubPlayerStore, winner string) {
	t.Helper()

	if len(store.winCalls) < 1 {
		t.Fatal("must wincalls > 1")
	}

	got := store.winCalls[0]
	want := winner

	assert.Equal(t, want, got)
}
