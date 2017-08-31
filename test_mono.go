package main

import (
    "github.com/xferd/golib/mono"
    "net/http"
    "fmt"
)

type HomeController struct {
    mono.Controller
}

func (c *HomeController)ServeHTTP(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "homepage")
}

func main() {
    mono.Handle("^\\/abc\\/\\d+$", &HomeController{})
    mono.ListenAndServe(":8081")
}