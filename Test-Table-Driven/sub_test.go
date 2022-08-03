package sub_test

import (
	"testing"

	. "github.com/rezairfanwijaya/Golang-Unit-Test/Test-Table-Driven"
)

// bikin test function
func TestSub(t *testing.T) {

	// create test table
	// test table ini merupakan slice of struct yang berisi value dan expectation
	testTable := []struct {
		a           int
		b           int
		expectation int
	}{
		{a: 3, b: 1, expectation: 2},
		{a: 0, b: 0, expectation: 0},
		{a: 1, b: 8, expectation: -7},
		{a: -3, b: 1, expectation: -4},
		{a: 9, b: -1, expectation: 10},
		{a: -3, b: -1, expectation: -2},
	}

	// looping tesTable
	for _, value := range testTable {
		// panggil function yang akan ditest dan assign ke variable. Dan isi params menggunakan value dari testTable
		result := Sub(value.a, value.b)

		// assertion
		if result != value.expectation {
			t.Logf("Got: %d But Expect %d", result, value.expectation)
			t.Fail()
		}
	}
}
