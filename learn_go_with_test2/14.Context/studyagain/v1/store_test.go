package store

import (
    "net/http"
    "net/http/httptest"
    "testing"
)
type StubStore struct {
    response string
}

func (s *StubStore)Fetch() string {
    return s.response
}

func TestHandler(t *testing.T) {
    t.Run("stub store", func(t *testing.T) {

        data := "mydata"
        stubStore := &StubStore{data}

        req := httptest.NewRequest(http.MethodGet, "/", nil)
        res := httptest.NewRecorder()

        svr := Server(stubStore)
        svr.ServeHTTP(res, req)

        got := res.Body.String()

        if got != data {
            t.Errorf("got %s want %s", got, data)
        }
    })
}
