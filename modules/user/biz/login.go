package userbiz

import (
	"context"
	"nhaancs/common"
	"nhaancs/component/tokenprovider"
	"nhaancs/modules/user/model"

	"go.opencensus.io/trace"
)

type LoginStorage interface {
	FindUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
}

type loginBiz struct {
	storeUser     LoginStorage
	tokenProvider tokenprovider.Provider
	hasher        Hasher
	expiry        int
}

func NewLoginBusiness(storeUser LoginStorage, tokenProvider tokenprovider.Provider, hasher Hasher, expiry int) *loginBiz {
	return &loginBiz{
		storeUser:     storeUser,
		tokenProvider: tokenProvider,
		hasher:        hasher,
		expiry:        expiry,
	}
}

// 1. Find user, email
// 2. Hash pass from input and compare with pass in db
// 3. Provider: issue JWT token for client
// 3.1. Access token and refresh token
// 4. Return token(s)
func (business *loginBiz) Login(ctx context.Context, data *usermodel.UserLogin) (*tokenprovider.Token, error) {
	if err := data.Validate(); err != nil {
		return nil, err
	}

	// note: span io only (read file, database, call other apis,...)
	ctx1, span1 := trace.StartSpan(ctx, "user.biz.login.find-user")
	span1.AddAttributes(
		trace.StringAttribute("email", data.Email),
		trace.StringAttribute("email-again", data.Email),
	)
	// note: use new created context ctx1
	user, err := business.storeUser.FindUser(ctx1, map[string]interface{}{"email": data.Email})
	span1.End()
	if err != nil {
		return nil, usermodel.ErrEmailOrPasswordInvalid
	}

	if user.DeletedAt != nil || !user.IsEnabled {
		panic(common.NewCustomError(nil, "user has been deleted or banned", "UserDeletedOrBanned"))
	}

	_, span2 := trace.StartSpan(ctx, "user.biz.login.gen-jwt")
	passHashed := business.hasher.Hash(data.Password + user.Salt)
	if user.Password != passHashed {
		span2.End()
		return nil, usermodel.ErrEmailOrPasswordInvalid
	}

	payload := tokenprovider.TokenPayload{
		UserId: user.Id,
		Role:   user.Role,
	}
	accessToken, err := business.tokenProvider.Generate(payload, business.expiry)
	span2.End()
	if err != nil {
		return nil, common.ErrInternal(err)
	}

	//refreshToken, err := business.tokenProvider.Generate(payload, business.tkCfg.GetRtExp())
	//if err != nil {
	//	return nil, common.ErrInternal(err)
	//}

	//account := usermodel.NewAccount(accessToken, refreshToken)

	return accessToken, nil
}
