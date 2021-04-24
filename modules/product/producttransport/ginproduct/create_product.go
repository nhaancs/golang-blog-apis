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

func CreateProduct(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data productmodel.ProductCreate
		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := productstore.NewSQLStore(appCtx.GetMainDBConnection())
		biz := productbiz.NewCreateProductBiz(store)
		if err := biz.CreateProduct(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))

	}
}
