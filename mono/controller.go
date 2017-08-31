package mono

import (
    . "github.com/xferd/golib/types"
    "net/http"
)

type IController interface {
    http.Handler
    Assign(string, Any)
    Display(string)
}

type Controller struct {
    data map[string]Any
}

func (c *Controller)Assign(key string, value Any) {
    if nil == c.data {
        c.data = make(map[string]Any)
    }
    c.data[key] = value
}

func (c *Controller)Display(template string) {

}

func (c *Controller)ServeHTTP(w http.ResponseWriter, r *http.Request) {

}