package middleware

import (
	"context"
	"errors"
	"nhaancs/common"
	"nhaancs/component"
	"nhaancs/component/tokenprovider/jwt"
	usermodel "nhaancs/modules/user/model"

	"github.com/gin-gonic/gin"
	"go.opencensus.io/trace"
)

type AuthenStore interface {
	FindUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
}

// 1. Get token from header
// 2. Validate token and parse to payload
// 3. From the token payload, we use user_id to find from DB
func RequiredAuth(appCtx component.AppContext, authStore AuthenStore) func(c *gin.Context) {
	tokenProvider := jwt.NewTokenJWTProvider(appCtx.SecretKey())
	return func(c *gin.Context) {
		token, err := common.ExtractTokenFromHeaderString(c.GetHeader("Authorization"))
		if err != nil {
			panic(err)
		}

		// db := appCtx.GetMainDBConnection()
		// store := userstorage.NewSQLStore(db)
		payload, err := tokenProvider.Validate(token)
		if err != nil {
			panic(err)
		}

		ctx, span := trace.StartSpan(c.Request.Context(), "middleware.RequiredAuth.find-user")
		user, err := authStore.FindUser(ctx, map[string]interface{}{"id": payload.UserId})
		span.End()
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
