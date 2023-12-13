package domain

type User struct {
	Id             string `json:"id" bson:"_id"`
	Email          string `json:"email" bson:"email"`
	Password       string `json:"password" bson:"password"`
	FavoritePhrase string `json:"favorite_phrase" bson:"favorite_phrase"`
}

type (
	RegisterRequest struct {
		Email          string `json:"email" validate:"required,email"`
		Password       string `json:"password" validate:"required"`
		FavoritePhrase string `json:"favorite_phrase" validate:"required"`
	}

	RegisterResponse struct {
		Id    string `json:"id"`
		Email string `json:"email"`
	}

	LoginRequest struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	}

	LoginResponse struct {
		Token string `json:"token"`
	}
)
