package customer

import (
	"github.com/lathief/crm-service/payload/response"
	"net/http"
	"strconv"
)

type customerController struct {
	CustomerUseCase UseCaseCustomer
}
type CustomerController interface {
	CreateCustomer(customer CustomerDTO) (response.Response, error)
	GetCustomerById(custId int) (response.Response, error)
	UpdateCustomer(customer CustomerDTO, custId int) (response.Response, error)
	DeleteCustomer(custId int) (response.Response, error)
}

func (cc *customerController) CreateCustomer(customer CustomerDTO) (response.Response, error) {
	err := cc.CustomerUseCase.CreateCustomer(customer)
	if err != nil {
		return response.HandleFailedResponse(err.Error(), 500), err
	}
	return response.HandleSuccessResponse(nil, "Create Customer Successfully", 500), err
}

func (cc *customerController) GetCustomerById(custId int) (response.Response, error) {
	user, err := cc.CustomerUseCase.GetCustomerById(custId)
	if err != nil {
		return response.HandleFailedResponse(err.Error(), http.StatusNotFound), err
	}
	return response.HandleSuccessResponse(user, "Success Get Customer By ID : "+strconv.Itoa(custId), 200), err
}
func (cc *customerController) UpdateCustomer(customer CustomerDTO, custId int) (response.Response, error) {
	err := cc.CustomerUseCase.UpdateCustomer(customer, custId)
	if err != nil {
		return response.HandleFailedResponse(err.Error(), http.StatusInternalServerError), err
	}
	return response.HandleSuccessResponse(nil, "Success Update Customer", 200), err
}

func (cc *customerController) DeleteCustomer(custId int) (response.Response, error) {
	err := cc.CustomerUseCase.DeleteCustomer(custId)
	if err != nil {
		return response.HandleFailedResponse(err.Error(), http.StatusInternalServerError), err
	}
	return response.HandleSuccessResponse(nil, "Success Delete Customer", 200), err
}
