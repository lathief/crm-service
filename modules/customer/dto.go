package customer

type CustomerDTO struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Avatar    string `json:"avatar"`
}
