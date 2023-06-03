package entity

import (
	"database/sql/driver"
	"github.com/lathief/crm-service/utils/security"
	"gorm.io/gorm"
)

type BoolType string

const (
	True  BoolType = "true"
	False BoolType = "false"
)

func (ct *BoolType) Scan(value interface{}) error {
	*ct = BoolType(value.([]byte))
	return nil
}

func (ct BoolType) Value() (driver.Value, error) {
	return string(ct), nil
}

type Actor struct {
	GormModel
	Username   string
	Password   string
	RoleId     uint
	IsVerified BoolType
	IsActive   BoolType
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
