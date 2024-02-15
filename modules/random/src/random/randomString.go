package random

import (
	cryptrand "crypto/rand"
)

const source = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

type Rand struct {
	len int
}

func (rnd *Rand) Generate() string {
	str, r := make([]rune, rnd.len), []rune(source)
	for i := range str {
		p, _ := cryptrand.Prime(cryptrand.Reader, rnd.len)
		x, y := p.Uint64(), uint64(rnd.len)
		str[i] = r[x%y]
	}
	return string(str)
}

func New(len int) *Rand {
	return &Rand{
		len: len,
	}
}
