package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETPlayer(t *testing.T) {
	stubStore := &StubPlayerStore {
		stores: map[string]int{
			"Pepper":20,
			"Floyd":10,
		},
	}
	t.Run("returns Pepper's screo", func(t *testing.T) {
		request := NewRequestWithName(http.MethodGet, "Pepper")
		respose := httptest.NewRecorder()

		server := &PlayerServer{store:stubStore}
		server.ServeHTTP(respose, request)

		got := respose.Body.String()
		want := "20"

		assertResponseBody(t, got, want)
	})

	t.Run("returns Floyd's screo", func(t *testing.T) {
		request := NewRequestWithName(http.MethodGet, "Floyd")
		respose := httptest.NewRecorder()

		server := &PlayerServer{store:stubStore}
		server.ServeHTTP(respose, request)

		got := respose.Body.String()
		want := "10"

		assertResponseBody(t, got, want)
	})

	t.Run("returns Floyd's screo", func(t *testing.T) {
		request := NewRequestWithName(http.MethodGet, "Floyd")
		respose := httptest.NewRecorder()

		memStore := &InMemoryPlayerStore{}
		server := &PlayerServer{store:memStore}
		server.ServeHTTP(respose, request)

		got := respose.Body.String()
		want := "0"

		assertResponseBody(t, got, want)
	})

	t.Run("returns 404 on missing players", func(t *testing.T) {
		request := NewRequestWithName(http.MethodGet, "Apollo")
		response := httptest.NewRecorder()

		store := &InMemoryPlayerStore{}
		server := &PlayerServer{store:store}

		server.ServeHTTP(response,request)

		got := response.Code
		want := http.StatusNotFound

		assertStatus(t, got, want)
	})

	t.Run("it returns accepted on POST", func(t *testing.T) {
		request := NewRequestWithName(http.MethodPost, "Pepper")
		response := httptest.NewRecorder()

		store := &StubPlayerStore{}
		server := &PlayerServer{store:store}

		server.ServeHTTP(response, request)

		got := response.Code
		want := http.StatusAccepted

		assertStatus(t, got, want)
	})
}

func TestStoreWins(t *testing.T) {
	store := StubPlayerStore{
		stores:map[string]int{},
	}

	t.Run("it records wins when POST", func(t *testing.T) {
		request := NewRequestWithName(http.MethodPost, "Pepper")
		response := httptest.NewRecorder()

		server := &PlayerServer{store:&store}
		server.ServeHTTP(response, request)

		got := response.Code
		want := http.StatusAccepted

		assertStatus(t, got, want)
	})

	t.Run("it records wins on POST", func(t *testing.T) {
		playerName := "Pepper"
		request := NewRequestWithName(http.MethodPost, playerName)
		response := httptest.NewRecorder()

		server := &PlayerServer{store:&store}
		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusAccepted)
		if len(store.winCalls) != 2 {
			t.Errorf("win call %d len is not 1", len(store.winCalls))
		}

		if store.winCalls[0] != playerName {
			t.Errorf("got %s want %s", store.winCalls[0], playerName)
		}
	})
}

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	store := NewInMemoryPlayerStore()
	server := &PlayerServer{store:store}
	playerName := "cloudrain"

	server.ServeHTTP(httptest.NewRecorder(), NewRequestWithName(http.MethodPost,playerName))
	server.ServeHTTP(httptest.NewRecorder(), NewRequestWithName(http.MethodPost,playerName))
	server.ServeHTTP(httptest.NewRecorder(), NewRequestWithName(http.MethodPost,playerName))

	response := httptest.NewRecorder()
	server.ServeHTTP(response, NewRequestWithName(http.MethodGet,playerName))

	assertStatus(t, response.Code, http.StatusOK)
	assertResponseBody(t, response.Body.String(), "3")
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
