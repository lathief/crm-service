package customer

type CustomerDTO struct {
	Firstname string `json:"firstname" valid:"required~Firstname is required"`
	Lastname  string `json:"lastname" valid:"required~Lastname is required"`
	Email     string `json:"email" valid:"required~Email is required, email~Email is invalid"`
	Avatar    string `json:"avatar"`
}
