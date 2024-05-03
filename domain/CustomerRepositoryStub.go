package domain

type CustomerRepository interface {
	FindAll() ([]Customer, error)
}

type CustomerRepositoryStub struct {
	customers []Customer
}

func (C CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return C.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{"10001", "Ashli", "New Delhi", "100100", "09-03-2005", "1"},
		{"738490", "Ashli", "New Delhi", "100100", "02-03-2010", "7"},
	}
	return CustomerRepositoryStub{customers}
}
