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

func CreateProductCategory(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data productcategorymodel.ProductCategoryCreate
		if err := c.ShouldBind(&data); err != nil {
			// todo: standardize error responses
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		store := productcategorystore.NewSQLStore(appCtx.GetMainDBConnection())
		biz := productcategorybiz.NewCreateProductCategoryBiz(store)

		if err := biz.CreateProductCategory(c.Request.Context(), &data); err != nil {
			// todo: new way to return appropriate http status
			c.JSON(http.StatusInternalServerError, gin.H{"errors": err.Error()})
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))

	}
}
