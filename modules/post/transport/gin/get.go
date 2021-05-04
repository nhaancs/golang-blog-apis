package ginpost

import (
	"net/http"
	"nhaancs/common"
	"nhaancs/component"
	favoritestore "nhaancs/modules/favorite/store"
	"nhaancs/modules/post/biz"
	"nhaancs/modules/post/store"

	"github.com/gin-gonic/gin"
)

func Get(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := poststore.NewSQLStore(appCtx.GetMainDBConnection())
		favoriteStore := favoritestore.NewSQLStore(appCtx.GetMainDBConnection())
		biz := postbiz.NewGetBiz(store, favoriteStore)
		data, err := biz.Get(c.Request.Context(), int(uid.GetLocalID()))
		if err != nil {
			panic(err)
		}

		data.Mask(false)
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
