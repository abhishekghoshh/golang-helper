package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	str := "sha256 this string"
	hsh := sha256.New()
	hsh.Write([]byte(str))

	bs := hsh.Sum(nil)

	fmt.Println(str)
	fmt.Printf("%x\n", bs)
}
