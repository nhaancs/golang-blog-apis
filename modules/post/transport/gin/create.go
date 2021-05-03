package ginpost

import (
	"net/http"
	"nhaancs/common"
	"nhaancs/component"
	"nhaancs/modules/post/biz"
	"nhaancs/modules/post/model"
	"nhaancs/modules/post/store"

	"github.com/gin-gonic/gin"
)

func Create(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data postmodel.PostCreate
		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		// requester := c.MustGet(common.CurrentUser).(common.Requester)
		// data.OwnerId = requester.GetUserId()

		store := poststore.NewSQLStore(appCtx.GetMainDBConnection())
		biz := postbiz.NewCreateBiz(store)
		if err := biz.Create(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		data.Mask(false)
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeId.String()))
	}
}
