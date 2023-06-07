package customer

import (
	"context"
	"github.com/lathief/crm-service/payload"
	"github.com/lathief/crm-service/utils/helper"
	"strconv"
)

type customerController struct {
	CustomerUseCase UseCaseCustomer
}
type CustomerController interface {
	CreateCustomer(ctx context.Context, customer CustomerDTO) (payload.Response, error)
	SearchCustomer(ctx context.Context, filter map[string]string) (payload.Response, error)
	GetCustomerById(ctx context.Context, custId int) (payload.Response, error)
	UpdateCustomer(ctx context.Context, customer payload.UpdateCustomer, custId int) (payload.Response, error)
	DeleteCustomer(ctx context.Context, custId int) (payload.Response, error)
}

func (cc *customerController) CreateCustomer(ctx context.Context, customer CustomerDTO) (payload.Response, error) {
	err := cc.CustomerUseCase.CreateCustomer(ctx, customer)
	if err != nil {
		return payload.HandleFailedResponse(err.Error(), helper.GetStatusCode(err)), err
	}
	return payload.HandleSuccessResponse(nil, "Create Customer Successfully", 201), err
}
func (cc *customerController) SearchCustomer(ctx context.Context, filter map[string]string) (payload.Response, error) {
	customers, err := cc.CustomerUseCase.SearchCustomer(ctx, filter)
	if err != nil {
		return payload.HandleFailedResponse(err.Error(), helper.GetStatusCode(err)), err
	}
	return payload.HandleSuccessResponse(customers, "Success Search Customer by :"+filter["name"]+" "+filter["email"], 200), err
}
func (cc *customerController) GetCustomerById(ctx context.Context, custId int) (payload.Response, error) {
	user, err := cc.CustomerUseCase.GetCustomerById(ctx, custId)
	if err != nil {
		return payload.HandleFailedResponse(err.Error(), helper.GetStatusCode(err)), err
	}
	return payload.HandleSuccessResponse(user, "Success Get Customer By ID : "+strconv.Itoa(custId), 200), err
}
func (cc *customerController) UpdateCustomer(ctx context.Context, customer payload.UpdateCustomer, custId int) (payload.Response, error) {
	err := cc.CustomerUseCase.UpdateCustomer(ctx, customer, custId)
	if err != nil {
		return payload.HandleFailedResponse(err.Error(), helper.GetStatusCode(err)), err
	}
	return payload.HandleSuccessResponse(nil, "Success Update Customer", 200), err
}
func (cc *customerController) DeleteCustomer(ctx context.Context, custId int) (payload.Response, error) {
	err := cc.CustomerUseCase.DeleteCustomer(ctx, custId)
	if err != nil {
		return payload.HandleFailedResponse(err.Error(), helper.GetStatusCode(err)), err
	}
	return payload.HandleSuccessResponse(nil, "Success Delete Customer", 200), err
}
