package actor

type ActorDTO struct {
	Username   string `json:"username" valid:"required~Firstname is required"`
	Password   string `json:"password" valid:"required~Password is required"`
	IsVerified string `json:"isVerified"`
	IsActive   string `json:"isActive"`
	Role       string `json:"role"`
}
