package gincategory

import (
	"net/http"
	"nhaancs/common"
	"nhaancs/component"
	categorybiz "nhaancs/modules/category/biz"
	categorymodel "nhaancs/modules/category/model"
	categorystore "nhaancs/modules/category/store"

	"github.com/gin-gonic/gin"
)

func List(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter categorymodel.Filter
		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		filter.Fulfill()

		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		paging.Fulfill()

		store := categorystore.NewSQLStore(appCtx.GetMainDBConnection())
		biz := categorybiz.NewListBiz(store)
		result, err := biz.List(c.Request.Context(), &filter, &paging)
		if err != nil {
			panic(err)
		}

		for i := range result {
			// todo: what if user and admin use the same list api
			result[i].Mask(false)

			if i == len(result)-1 {
				paging.NextCursor = result[i].FakeId.String()
			}
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))
	}
}
