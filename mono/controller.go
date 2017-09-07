package mono

import (
    . "github.com/xferd/golib/types"
    "net/http"
    "html/template"
    // "os"
    "log"
    "path/filepath"
)

type IController interface {
    http.Handler
    Assign(string, Any)
    Display(string, http.ResponseWriter)
    ViewPath(tplFile string) string
}

type Controller struct {
    data map[string]Any
}

var (
    viewPath string
)

func SetViewPath(path string) {
    viewPath = path
}

func (c *Controller)Assign(key string, value Any) {
    if nil == c.data {
        c.data = make(map[string]Any)
    }
    c.data[key] = value
}

func (c *Controller)ViewPath(tplFile string) string {
    return viewPath + tplFile
}

func (c *Controller)Display(tplFile string, w http.ResponseWriter) {
    tplname := filepath.Base(tplFile)
    tmpl, err := template.New(tplname).ParseFiles(c.ViewPath(tplFile))
    if err != nil {
        log.Println(err)
    }
    if e := tmpl.Execute(w, c.data); e != nil {
        log.Println(e)
    }
}

func (c *Controller)WriteString(s string) {
    // os.Stdout
}

func (c *Controller)ServeHTTP(w http.ResponseWriter, r *http.Request) {

}