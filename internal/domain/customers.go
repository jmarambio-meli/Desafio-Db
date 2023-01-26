package domain

type Customers struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Condition bool   `json:"condition"`
}

type TotalCustomerByCondition struct {
	Condition string
	Total     float64
}

type TotalAmountActiveCustomers struct {
	Last_name  string
	First_name string
	Amount     float64
}
