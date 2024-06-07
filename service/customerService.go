package service

import (
	"banking/domain"
	"banking/dto"
	"banking/errs"
	"strings"
)

type CustomerService interface {
	GetAllCustomers(string) ([]dto.CustomerResponse, *errs.AppError)
	GetCustomer(string) (*dto.CustomerResponse, *errs.AppError)
}

type DefaultCustomerService struct {
	Repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomers(status string) ([]dto.CustomerResponse, *errs.AppError) {
	if strings.ToLower(status) == "active" {
		status = "1"
	} else if strings.ToLower(status) == "inactive" {
		status = "0"
	} else {
		status = ""
	}

	var customerSlideResult []dto.CustomerResponse
	customerSlide, err := s.Repo.FindAll(status)
	if err != nil {
		return nil, err
	}

	for _, customer := range customerSlide {
		customerSlideResult = append(customerSlideResult, customer.ToDto())
	}
	return customerSlideResult, nil
}

func (s DefaultCustomerService) GetCustomer(id string) (*dto.CustomerResponse, *errs.AppError) {
	c, err := s.Repo.ById(id)
	if err != nil {
		return nil, err
	}

	response := c.ToDto()
	return &response, nil
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}
