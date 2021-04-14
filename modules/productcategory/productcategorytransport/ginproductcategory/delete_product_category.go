package ginproductcategory

import (
	"net/http"
	"nhaancs/common"
	"nhaancs/component"
	"nhaancs/modules/productcategory/productcategorybiz"
	"nhaancs/modules/productcategory/productcategorystore"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DeleteProductCategory(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"errors": err.Error()})
			return
		}

		store := productcategorystore.NewSQLStore(appCtx.GetMainDBConnection())
		biz := productcategorybiz.NewDeleteProductCategoryBiz(store)
		if err := biz.DeleteProductCategory(c.Request.Context(), id); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"errors": err.Error()})
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
