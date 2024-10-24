package xutil

import (
	rand "crypto/rand"
	rand2 "math/rand"
)

const (
	Letters          = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	LowerLetters     = "abcdefghijklmnopqrstuvwxyz"
	Digits           = "0123456789"
	LettersAndDigits = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	LowLettersDigits = "abcdefghijklmnopqrstuvwxyz0123456789"
)

func RandStr(str string, n int) string {
	if n <= 0 {
		return ""
	}
	var (
		b     = make([]rune, n)
		runes = []rune(str)
	)
	if len(runes) <= 255 {
		numberBytes := RandBytes(n)
		for i := range b {
			b[i] = runes[int(numberBytes[i])%len(runes)]
		}
	} else {
		for i := range b {
			b[i] = runes[rand2.Intn(len(runes))]
		}
	}
	return string(b)
}

func RandLetters(n int) string {
	return RandStr(Letters, n)
}

func RandLowerLetters(n int) string {
	return RandStr(LowerLetters, n)
}

func RandDigits(n int) string {
	return RandStr(Digits, n)
}

func RandLettersAndDigits(n int) string {
	return RandStr(LettersAndDigits, n)
}

func RandLowLettersDigits(n int) string {
	return RandStr(LowLettersDigits, n)
}

func RandBytes(n int) []byte {
	b := make([]byte, n)
	rand.Read(b)
	return b
}

func RandStringPointer(n int) *string {
	s := RandLetters(n)
	return &s
}
