package types

import (
    "math/rand"
)

type Slice []Any

func (s Slice)Random() Any {
    return s[rand.Intn(len(s))]
}
