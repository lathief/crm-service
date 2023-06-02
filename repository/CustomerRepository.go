package repository

import (
	"fmt"
	"github.com/lathief/crm-service/entity"
	"gorm.io/gorm"
)

type CustomerRepository struct {
	db *gorm.DB
}

type CustomerInterfaceRepository interface {
	CreateCustomer(customer entity.Customer) error
	GetCustomerById(id uint) (entity.Customer, error)
	UpdateCustomer(customer entity.Customer, id uint) error
	DeleteCustomer(id uint) error
}

func New(db *gorm.DB) *CustomerRepository {
	return &CustomerRepository{db: db}
}

func (c *CustomerRepository) CreateCustomer(customer entity.Customer) error {
	err := c.db.Model(&entity.Customer{}).Create(&customer).Error
	return err
}
func (c *CustomerRepository) GetCustomerById(id uint) (entity.Customer, error) {
	var customer entity.Customer
	err := c.db.First(&customer, "id = ? ", id).Error
	return customer, err
}
func (c *CustomerRepository) UpdateCustomer(customer entity.Customer, id uint) error {
	fmt.Println(customer)
	err := c.db.Model(&entity.Customer{}).Where("id = ?", id).Updates(entity.Customer{
		Firstname: customer.Firstname, Lastname: customer.Lastname, Email: customer.Email, Avatar: customer.Avatar}).Error
	return err
}
func (c *CustomerRepository) DeleteCustomer(id uint) error {
	err := c.db.First(&entity.Customer{}).Where("id = ?", id).Delete(&entity.Customer{}).Error
	return err
}
