package dto

type RegisterRequest struct {
	Username       string `json:"username"`
	Fullname       string `json:"fullname" validate:"required"`
	Email          string `json:"email" validate:"required,email"`
	Password       string `json:"password" validate:"required"`
	FavoritePhrase string `json:"favorite_phrase" validate:"required"`
}

type RegisterResponse struct {
	Id    string `json:"id"`
	Email string `json:"email"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	Email string `json:"email"`
	Token string `json:"token"`
}
