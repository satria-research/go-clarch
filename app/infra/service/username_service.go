package service

import (
	"fmt"
	"strings"

	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/ubaidillahhf/go-clarch/app/domain"
)

type usernameGenerator struct{}

func NewUsernameGenerator() domain.UsernameGenerator {
	return &usernameGenerator{}
}

func (u *usernameGenerator) Generate(fullname string) string {
	fullnameLowercase := strings.ToLower(fullname)
	firstName := strings.Split(fullnameLowercase, " ")[0]
	shortIDLowercase := strings.ToLower(gonanoid.Must())

	return fmt.Sprintf("%s%s", firstName, strings.Replace(shortIDLowercase, "-", "", -1))
}
