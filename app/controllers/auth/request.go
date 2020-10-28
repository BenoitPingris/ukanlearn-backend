package auth

// LoginRequest  struct
type LoginRequest struct {
	Email    string `json:"email" validate:"email"`
	Password string `json:"password" validate:"min=6"`
}

// RegisterRequest struct
type RegisterRequest struct {
	LoginRequest
	Confirm string `json:"confirm" validate:"eqfield=Password"`
}
