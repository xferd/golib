package main

import (
    "fmt"
    "strings"
)

func getk(p string) (k []int) {
    k = make([]int, len(p))
    k[0] = 0
    for i, c := range p {
        if i == 0 {
            continue
        }
        k[i] = k[i - 1]
        if p[k[i - 1]] == uint8(c) {
            k[i]++
        } else {
            k[i] = 0
        }
    }
    k = append([]int{-1}, k[:len(p) - 1]...)
    return k
}

func kmp(s, p string) int {
    k := getk(p)
    fmt.Println("k:", k)
    for i := 0; i < len(s) - len(p); {
        found := true
        for j, c := range p {
            fmt.Println("i:", i, "j:", j)
            fmt.Println(s)
            fmt.Println(strings.Repeat(" ", i + j) + "^")
            fmt.Println(strings.Repeat(" ", i) + p)
            fmt.Println(strings.Repeat("=", len(s)))
            if c != rune(s[i + j]) {
                found = false
                i = i + j - k[j]
                break
            }
        }
        if found {
            return i
        }
    }
    return -1
}

func main() {
    s := "BBC ABCDAB ABCDABCDABDE"
    p := "ABCDABD"
    pos := kmp(s, p)
    fmt.Println("string:", s, "\npattern:", p, "\npos:", pos)
}