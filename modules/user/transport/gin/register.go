package ginuser

import (
	"net/http"
	"nhaancs/common"
	"nhaancs/component"
	"nhaancs/component/hasher"
	"nhaancs/modules/user/biz"
	"nhaancs/modules/user/model"
	"nhaancs/modules/user/store"

	"github.com/gin-gonic/gin"
)

func Register(appCtx component.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()
		var data usermodel.UserCreate
		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		store := userstorage.NewSQLStore(db)
		md5 := hasher.NewMd5Hash()
		biz := userbiz.NewRegisterBusiness(store, md5)
		if err := biz.Register(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
