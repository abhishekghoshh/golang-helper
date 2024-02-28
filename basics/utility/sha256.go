package utility

import (
	"crypto/sha256"
	"fmt"
)

/*
https://gobyexample.com/sha256-hashes
https://en.wikipedia.org/wiki/SHA-2
https://en.wikipedia.org/wiki/Cryptographic_hash_function

SHA256 hashes are frequently used to compute short identities for binary or text blobs.
For example, TLS/SSL certificates use SHA256 to compute a certificate's signature. Here's how to compute SHA256 hashes in Go.

Go implements several hash functions in various crypto/* packages.
*/
func SHA256Hashing() {

	str := "sha256 this string"

	// Here we start with a new hash.
	hsh := sha256.New()

	// Write expects bytes. If you have a string s, use []byte(s) to coerce it to bytes.
	hsh.Write([]byte(str))

	// This gets the finalized hash result as a byte slice.
	// The argument to Sum can be used to append to an existing byte slice: it usually isn't needed.
	bs := hsh.Sum(nil)

	fmt.Println(str)
	fmt.Printf("%x\n", bs)

	// Running the program computes the hash and prints it in a human-readable hex format.
	// You can compute other hashes using a similar pattern to the one shown above.
	// For example, to compute SHA512 hashes import crypto/sha512 and use sha512.New().
	// Note that if you need cryptographically secure hashes, you should carefully research hash strength!
}
