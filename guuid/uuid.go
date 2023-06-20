package guuid

import (
	u "github.com/google/uuid"
	"math/rand"
)

const codes = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6
	letterIdxMask = 1<<letterIdxBits - 1
	letterIdxMax  = 63 / letterIdxBits
)

func GenUUID() string {
	return u.NewString()
}

// RandomUUID [0,n)
func RandomUUID(n int) string {
	b := make([]byte, n)
	for i, cache, remain := n-1, rand.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = rand.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(codes) {
			b[i] = codes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return string(b)
}
