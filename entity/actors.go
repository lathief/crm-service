package entity

import (
	"database/sql/driver"
	"github.com/lathief/crm-service/utils/security"
	"gorm.io/gorm"
)

type boolType string

const (
	True  boolType = "true"
	False boolType = "false"
)

func (ct *boolType) Scan(value interface{}) error {
	*ct = boolType(value.([]byte))
	return nil
}

func (ct boolType) Value() (driver.Value, error) {
	return string(ct), nil
}

type Actor struct {
	GormModel
	Username   string
	Password   string
	RoleId     uint
	IsVerified boolType
	IsActive   boolType
	Role       *Role
}

func (Actor) TableName() string {
	return "actor"
}

func (a *Actor) BeforeCreate(tx *gorm.DB) (err error) {
	a.Password = security.HashPass(a.Password)
	return
}
func (a *Actor) BeforeUpdate(tx *gorm.DB) (err error) {
	a.Password = security.HashPass(a.Password)
	return
}
