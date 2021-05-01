package gincategory

import (
	"net/http"
	"nhaancs/common"
	"nhaancs/component"
	"nhaancs/modules/category/categorybiz"
	"nhaancs/modules/category/categorystore"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Get(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		// uid, err := common.FromBase58(c.Param("id"))
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := categorystore.NewSQLStore(appCtx.GetMainDBConnection())
		biz := categorybiz.NewGetBiz(store)
		// data, err := biz.Get(c.Request.Context(), int(uid.GetLocalID()))
		data, err := biz.Get(c.Request.Context(), id)
		if err != nil {
			panic(err)
		}

		// data.Mask(false)
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
