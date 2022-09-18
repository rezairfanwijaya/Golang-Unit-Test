package testtabledriven_test

import (
	"fmt"
	"testing"

	c "github.com/rezairfanwijaya/Golang-Unit-Test/Test-Table-Driven"
	"github.com/stretchr/testify/assert"
)

func TestGetAllCustomer(t *testing.T) {
	var customer c.Customer
	datas := []c.Customer{
		{
			ID:      1,
			Name:    "Cust 1",
			Address: "Jakarta",
			Age:     20,
		}, {
			ID:      2,
			Name:    "Cust 2",
			Address: "Surabaya",
			Age:     17,
		}, {
			ID:      3,
			Name:    "Cust 3",
			Address: "Bandung",
			Age:     23,
		},
	}

	testCases := []struct {
		Name     string
		Customer []c.Customer
		Want     []c.Customer
	}{
		{
			Name:     "Success",
			Customer: datas,
			Want:     datas,
		}, {
			Name:     "Success with no data",
			Customer: []c.Customer{},
			Want:     []c.Customer{},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			actual := customer.GetAll(testCase.Customer)
			assert.Equal(t, testCase.Want, actual)
		})
	}

}

func TestFindById(t *testing.T) {
	var customer c.Customer
	var ERRO_NOT_FOUND error = fmt.Errorf("customer not found")
	datas := []c.Customer{
		{
			ID:      1,
			Name:    "Cust 1",
			Address: "Jakarta",
			Age:     20,
		}, {
			ID:      2,
			Name:    "Cust 2",
			Address: "Surabaya",
			Age:     17,
		}, {
			ID:      3,
			Name:    "Cust 3",
			Address: "Bandung",
			Age:     23,
		},
	}

	testCases := []struct {
		Name      string
		Customer  []c.Customer
		Id        int
		Want      interface{}
		WantError bool
	}{
		{
			Name:      "Success",
			Customer:  datas,
			Id:        2,
			Want:      c.Customer{ID: 2, Name: "Cust 2", Address: "Surabaya", Age: 17},
			WantError: false,
		}, {
			Name:      "Not Found",
			Customer:  datas,
			Id:        90,
			Want:      ERRO_NOT_FOUND,
			WantError: true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			actual, err := customer.FindById(testCase.Id, testCase.Customer)
			if testCase.WantError {
				assert.Equal(t, testCase.Want, err)
			} else {
				assert.Equal(t, testCase.Want, actual)
			}
		})
	}
}
