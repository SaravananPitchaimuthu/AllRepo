package service

import (
	"github.com/SaravananPitchaimuthu/pub-api/domain"
	"github.com/SaravananPitchaimuthu/pub-api/dto"
	"github.com/SaravananPitchaimuthu/pub-api/errs"
)

type CustomerService interface {
	GetAllCustomer(string) ([]dto.CustomerResponse, *errs.AppError)
	GetCustomer(string) (*dto.CustomerResponse, *errs.AppError)
}

type defaultCustomerService struct {
	repo domain.CustomerRepository
}

func NewCustomerService(repo domain.CustomerRepository) defaultCustomerService {
	return defaultCustomerService{repo: repo}
}

func (s defaultCustomerService) GetAllCustomer(status string) ([]dto.CustomerResponse, *errs.AppError) {
	if status == "active" {
		status = "1"
	} else if status == "inactive" {
		status = "0"
	} else {
		status = ""
	}

	customers, err := s.repo.FindAll(status)
	if err != nil {
		return nil, err
	}

	response := make([]dto.CustomerResponse, 0)
	for _, c := range customers {
		response = append(response, c.ToDto())
	}
	return response, err

}

func (s defaultCustomerService) GetCustomer(id string) (*dto.CustomerResponse, *errs.AppError) {
	customer, err := s.repo.ById(id)
	if err != nil {
		return nil, err
	}
	response := customer.ToDto()
	return &response, nil
}
