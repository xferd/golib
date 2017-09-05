package main

import (
    "github.com/xferd/golib/mono"
    "net/http"
    "log"
)

type HomeController struct {
    mono.Controller
}

func (c *HomeController)ServeHTTP(w http.ResponseWriter, r *http.Request) {
    log.Println("homepage")
    c.Assign("name", "xferd")
    c.Assign("company", "Lenovo")
    c.Display("/Users/leeyan/go/src/github.com/xferd/golib/template/index.tpl", w)
}

func main() {
    mono.Handle("^\\/abc\\/\\d+$", &HomeController{})
    mono.ListenAndServe(":8081")
}