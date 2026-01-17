package domain

type PasswordHasher interface {
	Hash(password string) (string, error)
	Compare(hashedPassword, password string) error
}

type UsernameGenerator interface {
	Generate(fullname string) string
}

type TokenGenerator interface {
	GenerateAccessToken(user *User, expiry int) (string, error)
	GenerateRefreshToken(user *User, expiry int) (string, error)
}

type ConfigProvider interface {
	Get(key string) string
}
