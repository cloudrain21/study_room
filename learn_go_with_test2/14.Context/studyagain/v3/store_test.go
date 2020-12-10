package store

import (
    "context"
    "errors"
    "net/http"
    "net/http/httptest"
    "testing"
    "time"
)

type SpyStore struct {
    response string
    t        *testing.T
}

func (s *SpyStore)Fetch(ctx context.Context) (string,error) {

    dataChan := make(chan string, 1)
    defer close(dataChan)

    go func() {
        result := ""

        for _, c := range s.response {
            select {
            case <- ctx.Done():
                s.t.Log("canceled in goroutine...")
                return
            default:
                time.Sleep(10 * time.Millisecond)
                result += string(c)
            }
        }

        dataChan <- result
    }()

    select {
    case <- ctx.Done():
        s.t.Log("canceled...")
        return "", ctx.Err()
    case result := <- dataChan:
        return result, nil
    }
}

type SpyResponseWriter struct {
    written bool
}

func (s *SpyResponseWriter)Header() http.Header {
    s.written = true
    return nil
}

func (s *SpyResponseWriter)WriteHeader(statusCode int) {
    s.written = true
}

func (s *SpyResponseWriter)Write([]byte) (int,error) {
    s.written = true
    return 0, errors.New("not implemented")
}

func TestHandler(t *testing.T) {

    t.Run("spy basic", func(t *testing.T) {
        spyStore := &SpyStore{"myresponse", t}

        svr := Server(spyStore)

        req := httptest.NewRequest(http.MethodGet, "/", nil)
        res := httptest.NewRecorder()

        svr.ServeHTTP(res, req)

        if res.Body.String() != spyStore.response {
            t.Errorf("got %s want %s", res.Body.String(), spyStore.response)
        }
    })

    t.Run("spy request with cancel context", func(t *testing.T) {
        spyStore := &SpyStore{"myresponse", t}

        svr := Server(spyStore)

        req := httptest.NewRequest(http.MethodGet, "/", nil)

        cancelingCtx, cancel := context.WithCancel(req.Context())
        time.AfterFunc(5 * time.Millisecond, cancel)
        req = req.WithContext(cancelingCtx)

        res := httptest.NewRecorder()

        svr.ServeHTTP(res, req)

        if res.Body.String() != "" {
            t.Error("must be canceled, but not")
        }
    })

    t.Run("spy response writer", func(t *testing.T) {
        data := "myresponse"
        spyStore := &SpyStore{data, t}

        req := httptest.NewRequest(http.MethodGet, "/", nil)
        res := &SpyResponseWriter{}

        svr := Server(spyStore)
        svr.ServeHTTP(res, req)

        if res.written != true {
            t.Error("spy response was not written")
        }
    })

    t.Run("spy response writer timeout", func(t *testing.T) {
        data := "myresponse"
        spyStore := &SpyStore{data, t}

        req := httptest.NewRequest(http.MethodGet, "/", nil)
        cancelingCtx, cancel := context.WithCancel(req.Context())
        time.AfterFunc(5 * time.Millisecond, cancel)
        req = req.WithContext(cancelingCtx)

        res := &SpyResponseWriter{}

        svr := Server(spyStore)
        svr.ServeHTTP(res, req)

        if res.written {
            t.Error("spy response must not be written")
        }
    })
}
