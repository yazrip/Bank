package service

import (
	"bank/entity"
	"bank/repo"
)

func NewCustomerService(customerRepo *repo.CustomerRepo) CustomerService {
	return &CustomerServiceImpl{CustomerRepo: *customerRepo}
}

type CustomerServiceImpl struct {
	CustomerRepo repo.CustomerRepo
}

func (service *CustomerServiceImpl) GetAll() []entity.Customer {
	customers := service.CustomerRepo.GetAll()
	return customers
}
