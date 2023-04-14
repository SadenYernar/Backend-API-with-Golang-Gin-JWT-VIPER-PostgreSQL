package service

import (
	"Backend-API-with-Golang-Gin-JWT-VIPER-PostgreSQL/internal/model"
	"Backend-API-with-Golang-Gin-JWT-VIPER-PostgreSQL/internal/repository"
	"Backend-API-with-Golang-Gin-JWT-VIPER-PostgreSQL/internal/token"
	"fmt"
	"net/http"

	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo repository.Repository
	gen  uuid.Generator
}

func NewUserService(repo repository.Repository) *UserService {
	return &UserService{
		repo: repo,
		gen:  uuid.NewGen(),
	}
}

func (u *UserService) CreateUserService(user model.User) (int, model.User, error) {
	var err error

	if !userValidation(user) {
		return http.StatusBadRequest, user, fmt.Errorf("Create user in service: %w", err)
	}

	user.Uuid, err = u.gen.NewV4()

	if err != nil {
		return http.StatusInternalServerError, user, fmt.Errorf("Create user in service: %w", err)
	}
	user.Password, err = generateHashPassword(user.Password)
	if err != nil {
		return http.StatusInternalServerError, user, fmt.Errorf("Create user in service: %w", err)
	}
	return u.repo.CreateUser(user)
}

func (u *UserService) AuthorizationUserService(user model.User) (string, error) {
	var err error

	userInfoResponse, err := u.repo.GetUserInfoByEmail(user)
	if err != nil {
		return "User is exist", err
	}

	if userInfoResponse.Email != user.Email {
		return "Not correct email", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(userInfoResponse.Password), []byte(user.Password))
	if err != nil {
		return "Not correct password", err
	}

	value, err := u.CreateSessionService(userInfoResponse)
	if err != nil {
		return "Session not created", err
	}

	return value, err
}

func (u *UserService) GetUserInfoServiceByUsername(username string) (model.User, error) {
	return u.repo.GetUserInfoByUsername(username)
}

func (u *UserService) CreateSessionService(user model.User) (string, error) {
	token, err := CreateToken(user)
	if err != nil {
		return "Token not created", err
	}

	return token, u.repo.SetSession(user, token)
}

func CreateToken(user model.User) (string, error) {
	tokenString, err := token.GenerateJWT(user.Email, user.Username)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
