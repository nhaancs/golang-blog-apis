package middleware

import (
	"nhaancs/common"
	"nhaancs/component"
	"nhaancs/component/tokenprovider/jwt"
	usermodel "nhaancs/modules/user/model"

	"github.com/gin-gonic/gin"
)

func RequiredAuthOrNot(appCtx component.AppContext) func(c *gin.Context) {
	tokenProvider := jwt.NewTokenJWTProvider(appCtx.SecretKey())
	return func(c *gin.Context) {
		token, err := common.ExtractTokenFromHeaderString(c.GetHeader("Authorization"))
		if err != nil {
			c.Next()
			return
		}

		payload, err := tokenProvider.Validate(token)
		if err != nil {
			c.Next()
			return
		}

		user := &usermodel.User{}
		user.Id = payload.UserId
		user.Role = payload.Role
		user.Mask(user.Role == common.AdminRole)
		c.Set(common.CurrentUser, user)
		c.Next()
	}
}
