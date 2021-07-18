package domain

import (
	"net/mail"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	errors "github.com/go-microservices/users-api/utils"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	FirstName       string `json:"first_name"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
}

type UserLogin struct {
	Id       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Token struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type Auth struct {
	Id string `json:"id"`
}

func (user *User) Validate() *errors.RestErr {
	if user.FirstName == "" || user.Email == "" || user.Password == "" || user.ConfirmPassword == "" {
		return errors.NewBadRequest("Must Input All Field")
	}

	if user.Password != user.ConfirmPassword {
		return errors.NewBadRequest("Password must same")
	}
	_, err := mail.ParseAddress(user.Email)

	if err != nil {
		return errors.NewBadRequest("Must Input Correct Email")
	}
	return nil
}

func (user *User) HashPassword(password string) (string, *errors.RestErr) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", errors.NewBadRequest("Failed Hash Password")
	}
	return string(hash), nil
}

func (userLogin *UserLogin) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (user *User) ValidatePagination(lmt string, oft string) (int, int, *errors.RestErr) {
	if lmt == "" {
		lmt = "0"
	}
	if oft == "" {
		oft = "0"
	}
	limit, err := strconv.ParseInt(lmt, 0, 64)

	if err != nil {
		return 0, 0, errors.NewBadRequest("invalid limit")
	}
	offset, err := strconv.ParseInt(oft, 0, 64)

	if err != nil {
		return 0, 0, errors.NewBadRequest("invalid offset")
	}

	return int(limit), int(offset), nil
}

func (userLogin *UserLogin) ValidateLoginUser() *errors.RestErr {

	return nil
}

func (userLogin *UserLogin) GenerateToken(id int) (*Token, *errors.RestErr) {
	token := &Token{}
	generateToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(time.Minute * 15).Unix(),
	})
	generateRefreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(time.Minute * 15).Unix(),
	})
	tokenString, err := generateToken.SignedString([]byte("accessToken"))
	refreshTokenString, errRefresh := generateRefreshToken.SignedString([]byte("refreshToken"))
	if err != nil || errRefresh != nil {
		return nil, errors.NewInternalServerError("Error Generate Token ")
	}
	token.AccessToken = tokenString
	token.RefreshToken = refreshTokenString
	return token, nil
}
