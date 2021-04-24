package ginproduct

import (
	"net/http"
	"nhaancs/common"
	"nhaancs/component"
	"nhaancs/modules/product/productbiz"
	"nhaancs/modules/product/productmodel"
	"nhaancs/modules/product/productstore"

	"github.com/gin-gonic/gin"
)

func ListProduct(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter productmodel.Filter
		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		paging.Fulfill()

		store := productstore.NewSQLStore(appCtx.GetMainDBConnection())
		biz := productbiz.NewListProductBiz(store)
		result, err := biz.ListProduct(c.Request.Context(), &filter, &paging)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, &paging, &filter))
	}
}
