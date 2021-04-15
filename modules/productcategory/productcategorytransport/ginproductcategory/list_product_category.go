package ginproductcategory

import (
	"net/http"
	"nhaancs/common"
	"nhaancs/component"
	"nhaancs/modules/productcategory/productcategorybiz"
	"nhaancs/modules/productcategory/productcategorymodel"
	"nhaancs/modules/productcategory/productcategorystore"

	"github.com/gin-gonic/gin"
)

func ListProductCategory(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter productcategorymodel.Filter
		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		paging.Fulfill()

		store := productcategorystore.NewSQLStore(appCtx.GetMainDBConnection())
		biz := productcategorybiz.NewListProductCategoryBiz(store)
		result, err := biz.ListProductCategory(c.Request.Context(), &filter, &paging)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, &paging, &filter))
	}
}
