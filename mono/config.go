package mono

import (
    . "github.com/xferd/golib/types"
    "strings"
)

type t_conf map[string]*t_conf

var (
    conf t_conf
)

func init() {
    conf = make(t_conf)
}

func Ref(keyPath string) *t_conf {
    var ref *t_conf = &conf
    for _, k := range strings.Split(keyPath, ".") {
        if _, ok := (*ref)[k]; !ok {
            var _conf t_conf = make(t_conf)
            (*ref)[k] = &_conf
        }
        ref = (*ref)[k]
    }
    return ref
}

func Set(keyPath string, val Any) {

}