package entity

import "database/sql/driver"

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
