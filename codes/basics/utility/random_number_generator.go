package utility

import (
	cryptorand "crypto/rand"
	"fmt"
	mathrand "math/rand"
	"time"
)

// https://gobyexample.com/random-numbers
func RandomNumberGenerator() {
	// Intn returns, as an
	// int, a non-negative
	// pseudo-random number in
	// [0,n) from the default Source.
	// i.e. simply call Intn to
	// get the next random integer.
	fmt.Println(mathrand.Intn(256))
	// Float64 returns, as
	// a float64, a pseudo-random
	// number in [0.0,1.0)
	// from the default Source.
	fmt.Println(mathrand.Float64() * 256)

	// Implementation is slow
	// to make it faster
	// rand.Seed(time.Now().UnixNano())
	// is added. Seed is the current time,
	// converted to int64 by UnixNano.
	// Gives constantly changing numbers
	source := mathrand.NewSource(time.Now().UnixNano())
	rand := mathrand.New(source)
	fmt.Println(rand.Intn(256))

	// The above method is not safe if the user
	// wants to keep the random numbers secret.
	// That’s why Golang provides Crypto rand to
	// varies the level of randomness of numbers to come.
	// It is crypto-ready to use and secure but it is slower.
	// It is used for generating passkeys, CSRF tokens or anything related to security.
	RandomCrypto, err := cryptorand.Prime(cryptorand.Reader, 10)
	fmt.Println(RandomCrypto, err)
}
