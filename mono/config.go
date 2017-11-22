package mono

import (
    // . "github.com/xferd/golib/types"
    // "strings"
    "sync"
)

type t_conf map[string]*config

type config struct {
    conf *t_conf
    val interface{}
}

var (
    global_conf config
    mu_conf sync.RWMutex
)

func init() {
    global_conf = config{make(t_conf), nil}
}

func (this *config)ref(keyPath string) *t_conf {
    mu_conf.Lock()
    defer mu_conf.Unlock()

    var conf t_conf = this.conf
    for _, k := range strings.Split(keyPath, ".") {
        if _, ok := (*conf)[k]; !ok {
            (*conf)[k] = &make(t_conf)
        }

    }
    return conf
}

func Config() config {
    return global_conf
}

func (this *config)Set(keyPath string, val interface{}) {
    var ref *t_conf_map = this.ref(keyPath)
}

// func Conf(keyPath string) interface{} {
//     var ref interface{} = Ref(keyPath)
//     return ref
// }
