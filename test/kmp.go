package main

import (
    "fmt"
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
    return append([]int{-1}, k...)[:len(p)]
}

func kmp(s, p string) int {
    k := getk(p)
    fmt.Println(k)
    for i := 0; i < len(s) - len(p); i++ {
        found := true
        fmt.Println(string(s[i:]))
        for j, cp := range p {
            if cp != rune(s[i + j]) {
                found = false
                fmt.Println(i + j, string(s[i + j]), j, string(cp), k[j])
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
    s := "acbabcabaaacdf"
    p := "abcaba"
    pos := kmp(s, p)
    fmt.Println(s, p, pos)
}