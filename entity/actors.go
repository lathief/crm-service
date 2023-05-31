package entity

type Actors struct {
	GormModel
	Username   string
	Password   string
	RoleId     uint
	IsVerified bool
	IsActive   bool
	Role       *Role
}
