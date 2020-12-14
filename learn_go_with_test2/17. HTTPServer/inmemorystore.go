package main

type InMemoryPlayerStore struct {
	scores map[string]int
}

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{
		scores: map[string]int{},
	}
}

func (s *InMemoryPlayerStore)GetPlayerScore(name string) int {
	if score, ok := s.scores[name]; ok {
		return score
	}
	return 0
}

func (s *InMemoryPlayerStore)RecordWin(name string) {
	s.scores[name]++
}