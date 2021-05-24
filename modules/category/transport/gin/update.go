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

func Update(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var data categorymodel.CategoryUpdate
		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := categorystore.NewSQLStore(appCtx.GetMainDBConnection())
		biz := categorybiz.NewUpdateBiz(store, appCtx.GetPubsub())
		if err := biz.Update(c.Request.Context(), int(uid.GetLocalID()), &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
