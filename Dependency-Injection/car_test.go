package dependencyinjection_test

import (
	"testing"

	DI "github.com/rezairfanwijaya/Golang-Unit-Test/Dependency-Injection"

	"github.com/stretchr/testify/assert"
)

func TestDefaultEngine_MaxSpeed(t *testing.T) {
	testCases := []struct {
		Name        string
		Expectation int
	}{
		{
			Name:        "Default enginge should have 50 max speed",
			Expectation: 50,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			de := DI.DefaultEngine{}
			assert.Equal(t, testCase.Expectation, de.MaxSpeed())
		})
	}
}

func TestTurboEngine_MaxSpeed(t *testing.T) {
	testCases := []struct {
		Name        string
		Expectation int
	}{
		{
			Name:        "Turbo enginge should have 100 max speed",
			Expectation: 100,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			te := DI.TurboEngine{}
			assert.Equal(t, testCase.Expectation, te.MaxSpeed())
		})
	}
}

func TestCar_Speed(t *testing.T) {
	type SpeederType struct {
		Speeder DI.Speeder
	}

	testCases := []struct {
		Name     string
		Speeder  SpeederType
		Expected int
	}{
		{
			Name: "speed should be 50 when use default engine",
			Speeder: SpeederType{
				Speeder: &DI.DefaultEngine{},
			},
			Expected: 50,
		},
		{
			Name: "speed should be 90 when use turbo engine",
			Speeder: SpeederType{
				Speeder: &DI.TurboEngine{},
			},
			Expected: 9 0,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			newCar := DI.Car{
				Speeder: testCase.Speeder.Speeder,
			}
			actual := newCar.Speed()
			assert.Equal(t, testCase.Expected, actual)
		})
	}
}
