package main

import (
	"errors"
	"fmt"
)

type User struct {
	Id      int
	Name    string
	Age     int
	Address string
}

func (u *User) FindById(users []User, id int) (User, error) {
	var user User
	for _, value := range users {
		if value.Id == id {
			user.Id = value.Id
			user.Address = value.Address
			user.Age = value.Age
			user.Name = value.Name
		}
	}

	if user.Id == 0 {
		return user, errors.New("not found")
	}

	return user, nil
}

func main() {
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

	var user User
	fmt.Println(user.FindById(datas, 1))
}
