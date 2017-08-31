package mono

import (
    . "github.com/xferd/golib/types"
    "net/http"
    "html/template"
    // "os"
    "log"
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

func (c *Controller)Assign(key string, value Any) {
    if nil == c.data {
        c.data = make(map[string]Any)
    }
    c.data[key] = value
}

func (c *Controller)ViewPath(tplFile string) string {
    return ""
    // dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
}

func (c *Controller)Display(tplFile string, w http.ResponseWriter) {
    tmpl, err := template.New("index.tpl").ParseFiles(tplFile)
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