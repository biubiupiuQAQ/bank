package util

import (
	"math/rand"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// 生成min到max间的随机数
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

const alphat = "abcdefghijklmnopqrstuvwxyz"

// 生成随机的字符串
func RandomString(n int) string {
	l := len(alphat)
	var str strings.Builder
	for i := 0; i < n; i++ {
		c := alphat[rand.Intn(l)]
		str.WriteByte(c)
	}
	return str.String()
}

// 生成随机姓名
func RandomName() string {
	return RandomString(5)
}

// 生成随机金钱
func RandomMoney() int64 {
	return RandomInt(0, 10000)
}

// 生成随机货币结算方式
func RandomCurrency() string {
	currencies := []string{"USD", "CHY", "EUR", "JPY", "GBP", "KRW"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}
