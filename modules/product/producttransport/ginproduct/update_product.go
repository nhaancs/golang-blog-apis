package ginproduct

import (
	"net/http"
	"nhaancs/common"
	"nhaancs/component"
	"nhaancs/modules/product/productbiz"
	"nhaancs/modules/product/productmodel"
	"nhaancs/modules/product/productstore"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UpdateProduct(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var data productmodel.ProductUpdate
		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := productstore.NewSQLStore(appCtx.GetMainDBConnection())
		biz := productbiz.NewUpdateProductBiz(store)
		if err := biz.UpdateProduct(c.Request.Context(), id, &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
