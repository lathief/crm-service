package customer

import (
	"github.com/gin-gonic/gin"
	"github.com/lathief/crm-service/middleware"
	"github.com/lathief/crm-service/repository"
	"gorm.io/gorm"
)

type CustomerRoute struct {
	CustomerHandler CustomerRequestHandler
}

func NewRouter(db *gorm.DB, auth middleware.AuthorizationInterface) CustomerRoute {
	return CustomerRoute{
		CustomerHandler: &customerRequestHandler{
			CustomerController: &customerController{
				CustomerUseCase: &useCaseCustomer{
					CustomerRepo: repository.New(db),
				},
			},
			Auth: auth,
		},
	}
}

func (cr *CustomerRoute) Handle(router *gin.Engine) {
	customerPath := "/customer"
	customerRG := router.Group(customerPath)
	customerRG.POST("", cr.CustomerHandler.CreateCustomer)
	customerRG.GET("/:id", cr.CustomerHandler.GetCustomerById)
	customerRG.GET("/search", cr.CustomerHandler.SearchCustomers)
	customerRG.PUT("/:id", cr.CustomerHandler.UpdateCustomer)
	customerRG.DELETE("/:id", cr.CustomerHandler.DeleteCustomer)
}
