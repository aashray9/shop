package models

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	//"golang.org/x/crypto/bcrypt"
)

type CrmUsers struct {
	Id       int     `json:"id"`
	Uuid     string  `json:"uuid"`
	Name     string  `json:"name"`
	Email    *string `json:"email"`
	Mobile   string  `json:"mobile"`
	Password string  `json:"password"`
	EmpCode  string  `json:"emp_code"`
	UserType string  `json:"user_type"`
	DeptId   int     `json:"dept_id"`
}

func (user *CrmUsers) GenerateJwtToken() string {

	jwt_token := jwt.New(jwt.SigningMethodHS512)

	jwt_token.Claims = jwt.MapClaims{
		"user_id":    user.Id,
		"email":      user.Email,
		"department": user.DeptId,
		"empcode":    user.EmpCode,
		"exp":        time.Now().Add(time.Hour * 24).Unix(),
	}
	// Sign and get the complete encoded token as a string
	token, _ := jwt_token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	return token
}
