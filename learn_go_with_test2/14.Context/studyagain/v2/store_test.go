package store

import (
    "context"
    "net/http"
    "net/http/httptest"
    "testing"
    "time"
)

type SpyStore struct {
    response string
    canceled bool
}

func (s *SpyStore)Fetch() string {
    time.Sleep(100 * time.Millisecond)
    return s.response
}

func (s *SpyStore)Cancel() {
    s.canceled = true
}

func TestHandler(t *testing.T) {
    t.Run("spy normal test", func(t *testing.T) {

        data := "mydata"
        spyStore := &SpyStore{data, false}

        req := httptest.NewRequest(http.MethodGet, "/", nil)
        res := httptest.NewRecorder()

        svr := Server(spyStore)
        svr.ServeHTTP(res, req)

        got := res.Body.String()

        if spyStore.canceled {
            t.Error("must not be canceled")
        }

        if got != data {
            t.Errorf("got %s want %s", got, data)
        }
    })

    t.Run("spy canceling request", func(t *testing.T) {

        data := "mydata"
        spyStore := &SpyStore{data, false}

        req := httptest.NewRequest(http.MethodGet, "/", nil)

        cancelingCtx, cancel := context.WithCancel(req.Context())
        time.AfterFunc(5 * time.Millisecond, cancel)
        req = req.WithContext(cancelingCtx)

        res := httptest.NewRecorder()

        svr := Server(spyStore)
        svr.ServeHTTP(res, req)

        if ! spyStore.canceled {
            t.Error("must be canceled, but not")
        }
    })
}
