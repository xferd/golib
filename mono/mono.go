package mono

import (
    "net/http"
    "fmt"
)

func ListenAndServe(addr string) {
    fmt.Println("start")
    if err := http.ListenAndServe(addr, getServeMux()); err != nil {
        fmt.Println("ListenAndServe error: ", err.Error())
    }
}
