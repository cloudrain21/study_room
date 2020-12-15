package main

type InMemoryPlayerStore struct {
	scores map[string]int
	league []Player
}

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{
		scores: map[string]int{},
		league: []Player{},
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

func (s *InMemoryPlayerStore)GetLeague() []Player {
	var league []Player
	for name, wins := range s.scores {
		league = append(league, Player{name,wins})
	}
	return league
}