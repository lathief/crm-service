package entity

type Customer struct {
	GormModel
	Firstname string
	Lastname  string
	Email     string
	Avatar    string
}
