package gincategory

import (
	"net/http"
	"nhaancs/common"
	"nhaancs/component"
	"nhaancs/modules/category/biz"
	"nhaancs/modules/category/store"

	"github.com/gin-gonic/gin"
)

func Delete(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := categorystore.NewSQLStore(appCtx.GetMainDBConnection())
		biz := categorybiz.NewDeleteBiz(store)
		if err := biz.Delete(c.Request.Context(), int(uid.GetLocalID())); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
