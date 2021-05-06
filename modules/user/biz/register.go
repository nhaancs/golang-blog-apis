package userbiz

import (
	"context"
	"nhaancs/common"
	"nhaancs/modules/user/model"
)

type RegisterStorage interface {
	FindUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
	CreateUser(ctx context.Context, data *usermodel.UserCreate) error
}

type Hasher interface {
	Hash(data string) string
}

type registerBusiness struct {
	registerStorage RegisterStorage
	hasher          Hasher
}

func NewRegisterBusiness(registerStorage RegisterStorage, hasher Hasher) *registerBusiness {
	return &registerBusiness{
		registerStorage: registerStorage,
		hasher:          hasher,
	}
}

func (business *registerBusiness) Register(ctx context.Context, data *usermodel.UserCreate) error {
	if err := data.Validate(); err != nil {
		return err
	}

	user, _ := business.registerStorage.FindUser(ctx, map[string]interface{}{"email": data.Email})
	if user != nil {
		return usermodel.ErrEmailExisted
	}

	salt := common.GenSalt(50)
	data.Password = business.hasher.Hash(data.Password + salt)
	data.Salt = salt
	data.Role = "user" // hard code
	if err := business.registerStorage.CreateUser(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(usermodel.EntityName, err)
	}

	return nil
}
