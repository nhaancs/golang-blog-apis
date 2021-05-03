package gincategory

import (
	"net/http"
	"nhaancs/common"
	"nhaancs/component"
	"nhaancs/modules/category/biz"
	"nhaancs/modules/category/model"
	"nhaancs/modules/category/store"

	"github.com/gin-gonic/gin"
)

func Create(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data categorymodel.CategoryCreate
		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		// requester := c.MustGet(common.CurrentUser).(common.Requester)
		// data.OwnerId = requester.GetUserId()

		store := categorystore.NewSQLStore(appCtx.GetMainDBConnection())
		biz := categorybiz.NewCreateBiz(store)
		if err := biz.Create(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		data.Mask(false)
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeId.String()))
	}
}
