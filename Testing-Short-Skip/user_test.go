package main

import (
	"testing"
	"time"
)

func TestUser(t *testing.T) {
	var user User
	datas := []User{
		{
			Id:      1,
			Name:    "User 1",
			Age:     20,
			Address: "Jakarta",
		},
		{
			Id:      2,
			Name:    "User 2",
			Age:     40,
			Address: "Surabaya",
		},
		{
			Id:      3,
			Name:    "User 3",
			Age:     17,
			Address: "Bandung",
		},
	}

	type args struct {
		name       string
		data       []User
		id         int
		expectedId int
	}

	testCases := []args{
		{
			name:       "Success With Id 2",
			data:       datas,
			id:         2,
			expectedId: 2,
		},
		{
			name: "Will Be skipped",
		},
		{
			name:       "Success With Id 3",
			data:       datas,
			id:         3,
			expectedId: 3,
		},
		{
			name:       "Id Not Found",
			data:       datas,
			id:         0,
			expectedId: 0,
		},
	}

	for _, testCase := range testCases {
		if testCase.name != "Will Be skipped" {
			t.Run(testCase.name, func(t *testing.T) {
				actual, _ := user.FindById(testCase.data, testCase.id)
				if actual.Id != testCase.expectedId {
					t.Logf("Found : %v But Expected : %v", actual.Id, testCase.expectedId)
				}
			})
		} else {
			t.Skip(testCase.name) // skiped testing

		}

	}

}

func TestCallToDB(t *testing.T) {
	if testing.Short() {
		t.Skip("Will be skip, because calling db is way to long")
	}

	<- time.After(3 *time.Second)

}
