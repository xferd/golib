package main

import (
    "fmt"
    "strings"
)

func getk(p string) (k []int) {
    k = make([]int, len(p))
    k[0] = 0
    for i := 1; i < len(p); i++ {
        k[i] = k[i - 1]
        if p[k[i - 1]] == p[i] {
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
    fmt.Println("pattern: ", p)
    fmt.Println("k:", k)
    st := 0
    for i, j := 0, 0; i < len(s) - len(p); {
        st++
        fmt.Printf("step %d -> i: %d, j: %d\n", st, i, j)
        fmt.Println(s)
        fmt.Println(strings.Repeat(" ", i + j) + "^")
        fmt.Println(strings.Repeat(" ", i) + p)
        fmt.Println(strings.Repeat("=", len(s)))
        if s[i + j] == p[j] {
            j++;
        } else {
            i = i + j - k[j]
            if j = k[j]; j < 0 {
                j = 0
            }
            continue
        }
        if j >= len(p) {
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