package request

type AuthActor struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UpdateActor struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UpdateFlagActor struct {
	IsVerified string `json:"is_verified"`
	IsActive   string `json:"is_active"`
}

type ApprovalStatus struct {
	Status string `json:"status"`
}
