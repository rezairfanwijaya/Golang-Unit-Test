// package test harus memiliki sufiks _test agar ketika file di build, file test ini tidak ikut terbuild menjadi executable file
package math_test

import (
	"testing"
	. "github.com/rezairfanwijaya/Golang-Unit-Test"
)

// bikin function untuk testing
// biasanya memiliki prefiks Test dan parameter t *testing.T

func TestAdd(t *testing.T) {
	// panggil function yang mau di test dan tampung dalam varibale
	result := Add(1, 4)

	// bikin assertion atau sebuah tuntuan untuk diterapkan pada function yang akan ditest
	if result != 5 {
		t.Logf("Got: %d but expect: %d", result, 5)
		t.Fail()
	}

}
