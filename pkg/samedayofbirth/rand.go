package samedayofbirth

import (
    "math/rand"
    "time"
)

func getRandom() int {
    min := 1
    max := 365
    rand.Seed(time.Now().UnixNano())
    return rand.Intn(max - min) + min
}