package multiplication_test

import (
	"testing"
	. "github.com/rezairfanwijaya/Golang-Unit-Test/Heirarchical-Test-With-Subtest"
)

// pada hierarchical test with subtest ini kita bisa melakukan nested test, jadi di dalam satu function test bisa memiliki sub test lain dan sub test bisa memiliki sub test juga
// struktur
/*
	TestMain
		SubTest1
			SubTest_Subtest1
			SubTest_Subtest1
			SubTest_Subtest1
		SubTest2
			SubTest_Subtest2
			SubTest_Subtest2
			SubTest_Subtest2
*/

// pada test kali ini kita akan mencoba test function perkalian
// dengan rules
// + x + = +
// + x - = -
// - x - = +


func TestMultiplication(t *testing.T) {
	// subtest untuk a positive
	t.Run("A = Positive", func(t *testing.T){
		// deklarasi untuk a = positive
		a := 10

		// bikin subtest untuk jika b positive
		t.Run("B = Positive", func(t *testing.T){
			// panggil function multiplication dan assign ke variable
			result := Multiplication(a, 10)

			// assertion
			if result != 100 {
				t.Logf("Got %d But Expect %d", result, 100)
				t.Fail()
			}
		})

		// bikin subtest untuk jika b negative
		t.Run("B = Negative", func(t *testing.T){
			// panggil function multiplication dan assign ke variable
			result := Multiplication(a, -10)

			// assertion
			if result != -100 {
				t.Logf("Got %d But Expect %d", result, -100)
				t.Fail()
			}
		})

		// bikin subtest untuk jika b 0
		t.Run("B = Nol", func(t *testing.T){
			// panggil function multiplication dan assign ke variable
			result := Multiplication(a, 0)

			// assertion
			if result != 0 {
				t.Logf("Got %d But Expect %d", result, 0)
				t.Fail()
			}
		})
	})

	// subtest untuk A negative
	t.Run("A = Negative", func(t *testing.T){
		a := -10

		// bikin subtest jika b negative
		t.Run("B = Negative", func(t *testing.T){
			// panggil function multiplication dan assign ke variable
			result := Multiplication(a, -10)

			// assertion
			if result != 100 {
				t.Logf("Got %d But Expect %d", result, 100)
				t.Fail()
			}
		})

		// bikin subtest jika b nol
		t.Run("B = Nol", func(t *testing.T){
			// panggil function multiplication dan assign ke variable
			result := Multiplication(a, 0)

			// assertion
			if result != 0 {
				t.Logf("Got %d But Expect %d", result, 0)
				t.Fail()
			}
		})
	})

}
