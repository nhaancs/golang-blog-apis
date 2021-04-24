package ginproduct

import (
	"net/http"
	"nhaancs/common"
	"nhaancs/component"
	"nhaancs/modules/product/productbiz"
	"nhaancs/modules/product/productstore"

	"github.com/gin-gonic/gin"
)

func GetProductBySlug(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		slug := c.Param("slug")

		store := productstore.NewSQLStore(appCtx.GetMainDBConnection())
		biz := productbiz.NewGetProductBiz(store)
		data, err := biz.GetProductBySlug(c.Request.Context(), slug)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}