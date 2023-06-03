package actor

type ActorDTO struct {
	Username   string `json:"username" valid:"required~Firstname is required"`
	Password   string `json:"password" valid:"required~Password is required"`
	Role       string `json:"role"`
	IsVerified string `json:"isVerified"`
	IsActive   string `json:"isActive"`
}

type ApprovalDTO struct {
	Admin  ActorDTO `json:"admin"`
	Status string   `json:"status"`
}
