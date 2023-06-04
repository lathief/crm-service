package repository

import (
	"fmt"
	"github.com/lathief/crm-service/entity"
	"github.com/lathief/crm-service/utils/helper"
	"gorm.io/gorm"
)

type CustomerRepository struct {
	db *gorm.DB
}

type CustomerInterfaceRepository interface {
	GetCustomerById(id uint) (entity.Customer, error)
	GetAllCustomer(pagination helper.Pagination, totalRows int64) (*helper.Pagination, error)
	CountRowCustomer(totalRows *int64) error
	SearchCustomerByName(pagination helper.Pagination, name string, totalRows int64) (*helper.Pagination, error)
	SearchCustomerByEmail(pagination helper.Pagination, email string, totalRows int64) (*helper.Pagination, error)
	SearchCustomerByNameOrEmail(pagination helper.Pagination, name string, email string, totalRows int64) (*helper.Pagination, error)
	CreateCustomer(customer entity.Customer) error
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
func (c *CustomerRepository) GetAllCustomer(pagination helper.Pagination, totalRows int64) (*helper.Pagination, error) {
	var customers []*entity.Customer
	err := c.db.Scopes(helper.Paginate(customers, &pagination, totalRows)).Find(&customers).Error
	pagination.Rows = customers
	return &pagination, err
}
func (c *CustomerRepository) CountRowCustomer(totalRows *int64) error {
	err := c.db.Model(&entity.Customer{}).Count(totalRows).Error
	return err
}
func (c *CustomerRepository) SearchCustomerByName(pagination helper.Pagination, name string, totalRows int64) (*helper.Pagination, error) {
	var customers []*entity.Customer
	err := c.db.Scopes(helper.Paginate(customers, &pagination, totalRows)).Where("CONCAT(firstname, \" \", lastname) LIKE ?", "%"+name+"%").Find(&customers).Error
	//err := c.db.Model(&entity.Customer{}).Limit(page).Offset(offset).Where("CONCAT(firstname, \" \", lastname) LIKE ?", "%"+name+"%").Find(&customers).Error
	pagination.Rows = customers
	return &pagination, err
}
func (c *CustomerRepository) SearchCustomerByEmail(pagination helper.Pagination, email string, totalRows int64) (*helper.Pagination, error) {
	var customers []*entity.Customer
	//err := c.db.Model(&entity.Customer{}).Limit(page).Offset(offset).Where("email LIKE ?", "%"+email+"%").Find(&customers).Error
	err := c.db.Scopes(helper.Paginate(customers, &pagination, totalRows)).Where("email LIKE ?", "%"+email+"%").Find(&customers).Error
	pagination.Rows = customers
	return &pagination, err
}
func (c *CustomerRepository) SearchCustomerByNameOrEmail(pagination helper.Pagination, name string, email string, totalRows int64) (*helper.Pagination, error) {
	var customers []*entity.Customer
	err := c.db.Scopes(helper.Paginate(customers, &pagination, totalRows)).Where("CONCAT(firstname, \" \", lastname) "+
		"LIKE ?", "%"+name+"%").Or("email LIKE ?", "%"+email+"%").Find(&customers).Error
	pagination.Rows = customers
	//err := c.db.Model(&entity.Customer{}).Limit(page).Offset(offset).Where("CONCAT(firstname, \" \", lastname) "+
	//"LIKE ?", "%"+name+"%").Or("email LIKE ?", "%"+email+"%").Find(&customers).Error
	return &pagination, err
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
