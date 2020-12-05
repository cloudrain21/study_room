package server

import (
    "context"
    "net/http"
    "net/http/httptest"
    "testing"
)

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
        data := "hello world"
        svr := Server(&StubStore{data})

        request := httptest.NewRequest(http.MethodGet, "/", nil)

        cancellingCtx, cancel := context.WithCancel(request.Context())

        response := httptest.NewRecorder()
    })
}
