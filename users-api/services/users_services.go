package services

import (
	"github.com/go-microservices/users-api/domain"
	errors "github.com/go-microservices/users-api/utils"
)

func RegisterUser(user *domain.User) (*domain.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := user.RegisterUser(); err != nil {
		return nil, err
	}
	return user, nil
}

func LoginUser(userLogin *domain.UserLogin) (*domain.Token, *errors.RestErr) {
	if err := userLogin.ValidateLoginUser(); err != nil {
		return nil, err
	}
	res, err := userLogin.LoginUser()
	if err != nil {
		return nil, err
	}

	token, errGenerateToken := userLogin.GenerateToken(res.Id)
	if errGenerateToken != nil {
		return nil, errGenerateToken
	}
	return token, nil
}
