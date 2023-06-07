package payload

type AuthActor struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UpdateActor struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UpdateCustomer struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Avatar    string `json:"avatar"`
}

type UpdateFlagActor struct {
	IsVerified string `json:"is_verified" validate:"required,enum-flag"`
	IsActive   string `json:"is_active" validate:"required,enum-flag"`
}

type ApprovalStatus struct {
	Status string `json:"status" validate:"required,enum-approval"`
}
