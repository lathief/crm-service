package customer

import (
	"github.com/lathief/crm-service/entity"
	"github.com/lathief/crm-service/repository"
	"github.com/lathief/crm-service/utils/helper"
	"strconv"
)

type useCaseCustomer struct {
	CustomerRepo repository.CustomerInterfaceRepository
}
type UseCaseCustomer interface {
	CreateCustomer(customer CustomerDTO) error
	GetCustomerById(id int) (CustomerDTO, error)
	SearchCustomer(filter map[string]string) (*helper.Pagination, error)
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
func (uc *useCaseCustomer) SearchCustomer(filter map[string]string) (*helper.Pagination, error) {
	var customers *helper.Pagination
	var totalRows int64
	var err error
	page, err := strconv.Atoi(filter["page"])
	if err != nil {
		return &helper.Pagination{}, err
	}
	limit, err := strconv.Atoi(filter["limit"])
	if err != nil {
		return &helper.Pagination{}, err
	}
	pagination := helper.Pagination{
		Limit: limit,
		Page:  page,
		Sort:  "Id desc",
	}
	err = uc.CustomerRepo.CountRowCustomer(&totalRows)
	if err != nil {
		return &helper.Pagination{}, err
	}
	if filter["name"] != "" && filter["email"] == "" {
		customers, err = uc.CustomerRepo.SearchCustomerByName(pagination, filter["name"], totalRows)
		if err != nil {
			return &helper.Pagination{}, err
		}
	} else if filter["email"] != "" && filter["name"] == "" {
		customers, err = uc.CustomerRepo.SearchCustomerByEmail(pagination, filter["email"], totalRows)
		if err != nil {
			return &helper.Pagination{}, err
		}
	} else if filter["name"] != "" && filter["email"] != "" {
		customers, err = uc.CustomerRepo.SearchCustomerByNameOrEmail(pagination, filter["name"], filter["email"], totalRows)
		if err != nil {
			return &helper.Pagination{}, err
		}
		//customersName, err := uc.CustomerRepo.SearchCustomerByEmail(filter["email"])
		//if err != nil {
		//	return nil, err
		//}
		//customersEmail, err := uc.CustomerRepo.SearchCustomerByName(filter["name"])
		//if err != nil {
		//	return nil, err
		//}
		//customers = append(customers, customersName...)
		//customers = append(customers, customersEmail...)
	} else {
		customers, err = uc.CustomerRepo.GetAllCustomer(pagination, totalRows)
		if err != nil {
			return &helper.Pagination{}, err
		}
	}
	//for _, item := range customers {
	//	customerDTO := CustomerDTO{
	//		Firstname: item.Firstname,
	//		Lastname:  item.Lastname,
	//		Email:     item.Email,
	//		Avatar:    item.Avatar,
	//	}
	//	customersDTO = append(customersDTO, customerDTO)
	//}
	return customers, nil
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
