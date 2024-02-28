package utility

import (
	cryptorand "crypto/rand"
	"fmt"
	rand "math/rand/v2"
)

/*
https://gobyexample.com/random-numbers
https://en.wikipedia.org/wiki/Pseudorandom_number_generator
https://en.wikipedia.org/wiki/Permuted_congruential_generator

Go's math/rand/v2 package provides pseudorandom number generation.
*/

func RandomNumberGenerator() {

	// For example, rand.IntN returns a random int n, 0 <= n < 100.
	// Intn returns, as an int, a non-negative pseudo-random number in
	// [0,n) from the default Source. i.e. simply call Intn to get the next random integer.
	fmt.Print(rand.IntN(100), ",")
	fmt.Print(rand.IntN(100))
	fmt.Println()

	// rand.Float64 returns a float64 f, 0.0 <= f < 1.0.
	// Float64 returns, as a float64, a pseudo-random number in [0.0,1.0) from the default Source.
	fmt.Print((rand.Float64()*5)+5, ",")
	fmt.Print((rand.Float64() * 5) + 5)
	fmt.Println()

	// If you want a known seed, create a new rand.Source and pass it into the New constructor.
	// NewPCG creates a new PCG source that requires a seed of two uint64 numbers.
	s2 := rand.NewPCG(42, 1024)
	r2 := rand.New(s2)
	fmt.Print(r2.IntN(100), ",")
	fmt.Print(r2.IntN(100))
	fmt.Println()

	s3 := rand.NewPCG(42, 1024)
	r3 := rand.New(s3)
	fmt.Print(r3.IntN(100), ",")
	fmt.Print(r3.IntN(100))
	fmt.Println()

	// The above method is not safe if the user wants to keep the random numbers secret.
	// That's why Golang provides Crypto rand to varies the level of randomness of numbers to come.
	// It is crypto-ready to use and secure but it is slower.
	// It is used for generating passkeys, CSRF tokens or anything related to security.
	RandomCrypto, err := cryptorand.Prime(cryptorand.Reader, 10)
	fmt.Println(RandomCrypto, err)
}
