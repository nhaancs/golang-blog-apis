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
			c.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
			return
		}

		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
			return
		}
		paging.Fulfill()

		store := productcategorystore.NewSQLStore(appCtx.GetMainDBConnection())
		biz := productcategorybiz.NewListProductCategoryBiz(store)

		result, err := biz.ListProductCategory(c.Request.Context(), &filter, &paging)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"errors": err.Error()})
			return
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, &paging, &filter))
	}
}
