package usermodel

import (
	"github.com/badoux/checkmail"
	"nhaancs/common"
	"strings"
)

type UserCreate struct {
	common.SQLCreateModel `json:",inline"`
	Email                 string        `json:"email" gorm:"column:email;"`
	Password              string        `json:"password" gorm:"column:password;"`
	LastName              string        `json:"last_name" gorm:"column:last_name;"`
	FirstName             string        `json:"first_name" gorm:"column:first_name;"`
	Bio                   string        `json:"bio" gorm:"column:bio;"`
	Role                  string        `json:"-" gorm:"column:role;"`
	Salt                  string        `json:"-" gorm:"column:salt;"`
	Avatar                *common.Image `json:"avatar,omitempty" gorm:"column:avatar;type:json"`
}

func (UserCreate) TableName() string {
	return User{}.TableName()
}

func (u *UserCreate) Mask(isAdmin bool) {
	u.GenUID(common.DbTypeUser)
}

func (res *UserCreate) Validate() error {
	res.Email = strings.TrimSpace(res.Email)
	res.Password = strings.TrimSpace(res.Password)
	res.FirstName = strings.TrimSpace(res.FirstName)
	res.LastName = strings.TrimSpace(res.LastName)
	res.Bio = strings.TrimSpace(res.Bio)

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
	if len(res.FirstName) == 0 {
		return ErrFirstNameCannotBeEmpty
	}
	if len(res.FirstName) > 200 {
		return ErrFirstNameIsTooLong
	}
	if len(res.LastName) == 0 {
		return ErrLastNameCannotBeEmpty
	}
	if len(res.LastName) > 200 {
		return ErrLastNameIsTooLong
	}
	if len(res.Bio) > 500 {
		return ErrBioIsTooLong
	}
	if res.Avatar == nil {
		return ErrAvatarCannotBeEmpty
	}

	return nil
}
