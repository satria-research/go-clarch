package helper

import (
	"fmt"
	"strings"

	gonanoid "github.com/matoous/go-nanoid/v2"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}
	return string(hashedPassword), nil
}

func CheckPassword(password string, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func RandomUsername(fullname string) string {
	fullnameLowercase := strings.ToLower(fullname)
	firstName := strings.Split(fullnameLowercase, " ")[0]
	shortIDLowercase := strings.ToLower(gonanoid.Must())

	return fmt.Sprintf("%s%s", firstName, strings.Replace(shortIDLowercase, "-", "", -1))
}
