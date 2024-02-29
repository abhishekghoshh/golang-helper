package test

import (
	"basics/testings/src"
	"fmt"
	"testing"
)

/*
https://gobyexample.com/testing-and-benchmarking

*/

type MinTestCase struct {
	a, b int
	want int
}

func minTestCases() []MinTestCase {
	return []MinTestCase{
		{0, 1, 0},
		{1, 0, 0},
		{2, -2, -2},
		{0, -1, -1},
		{-1, 0, -1},
	}
}

func cleanUp() {
	fmt.Println("This is cleanup function")
}

func TestReturnGeeks(t *testing.T) {

	actualString := src.SayHello()

	expectedString := "hello"

	if actualString != expectedString {
		t.Errorf("Expected String(%s) is not same as"+
			" actual string (%s)", expectedString, actualString)
	}

	t.Cleanup(cleanUp)
}

func TestMinBasic(t *testing.T) {
	ans := src.Min(2, -2)

	if ans != -2 {
		t.Errorf("IntMin(2, -2) = %d; want -2", ans)
	}
	t.Cleanup(cleanUp)
}

func TestMinTableDriven(t *testing.T) {
	for _, tt := range minTestCases() {

		testname := fmt.Sprintf("%d,%d", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			ans := src.Min(tt.a, tt.b)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}

func BenchmarkMin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src.Min(1, 2)
	}
}
