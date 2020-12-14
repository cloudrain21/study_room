package main

type StubPlayerStore struct {
	stores  map[string]int
	winCalls []string
}

func (s *StubPlayerStore)GetPlayerScore(name string) int {
	return s.stores[name]
}

func (s *StubPlayerStore)RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}