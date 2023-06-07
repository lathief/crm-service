package entity

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Customer struct {
	GormModel
	Firstname string
	Lastname  string
	Email     string `valid:"email"`
	Avatar    string `valid:"url"`
}

func (Customer) TableName() string {
	return "customer"
}
func (c *Customer) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(c)

	if errCreate != nil {
		err = errCreate
		return
	}
	return
}
