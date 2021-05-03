package ginpost

import (
	"net/http"
	"nhaancs/common"
	"nhaancs/component"
	"nhaancs/modules/post/biz"
	"nhaancs/modules/post/store"

	"github.com/gin-gonic/gin"
)

func Delete(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := poststore.NewSQLStore(appCtx.GetMainDBConnection())
		biz := postbiz.NewDeleteBiz(store)
		if err := biz.Delete(c.Request.Context(), int(uid.GetLocalID())); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
