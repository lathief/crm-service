package customer

import (
	"github.com/gin-gonic/gin"
	"github.com/lathief/crm-service/payload/response"
	"net/http"
	"strconv"
)

type customerRequestHandler struct {
	CustomerController CustomerController
}
type CustomerRequestHandler interface {
	CreateCustomer(c *gin.Context)
	GetCustomerById(c *gin.Context)
	UpdateCustomer(c *gin.Context)
	DeleteCustomer(c *gin.Context)
}

func (cr *customerRequestHandler) CreateCustomer(c *gin.Context) {
	custReq := new(CustomerDTO)
	err := c.ShouldBindJSON(&custReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.HandleFailedResponse(err.Error(), 400))
		return
	}
	res, err := cr.CustomerController.CreateCustomer(*custReq)
	if err != nil {
		c.JSON(res.Status, res)
		return
	}
	c.JSON(http.StatusOK, res)
}
func (cr *customerRequestHandler) GetCustomerById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, response.HandleFailedResponse(err.Error(), 400))
		return
	}
	res, err := cr.CustomerController.GetCustomerById(id)
	if err != nil {
		c.JSON(res.Status, res)
		return
	}
	c.JSON(http.StatusOK, res)
}
func (cr *customerRequestHandler) UpdateCustomer(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, response.HandleFailedResponse(err.Error(), 400))
		return
	}
	custReq := new(CustomerDTO)
	err = c.ShouldBindJSON(&custReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.HandleFailedResponse(err.Error(), 400))
		return
	}
	res, err := cr.CustomerController.UpdateCustomer(*custReq, id)
	if err != nil {
		c.JSON(res.Status, res)
		return
	}
	c.JSON(http.StatusOK, res)
}
func (cr *customerRequestHandler) DeleteCustomer(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, response.HandleFailedResponse(err.Error(), 400))
		return
	}
	res, err := cr.CustomerController.DeleteCustomer(id)
	if err != nil {
		c.JSON(res.Status, res)
		return
	}
	c.JSON(http.StatusOK, res)
}
