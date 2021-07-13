package domain

import (
	"github.com/go-microservices/users-api/repositories"
	errors "github.com/go-microservices/users-api/utils"
)

func (user *User) RegisterUser() *errors.RestErr {
	if result := repositories.Client.Model(&user).Select("first_name", "email", "password").Create(&user); result.Error != nil {
		return errors.NewInternalServerError("Something Bad Happened")
	}

	return nil

}

func (userLogin *UserLogin) LoginUser() (*UserLogin, *errors.RestErr) {
	user := &UserLogin{}
	if result := repositories.Client.Table("users").Where("email = ?", userLogin.Email).Find(&user); result.Error != nil {
		return nil, errors.NewInternalServerError("Something Bad Happened")
	}
	if userLogin.Password != user.Password {
		return nil, errors.NewBadRequest("Wrong Password")
	}
	return user, nil
}
