package entity

import "database/sql/driver"

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
