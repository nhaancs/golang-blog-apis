package middleware

// todo: implement authorize or not middleware

import (
	"errors"
	"nhaancs/common"
	"nhaancs/component"
	"nhaancs/component/tokenprovider/jwt"
	userstorage "nhaancs/modules/user/store"

	"github.com/gin-gonic/gin"
)

// 1. Get token from header
// 2. Validate token and parse to payload
// 3. From the token payload, we use user_id to find from DB
func RequiredAuthOrNot(appCtx component.AppContext) func(c *gin.Context) {
	tokenProvider := jwt.NewTokenJWTProvider(appCtx.SecretKey())
	return func(c *gin.Context) {
		token, err := common.ExtractTokenFromHeaderString(c.GetHeader("Authorization"))
		if err != nil {
			panic(err)
		}

		db := appCtx.GetMainDBConnection()
		store := userstorage.NewSQLStore(db)
		payload, err := tokenProvider.Validate(token)
		if err != nil {
			panic(err)
		}

		user, err := store.FindUser(c.Request.Context(), map[string]interface{}{"id": payload.UserId})
		if err != nil {
			panic(err)
		}

		if user.DeletedAt != nil || !user.IsEnabled {
			panic(common.ErrNoPermission(errors.New("user has been deleted or banned")))
		}

		user.Mask(user.Role == common.AdminRole)
		c.Set(common.CurrentUser, user)
		c.Next()
	}
}
