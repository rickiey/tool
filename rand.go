package tool

import (
	"crypto/rand"
	mrand "math/rand"
	"strings"
	"time"
)

// generate random numbers
func GenerateRandom() uint64 {
	b := make([]byte, 8)
	_, _ = rand.Read(b) //在byte切片中随机写入元素
	var res uint64
	for i, v := range b {
		vv := uint64(v) << (uint(i) * 8)
		res = res + vv
	}
	return res
}

//Random string lowercase
func RandString(len int) string {
	r := mrand.New(mrand.NewSource(time.Now().Unix()))
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		b := r.Intn(26) + 65
		bytes[i] = byte(b)
	}

	return strings.ToLower(string(bytes))
}
