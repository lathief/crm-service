package customer

import (
	"github.com/gin-gonic/gin"
	"github.com/lathief/crm-service/repository"
	"gorm.io/gorm"
)

type CustomerRoute struct {
	CustomerHandler CustomerRequestHandler
}

func NewRouter(db *gorm.DB) CustomerRoute {
	return CustomerRoute{
		CustomerHandler: &customerRequestHandler{
			CustomerController: &customerController{
				CustomerUseCase: &useCaseCustomer{
					CustomerRepo: repository.New(db),
				},
			},
		},
	}
}

func (cr *CustomerRoute) Handle(router *gin.Engine) {
	basePath := "/customer"
	customer := router.Group(basePath)
	customer.POST("", cr.CustomerHandler.CreateCustomer)
	customer.GET("/:id", cr.CustomerHandler.GetCustomerById)
	customer.PUT("/:id", cr.CustomerHandler.UpdateCustomer)
	customer.DELETE("/:id", cr.CustomerHandler.DeleteCustomer)
}
