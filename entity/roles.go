package entity

const (
	ROLE_ADMIN       = "role_admin"
	ROLE_SUPER_ADMIN = "role_super_admin"
)

type Role struct {
	ID       uint
	Rolename string
}

func (Role) TableName() string {
	return "role"
}
