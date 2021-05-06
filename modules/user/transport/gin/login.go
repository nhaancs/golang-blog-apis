package ginuser

import (
	"net/http"
	"nhaancs/common"
	"nhaancs/component"
	"nhaancs/component/hasher"
	"nhaancs/component/tokenprovider/jwt"
	"nhaancs/modules/user/biz"
	"nhaancs/modules/user/model"
	"nhaancs/modules/user/store"

	"github.com/gin-gonic/gin"
)

func Login(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var loginUserData usermodel.UserLogin
		if err := c.ShouldBind(&loginUserData); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := appCtx.GetMainDBConnection()
		tokenProvider := jwt.NewTokenJWTProvider(appCtx.SecretKey())
		store := userstorage.NewSQLStore(db)
		md5 := hasher.NewMd5Hash()
		business := userbiz.NewLoginBusiness(store, tokenProvider, md5, 60*60*24*30)
		account, err := business.Login(c.Request.Context(), &loginUserData)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(account))
	}
}
