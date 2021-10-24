package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{Id: "001", Name: "hisam", City: "bogor", Zipcode: "17021", DateOfBirth: "2000-01-01", Status: "1"},
		{Id: "002", Name: "maulana", City: "bogor", Zipcode: "17021", DateOfBirth: "2000-01-01", Status: "1"},
	}
	return CustomerRepositoryStub{customers: customers}
}
