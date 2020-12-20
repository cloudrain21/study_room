package main

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestLeague(t *testing.T) {
	//t.Run("it returns 200 on /league", func(t *testing.T){
	//	store := &StubPlayerStore{}
	//	server := &PlayerServer{store:store}
	//
	//	request := httptest.NewRequest(http.MethodGet, "/league", nil)
	//	response := httptest.NewRecorder()
	//
	//	server.ServeHTTP(response, request)
	//
	//	got := response.Code
	//	want := http.StatusOK
	//
	//	assertStatus(t, got, want)
	//})

	t.Run("use servermux", func(t *testing.T){
		store := NewInMemoryPlayerStore()
		server := NewPlayerServer(store)

		request := httptest.NewRequest(http.MethodGet, "/league", nil)
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)
		assertStatus(t, response.Code, http.StatusOK)

		request = httptest.NewRequest(http.MethodPost, "/players/cloudrain", nil)
		response = httptest.NewRecorder()
		server.ServeHTTP(response, request)
		assertStatus(t, response.Code, http.StatusAccepted)

		request = httptest.NewRequest(http.MethodGet, "/players/cloudrain", nil)
		response = httptest.NewRecorder()
		server.ServeHTTP(response, request)
		assert.Equal(t, "1", response.Body.String(), "body should be the same")
	})

	t.Run("it returns 200 on /league", func(t *testing.T) {
		stubStore := &StubPlayerStore{}
		server := NewPlayerServer(stubStore)

		request := httptest.NewRequest(http.MethodGet, "/league", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		var got League
		err := json.NewDecoder(response.Body).Decode(&got)
		if err != nil {
			t.Fatalf("decode body fail: %s", err)
		}
		assert.Equal(t, 200, response.Code)
	})

	t.Run("it returns player results", func(t *testing.T) {
		want := League {
			{"Cleo", 32},
			{"Chris", 20},
			{"Tiest", 14},
		}

		stubStore := &StubPlayerStore{league:want}
		server := NewPlayerServer(stubStore)

		request := httptest.NewRequest(http.MethodGet, "/league", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		var got League
		err := json.NewDecoder(response.Body).Decode(&got)
		if err != nil {
			t.Fatalf("decode body fail : %s", err)
		}

		assert.Equal(t, want, got)
	})

	t.Run("check content-type header", func(t *testing.T) {
		want := League{
			{"Chris",10},
		}

		stubStore := &StubPlayerStore{league:want}
		server := NewPlayerServer(stubStore)

		request := httptest.NewRequest(http.MethodGet, "/league", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assert.Equal(t, "application/json", response.Header().Get("content-type"))
	})

	t.Run("inmemory league test", func(t *testing.T) {
		var want League

		memStore := NewInMemoryPlayerStore()
		server := NewPlayerServer(memStore)

		request := httptest.NewRequest(http.MethodGet, "/league", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		var got League
		err := json.NewDecoder(response.Body).Decode(&got)
		assert.Nil(t, err)

		assert.Equal(t, want, got)
	})
}

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

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	store := NewInMemoryPlayerStore()
	server := NewPlayerServer(store)
	player := "Pepper"

	server.ServeHTTP(httptest.NewRecorder(), myRequestPost(player))
	server.ServeHTTP(httptest.NewRecorder(), myRequestPost(player))
	server.ServeHTTP(httptest.NewRecorder(), myRequestPost(player))

	t.Run("get score", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, myRequestGet(player))

		want := "3"
		assert.Equal(t, want, response.Body.String())
	})

	t.Run("get league", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, getLeagueRequest(player))

		want := League {
			{"Pepper", 3},
		}

		var got League
		err := json.NewDecoder(response.Body).Decode(&got)
		assert.Nil(t, err)

		assert.Equal(t, want, got)
	})
}

func TestFileSystemStore(t *testing.T) {
//	t.Run("filesystem store test", func(t *testing.T) {
//		database := strings.NewReader(`[
//{"Name":"Cleo", "Wins":10},
//{"Name":"Chris", "Wins":33}]`)
//
//		store := FileSystemPlayerStore{database:database}
//
//		got := store.GetLeague()
//		got = store.GetLeague()
//
//		want := League {
//			{"Cleo",10},
//			{"Chris",33},
//		}
//
//		assert.Equal(t, want, got)
//	})
//
//	t.Run("get player store", func(t *testing.T){
//		database := strings.NewReader(`[
//{"Name":"Cleo", "Wins":10},
//{"Name":"Chris", "Wins":33}]`)
//
//		store := FileSystemPlayerStore{database:database}
//
//		got := store.GetPlayerScore("Chris")
//		want := 33
//
//		assert.Equal(t, want, got)
//	})

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
