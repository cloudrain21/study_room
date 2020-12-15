package main

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
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

		var got []Player
		err := json.NewDecoder(response.Body).Decode(&got)
		if err != nil {
			t.Fatalf("decode body fail: %s", err)
		}
		assert.Equal(t, 200, response.Code)
	})

	t.Run("it returns player results", func(t *testing.T) {
		want := []Player {
			{"Cleo", 32},
			{"Chris", 20},
			{"Tiest", 14},
		}

		stubStore := &StubPlayerStore{league:want}
		server := NewPlayerServer(stubStore)

		request := httptest.NewRequest(http.MethodGet, "/league", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		var got []Player
		err := json.NewDecoder(response.Body).Decode(&got)
		if err != nil {
			t.Fatalf("decode body fail : %s", err)
		}

		assert.Equal(t, want, got)
	})

	t.Run("check content-type header", func(t *testing.T) {
		want := []Player{
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
		var want []Player

		memStore := NewInMemoryPlayerStore()
		server := NewPlayerServer(memStore)

		request := httptest.NewRequest(http.MethodGet, "/league", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		var got []Player
		err := json.NewDecoder(response.Body).Decode(&got)
		assert.Nil(t, err)

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

		want := []Player {
			{"Pepper", 3},
		}

		var got []Player
		err := json.NewDecoder(response.Body).Decode(&got)
		assert.Nil(t, err)

		assert.Equal(t, want, got)
	})
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
