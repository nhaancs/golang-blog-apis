package usermodel

import (
	"strings"

	"github.com/badoux/checkmail"
)

type UserLogin struct {
	Email    string `json:"email" form:"email" gorm:"column:email;"`
	Password string `json:"password" form:"password" gorm:"column:password;"`
}

func (UserLogin) TableName() string {
	return User{}.TableName()
}

func (res *UserLogin) Validate() error {
	res.Email = strings.TrimSpace(res.Email)
	res.Password = strings.TrimSpace(res.Password)

	if len(res.Email) == 0 {
		return ErrEmailCannotBeEmpty
	}
	if err := checkmail.ValidateFormat(res.Email); err != nil {
		return ErrInvalidEmail
	}
	if len(res.Password) == 0 {
		return ErrPasswordCannotBeEmpty
	}
	if len(res.Password) < 6 || len(res.Password) > 50 {
		return ErrInvalidPassword
	}

	return nil
}