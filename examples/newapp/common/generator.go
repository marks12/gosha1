package common

import (
    "math/rand"
    "time"
)

func RandomString(size int) string {
    var alpha = "abcdefghijkmnpqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ23456789"
    rand.Seed(time.Now().UTC().UnixNano())
    buf := make([]byte, size)
    for i := 0; i < size; i++ {
        buf[i] = alpha[rand.Intn(len(alpha))]
    }
    return string(buf)
}
