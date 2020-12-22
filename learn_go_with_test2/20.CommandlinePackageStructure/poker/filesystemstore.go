package poker

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sort"
)

type FileSystemPlayerStore struct {
	database *json.Encoder
	league League
}

func NewFileSystemPlayerStore(file *os.File) (*FileSystemPlayerStore,error) {
	file.Seek(0,io.SeekStart)

	info, err := file.Stat()
	if err != nil {
		return nil, fmt.Errorf("stat error : %s", err)
	}

	if info.Size() == 0 {
		file.Write([]byte("[]"))
		file.Seek(0,io.SeekStart)
	}

	league, err := LoadLeague(file)
	if err != nil {
		return nil, fmt.Errorf("loading (%s) error : %s", file.Name(), err)
	}

	return &FileSystemPlayerStore{
		json.NewEncoder(&tape{file}),
		league,
	}, nil
}

func NewFileSystemPlayerStoreFromFile(fileName string) (*FileSystemPlayerStore,func(),error) {
	file, err := os.OpenFile(fileName,os.O_RDWR|os.O_CREATE,0666)
	if err != nil {
		return nil, nil, err
	}

	closeFunc := func() {
		file.Close()
	}

	store, err := NewFileSystemPlayerStore(file)
	if err != nil {
		return nil, nil, err
	}

	return store, closeFunc, nil
}

func (f *FileSystemPlayerStore)GetLeague() League {
	sort.Slice(f.league, func(i, j int) bool {
		return f.league[i].Wins > f.league[j].Wins
	})
	return f.league
}

func (f *FileSystemPlayerStore)GetPlayerScore(name string) (wins int) {
	player := f.GetLeague().Find(name)
	if player != nil {
		return player.Wins
	}

	return 0
}

func (f *FileSystemPlayerStore)RecordWin(name string) {
	player := f.GetLeague().Find(name)
	if player != nil {
		player.Wins++
	} else {
		f.league = append(f.league, Player{name, 1})
	}

	//f.database.Seek(0, io.SeekStart)
	f.database.Encode(f.league)
}

func LoadLeague(r io.Reader) (League, error) {
	var league League
	err := json.NewDecoder(r).Decode(&league)
	if err != nil {
		fmt.Errorf("decode error : %s", err)
	}

	return league, err
}