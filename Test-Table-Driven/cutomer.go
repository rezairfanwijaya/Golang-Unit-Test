package testtabledriven

import "errors"

type Customer struct {
	ID      int
	Name    string
	Address string
	Age     int
}

func (c Customer) GetAll(datas []Customer) []Customer {
	if len(datas) == 0 {
		return datas
	}
	return datas
}

func (c Customer) FindById(id int, datas []Customer) (Customer, error) {
	var cust Customer
	for _, data := range datas {
		if id == data.ID {
			cust.Address = data.Address
			cust.ID = data.ID
			cust.Name = data.Name
			cust.Age = data.Age
		}
	}

	if cust.ID == 0 {
		return cust, errors.New("customer not found")
	}

	return cust, nil
}
