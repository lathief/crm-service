package entity

type Approval struct {
	ID            uint
	Admin_id      uint
	Admin         *Actors
	Superadmin_id uint
	Superadmin    *Actors
	Status        string
}
