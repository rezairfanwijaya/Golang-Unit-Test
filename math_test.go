package math

import "testing"

// bikin function untuk testing
// biasanya memiliki prefik Test dan parameter t *testing.T

func TestAdd(t *testing.T) {
	// panggil function yang mau di test dan tampung dalam varibale
	result := add(1, 4)

	// bikin assertion atau sebuah tuntuan untuk diterapkan pada function yang akan ditest
	if result != 5 {
		t.Logf("Got: %d but expect: %d", result, 5)
		t.Fail()
	}

}
