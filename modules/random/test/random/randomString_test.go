package random

import (
	"fmt"
	"testing"

	"github.com/abhishekghoshh/golang-module/src/random"
)

func RandomStringGenerate(t *testing.T) {
	l := 100
	rand := random.New(l)
	if len(rand.Generate()) != l {
		fmt.Errorf("length is not same")
	}
	str1 := rand.Generate()
	str2 := rand.Generate()
	if str1 == str2 {
		fmt.Errorf("random is not random")
	}
}
