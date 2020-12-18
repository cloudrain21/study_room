package main

import (
	"encoding/json"
	"fmt"
	"io"
)

type FileSystemPlayerStore struct {
	database io.ReadWriteSeeker
}

func (f *FileSystemPlayerStore)GetLeague() []Player {
	f.database.Seek(0,io.SeekStart)
	players, _ := NewLeague(f.database)

	return players
}

func (f *FileSystemPlayerStore)GetPlayerScore(name string) (wins int) {
	for _, player := range f.GetLeague() {
		if name == player.Name {
			wins = player.Wins
			break
		}
	}
	return wins
}

func (f *FileSystemPlayerStore)RecordWin(name string) {
	league := f.GetLeague()

	for i, player := range league {
		if player.Name == name {
			league[i].Wins++
			break
		}
	}

	f.database.Seek(0, io.SeekStart)
	json.NewEncoder(f.database).Encode(league)
}

func NewLeague(r io.Reader) ([]Player, error) {
	var league []Player
	err := json.NewDecoder(r).Decode(&league)
	if err != nil {
		fmt.Errorf("decode error : %s", err)
	}

	return league, err
}