package entity

type Approval struct {
	ID            uint
	Admin_id      uint
	Admin         *Actor
	Superadmin_id uint
	Superadmin    *Actor
	Status        string
}

func (Approval) TableName() string {
	return "approval"
}
