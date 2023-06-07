package customer

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/lathief/crm-service/middleware"
	"github.com/lathief/crm-service/payload"
	"net/http"
	"strconv"
)

type customerRequestHandler struct {
	CustomerController CustomerController
	Auth               middleware.AuthorizationInterface
	Validation         middleware.ValidationInterface
}
type CustomerRequestHandler interface {
	CreateCustomer(c *gin.Context)
	GetCustomerById(c *gin.Context)
	SearchCustomers(c *gin.Context)
	UpdateCustomer(c *gin.Context)
	DeleteCustomer(c *gin.Context)
}

func (cr *customerRequestHandler) CreateCustomer(c *gin.Context) {
	err := cr.Auth.Authentication(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, payload.HandleFailedResponse(err.Error(), 401))
		return
	}
	//custReq := new(CustomerDTO)
	//err = c.ShouldBindJSON(&custReq)
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, payload.HandleFailedResponse(err.Error(), 400))
	//	return
	//}
	var custReq CustomerDTO
	if errs := cr.Validation.BindAndValidate(c, &custReq); len(errs) > 0 {
		c.JSON(http.StatusBadRequest, payload.HandleSuccessResponse(errs, "", 400))
		return
	}
	fmt.Println("error")
	ctx := context.Background()
	adminData := c.MustGet("adminData").(jwt.MapClaims)
	ctx = context.WithValue(ctx, "adminId", uint(adminData["id"].(float64)))
	res, err := cr.CustomerController.CreateCustomer(ctx, custReq)
	if err != nil {
		c.JSON(res.Status, res)
		return
	}
	c.JSON(http.StatusOK, res)
}
func (cr *customerRequestHandler) GetCustomerById(c *gin.Context) {
	err := cr.Auth.Authentication(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, payload.HandleFailedResponse(err.Error(), 401))
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, payload.HandleFailedResponse(err.Error(), 400))
		return
	}
	ctx := context.Background()
	adminData := c.MustGet("adminData").(jwt.MapClaims)
	ctx = context.WithValue(ctx, "adminId", uint(adminData["id"].(float64)))
	res, err := cr.CustomerController.GetCustomerById(ctx, id)
	if err != nil {
		c.JSON(res.Status, res)
		return
	}
	c.JSON(http.StatusOK, res)
}
func (cr *customerRequestHandler) SearchCustomers(c *gin.Context) {
	err := cr.Auth.Authentication(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, payload.HandleFailedResponse(err.Error(), 401))
		return
	}
	name := c.Query("name")
	email := c.Query("email")
	pageStr := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "10")
	sortBy := c.DefaultQuery("sort_by", "id")
	orderBy := c.DefaultQuery("order_by", "asc")
	filter := map[string]string{
		"name":    name,
		"email":   email,
		"page":    pageStr,
		"limit":   limit,
		"sortby":  sortBy,
		"orderby": orderBy,
	}
	ctx := context.Background()
	adminData := c.MustGet("adminData").(jwt.MapClaims)
	ctx = context.WithValue(ctx, "adminId", uint(adminData["id"].(float64)))
	res, err := cr.CustomerController.SearchCustomer(ctx, filter)
	if err != nil {
		c.JSON(res.Status, res)
		return
	}
	c.JSON(http.StatusOK, res)
}
func (cr *customerRequestHandler) UpdateCustomer(c *gin.Context) {
	err := cr.Auth.Authentication(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, payload.HandleFailedResponse(err.Error(), 401))
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, payload.HandleFailedResponse(err.Error(), 400))
		return
	}
	custReq := new(payload.UpdateCustomer)
	err = c.ShouldBindJSON(&custReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, payload.HandleFailedResponse(err.Error(), 400))
		return
	}
	ctx := context.Background()
	adminData := c.MustGet("adminData").(jwt.MapClaims)
	ctx = context.WithValue(ctx, "adminId", uint(adminData["id"].(float64)))
	res, err := cr.CustomerController.UpdateCustomer(ctx, *custReq, id)
	if err != nil {
		c.JSON(res.Status, res)
		return
	}
	c.JSON(http.StatusOK, res)
}
func (cr *customerRequestHandler) DeleteCustomer(c *gin.Context) {
	err := cr.Auth.Authentication(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, payload.HandleFailedResponse(err.Error(), 401))
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, payload.HandleFailedResponse(err.Error(), 400))
		return
	}
	ctx := context.Background()
	adminData := c.MustGet("adminData").(jwt.MapClaims)
	ctx = context.WithValue(ctx, "adminId", uint(adminData["id"].(float64)))
	res, err := cr.CustomerController.DeleteCustomer(ctx, id)
	if err != nil {
		c.JSON(res.Status, res)
		return
	}
	c.JSON(http.StatusOK, res)
}
