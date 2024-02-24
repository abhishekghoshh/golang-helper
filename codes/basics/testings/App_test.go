package testings

import (
	"fmt"
	"testing"
)

func TestReturnGeeks(t *testing.T) {
	actualString := SayHello()
	expectedString := "hello"
	if actualString != expectedString {
		t.Errorf("Expected String(%s) is not same as"+
			" actual string (%s)", expectedString, actualString)
	}
	t.Cleanup(func() { fmt.Println("This is cleanup function") })
}
