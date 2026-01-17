package service

import (
	"github.com/ubaidillahhf/go-clarch/app/domain"
	"golang.org/x/crypto/bcrypt"
)

type passwordHasher struct{}

func NewPasswordHasher() domain.PasswordHasher {
	return &passwordHasher{}
}

func (p *passwordHasher) Hash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (p *passwordHasher) Compare(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
