package customer

type CustomerDTO struct {
	Firstname string `json:"firstname" validate:"required"`
	Lastname  string `json:"lastname" validate:"required"`
	Email     string `json:"email"`
	Avatar    string `json:"avatar"`
}
