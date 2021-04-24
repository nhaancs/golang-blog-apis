package ginproduct

import (
	"net/http"
	"nhaancs/common"
	"nhaancs/component"
	"nhaancs/modules/product/productbiz"
	"nhaancs/modules/product/productstore"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DeleteProduct(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := productstore.NewSQLStore(appCtx.GetMainDBConnection())
		biz := productbiz.NewDeleteProductBiz(store)
		if err := biz.DeleteProduct(c.Request.Context(), id); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
