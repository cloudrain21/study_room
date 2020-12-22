package poker

import (
	"github.com/stretchr/testify/assert"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)


func TestTape_Write(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		file, clean := createTempFile(t, "12345")
		defer clean()

		tape := &tape{file}
		tape.Write([]byte("abc"))

		file.Seek(0,io.SeekStart)
		newFileContents, _ := ioutil.ReadAll(file)

		got := string(newFileContents)
		want := "abc"

		assert.Equal(t, want, got)
	})
}

func TestFileSystemStore(t *testing.T) {
	t.Run("get player score from tempfile store", func(t *testing.T) {
		file, remf := createTempFile(t, `[
{"Name":"Cleo", "Wins":10},
{"Name":"Chris", "Wins":33}]`)

		store, _ := NewFileSystemPlayerStore(file)
		defer remf()

		got := store.GetPlayerScore("Chris")
		want := 33

		assert.Equal(t, want, got)
	})

	t.Run("get league from tempfile store", func(t *testing.T) {
		file, remf := createTempFile(t, `[
{"Name":"Cleo", "Wins":10},
{"Name":"Chris", "Wins":33}]`)
		defer remf()

		store, _ := NewFileSystemPlayerStore(file)
		got := store.GetLeague()
		want := League {
			{"Chris", 33},
			{"Cleo", 10},
		}

		assert.Equal(t, want, got)
	})

	t.Run("record win tempfile store", func(t *testing.T) {
		file, remf := createTempFile(t, `[
{"Name":"Cleo", "Wins":10},
{"Name":"Chris", "Wins":33}]`)
		defer remf()

		store, _ := NewFileSystemPlayerStore(file)

		playerName := "Chris"
		store.RecordWin(playerName)

		got := store.GetPlayerScore(playerName)
		want := 34

		assert.Equal(t, want, got)
	})

	t.Run("store wins for new players", func(t *testing.T) {
		database, remf := createTempFile(t, `[
{"Name":"Cleo", "Wins":10},
{"Name":"Chris", "Wins":33}]`)
		defer remf()

		store, _ := NewFileSystemPlayerStore(database)
		store.RecordWin("Pepper")

		got := store.GetPlayerScore("Pepper")
		want := 1

		assert.Equal(t, want, got)
	})

	t.Run("json test", func(t *testing.T) {
		database, remf := createTempFile(t, `[]`)
		defer remf()

		store, _ := NewFileSystemPlayerStore(database)
		store.RecordWin("Pepper")
		store.RecordWin("Pepper")
		store.RecordWin("Pepper")

		got := store.GetPlayerScore("Pepper")
		want := 3

		assert.Equal(t, want, got)
	})

	t.Run("empty file", func(t *testing.T) {
		database, remf := createTempFile(t, "")
		defer remf()

		store, err := NewFileSystemPlayerStore(database)
		if err != nil {
			t.Fatal(err)
		}

		store.RecordWin("Pepper")
		store.RecordWin("Pepper")
		store.RecordWin("Pepper")

		got := store.GetPlayerScore("Pepper")
		want := 3

		assert.Equal(t, want, got)
	})

	t.Run("league sorted", func(t *testing.T) {
		database, removef := createTempFile(t, `[
{"Name":"Cleo", "Wins":10},
{"Name":"Chris", "Wins":33}
]`)
		defer removef()

		store, err := NewFileSystemPlayerStore(database)
		if err != nil {
			t.Fatal(err)
		}

		got := store.GetLeague()
		want := League{
			{"Chris", 33},
			{"Cleo", 10},
		}

		assert.Equal(t, want, got)
	})
}

func createTempFile(t *testing.T, initData string) (*os.File, func()) {
	t.Helper()

	tmpfile, err := ioutil.TempFile("", "db")
	if err != nil {
		t.Fatalf("create temp file : %v", err)
	}

	tmpfile.Write([]byte(initData))

	removeFile := func() {
		tmpfile.Close()
		os.Remove(tmpfile.Name())
	}

	return tmpfile, removeFile
}

func myRequestPost(name string) *http.Request {
	url := "/players/" + name
	return httptest.NewRequest(http.MethodPost, url, nil)
}

func myRequestGet(name string) *http.Request {
	url := "/players/" + name
	return httptest.NewRequest(http.MethodGet, url, nil)
}

func getLeagueRequest(name string) *http.Request {
	url := "/league"
	return httptest.NewRequest(http.MethodGet, url, nil)
}

func assertStatus(t *testing.T, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func assertResponseBody(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func NewRequestWithName(method, name string) (*http.Request) {
	url := "/players/" + name
	req, _ := http.NewRequest(method, url, nil)
	return req
}
