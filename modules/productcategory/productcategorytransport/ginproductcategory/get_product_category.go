package ginproductcategory

import (
	"net/http"
	"nhaancs/common"
	"nhaancs/component"
	"nhaancs/modules/productcategory/productcategorybiz"
	"nhaancs/modules/productcategory/productcategorystore"

	"github.com/gin-gonic/gin"
)

func GetProductCategoryBySlug(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		slug := c.Param("slug")

		store := productcategorystore.NewSQLStore(appCtx.GetMainDBConnection())
		biz := productcategorybiz.NewGetProductCategoryBiz(store)
		data, err := biz.GetProductCategoryBySlug(c.Request.Context(), slug)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}