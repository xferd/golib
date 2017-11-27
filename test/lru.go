package main

import (
)

type Node struct {
    Val int
    Key string
    prev *Node
    next *Node
}

type CacheList struct {
    head *Node
    tail *Node
}

func (c *CacheList)push(key string, val int) {
    node := Node{Val: val, Key: key, prev: nil, next: c.head}
    c.head = &node
}

type LRU struct {
    cache *CacheList
}

func NewLRU() *LRU {
    return &LRU{CacheList{nil, nil}}
}

func main() {

}
