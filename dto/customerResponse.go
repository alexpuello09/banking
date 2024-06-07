package dto

type CustomerResponse struct {
	Id          string `json:"customer_id"`
	Name        string `json:"full_name"`
	City        string `json:"city"`
	ZipCode     string `json:"zipcode"`
	DateOfBirth string `db:"date_of_birth"`
	Status      string `json:"status"`
}
