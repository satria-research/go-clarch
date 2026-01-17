package usecases

import (
	"context"
	"errors"
	"strconv"

	"github.com/ubaidillahhf/go-clarch/app/domain"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserUsecase interface {
	Register(ctx context.Context, username, fullname, email, password, favoritePhrase string) (domain.User, error)
	Login(ctx context.Context, email, password string) (string, error)
}

func NewUserUsecase(
	repo domain.UserRepository,
	passwordHasher domain.PasswordHasher,
	usernameGenerator domain.UsernameGenerator,
	tokenGenerator domain.TokenGenerator,
	config domain.ConfigProvider,
) UserUsecase {
	return &userUsecase{
		repo:              repo,
		passwordHasher:    passwordHasher,
		usernameGenerator: usernameGenerator,
		tokenGenerator:    tokenGenerator,
		config:            config,
	}
}

type userUsecase struct {
	repo              domain.UserRepository
	passwordHasher    domain.PasswordHasher
	usernameGenerator domain.UsernameGenerator
	tokenGenerator    domain.TokenGenerator
	config            domain.ConfigProvider
}

func (uc *userUsecase) Register(ctx context.Context, username, fullname, email, password, favoritePhrase string) (domain.User, error) {
	existingUser, err := uc.repo.FindByIdentifier(ctx, username, email)
	if err != nil && err != mongo.ErrNoDocuments {
		return domain.User{}, err
	}
	if existingUser.Id != "" {
		return domain.User{}, errors.New("username or email already registered")
	}

	if username == "" {
		username = uc.usernameGenerator.Generate(fullname)
	}

	hashedPassword, err := uc.passwordHasher.Hash(password)
	if err != nil {
		return domain.User{}, err
	}

	newUser := domain.User{
		Username:       username,
		Fullname:       fullname,
		Email:          email,
		Password:       hashedPassword,
		FavoritePhrase: favoritePhrase,
	}

	return uc.repo.Insert(ctx, newUser)
}

func (uc *userUsecase) Login(ctx context.Context, email, password string) (string, error) {
	user, err := uc.repo.FindByIdentifier(ctx, "", email)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return "", errors.New("user not found")
		}
		return "", err
	}

	if err := uc.passwordHasher.Compare(user.Password, password); err != nil {
		return "", errors.New("wrong password")
	}

	exp := uc.config.Get("ACCESS_TOKEN_EXPIRY_HOUR")
	expAsInt, _ := strconv.Atoi(exp)

	token, err := uc.tokenGenerator.GenerateAccessToken(&user, expAsInt)
	if err != nil {
		return "", err
	}

	return token, nil
}
