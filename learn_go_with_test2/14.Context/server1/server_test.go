package server

import (
    "context"
    "net/http"
    "net/http/httptest"
    "testing"
    "time"
)

type StubStore struct {
    response string
}

func (s *StubStore)Fetch() string {
    return s.response
}

func (s *StubStore)Cancel() {
}

type SpyStore struct {
    response string
    cancelled bool
}

func (s *SpyStore)Fetch() string {
    time.Sleep(100 * time.Millisecond)
    return s.response
}

func (s *SpyStore)Cancel() {
    s.cancelled = true
}

func TestHandler(t *testing.T) {
    t.Run("server1", func(t *testing.T) {
        data := "hello, world"
        svr := Server(&StubStore{data})

        request := httptest.NewRequest(http.MethodGet, "/", nil)
        response := httptest.NewRecorder()

        svr.ServeHTTP(response, request)

        if response.Body.String() != data {
            t.Errorf("got %s want %s", response.Body.String(), data)
        }
    })

    t.Run("cancelling context", func(t *testing.T) {
        store := &SpyStore{"hello world", false}
        svr := Server(store)

        request := httptest.NewRequest(http.MethodGet, "/", nil)

        cancellingCtx, cancel := context.WithCancel(request.Context())

        // 5 millisecond 후에 request context 의 cancel 함수를 호출함
        time.AfterFunc(5 * time.Millisecond, cancel)

        // request 는 done channel 에서 기다리게 됨
        request = request.WithContext(cancellingCtx)
        response := httptest.NewRecorder()

        svr.ServeHTTP(response, request)

        if !store.cancelled {
            t.Errorf("store was not cancelled")
        }
    })

    t.Run("returns data from store", func(t *testing.T) {
        store := &SpyStore{response: "hello, world"}
        svr := Server(store)

        request := httptest.NewRequest(http.MethodGet, "/", nil)
        response := httptest.NewRecorder()

        svr.ServeHTTP(response, request)

        if response.Body.String() != "hello, world" {
            t.Errorf("got (%s) want (%s)", response.Body.String(), "hello, world")
        }

        if store.cancelled {
            t.Error("must not be cancelled")
        }
    })
}
