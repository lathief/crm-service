package entity

import (
	"github.com/lathief/crm-service/constant"
	"github.com/lathief/crm-service/utils/security"
	"gorm.io/gorm"
)

type Actor struct {
	GormModel
	Username   string
	Password   string
	RoleId     uint
	IsVerified constant.BoolType
	IsActive   constant.BoolType
	Role       *Role
}

func (Actor) TableName() string {
	return "actor"
}

func (a *Actor) BeforeCreate(tx *gorm.DB) (err error) {
	a.Password = security.HashPass(a.Password)
	return
}
