package actor

type ActorDTO struct {
	Username   string `json:"username"`
	Role       string `json:"role"`
	IsVerified string `json:"is_verified"`
	IsActive   string `json:"is_active"`
}

type ApprovalDTO struct {
	ID     uint     `json:"approval_id"`
	Admin  ActorDTO `json:"admin"`
	Status string   `json:"status"`
}
