package customer

import (
	"github.com/lathief/crm-service/entity"
	"github.com/lathief/crm-service/repository"
)

type useCaseCustomer struct {
	CustomerRepo repository.CustomerInterfaceRepository
}
type UseCaseCustomer interface {
	CreateCustomer(customer CustomerDTO) error
	GetCustomerById(id int) (CustomerDTO, error)
	UpdateCustomer(customer CustomerDTO, id int) error
	DeleteCustomer(id int) error
}

func (uc *useCaseCustomer) CreateCustomer(customer CustomerDTO) error {
	customerSave := entity.Customer{
		Firstname: customer.Firstname,
		Lastname:  customer.Lastname,
		Email:     customer.Email,
		Avatar:    customer.Avatar,
	}
	err := uc.CustomerRepo.CreateCustomer(customerSave)
	if err != nil {
		return err
	}
	return nil
}
func (uc *useCaseCustomer) GetCustomerById(id int) (CustomerDTO, error) {
	get, err := uc.CustomerRepo.GetCustomerById(uint(id))
	if err != nil {
		return CustomerDTO{}, err
	}
	getCust := CustomerDTO{
		Firstname: get.Firstname,
		Lastname:  get.Lastname,
		Email:     get.Email,
		Avatar:    get.Avatar,
	}
	return getCust, nil
}
func (uc *useCaseCustomer) UpdateCustomer(customer CustomerDTO, id int) error {
	_, err := uc.CustomerRepo.GetCustomerById(uint(id))
	if err != nil {
		return err
	}
	customerUpdate := entity.Customer{
		Firstname: customer.Firstname,
		Lastname:  customer.Lastname,
		Email:     customer.Email,
		Avatar:    customer.Avatar,
	}
	err = uc.CustomerRepo.UpdateCustomer(customerUpdate, uint(id))
	if err != nil {
		return err
	}
	return nil
}
func (uc *useCaseCustomer) DeleteCustomer(id int) error {
	_, err := uc.CustomerRepo.GetCustomerById(uint(id))
	if err != nil {
		return err
	}
	err = uc.CustomerRepo.DeleteCustomer(uint(id))
	if err != nil {
		return err
	}
	return nil
}
