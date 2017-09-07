package mono

import (
    // _ "github.com/xferd/golib/types"
    "net/http"
    "log"
    "sync"
    "regexp"
    "fmt"
)


type monoServeMux struct {
    mu sync.RWMutex
    m map[string]IController
}

var (
    serveMux monoServeMux = monoServeMux{}
)

func getServeMux() *monoServeMux {
    return &serveMux
}

func Handle(pattern string, controller IController) {
    serveMux.handle(pattern, controller)
}

func (mux *monoServeMux)controller(w http.ResponseWriter, r *http.Request) (c IController, pattern string) {
    if nil == mux.m {
        return nil, ""
    }

    for p, c := range mux.m {
        if ok, _ := regexp.MatchString(p, r.RequestURI); ok {
            return c, p
        }
    }

    panic(404)
    return nil, ""
}

func (mux *monoServeMux)handle(pattern string, controller IController) (e error) {
    e = nil
    mux.mu.Lock()
    defer mux.mu.Unlock()

    if pattern == "" {
        panic("http: invalid pattern " + pattern)
    }
    if controller == nil {
        panic("http: nil controller")
    }

    if mux.m == nil {
        mux.m = make(map[string]IController)
    }
    mux.m[pattern] = controller
    return
}

func (mux *monoServeMux)ServeHTTP(w http.ResponseWriter, r *http.Request) {
    defer func() {
        err := recover()
        if nil == err {
            return
        }

        if v, ok := err.(int); ok && 404 == v {
            fmt.Fprint(w, 404)
        }
    }()

    log.Println("uri:", r.RequestURI)
    if r.RequestURI == "*" {
        if r.ProtoAtLeast(1, 1) {
            w.Header().Set("Connection", "close")
        }
        w.WriteHeader(http.StatusBadRequest)
        return
    }
    c, _ := mux.controller(w, r)
    c.ServeHTTP(w, r)
}
