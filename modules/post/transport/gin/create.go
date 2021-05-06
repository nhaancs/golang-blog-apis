package ginpost

import (
	"net/http"
	"nhaancs/common"
	"nhaancs/component"
	postbiz "nhaancs/modules/post/biz"
	postmodel "nhaancs/modules/post/model"
	poststore "nhaancs/modules/post/store"

	"github.com/gin-gonic/gin"
)

func Create(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data = new(postmodel.PostCreate)
		if err := c.ShouldBind(data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		data.Fulfill()
		requester := c.MustGet(common.CurrentUser).(common.Requester)
		data.UserId = requester.GetUserId()

		store := poststore.NewSQLStore(appCtx.GetMainDBConnection())
		biz := postbiz.NewCreateBiz(store)
		if err := biz.Create(c.Request.Context(), data); err != nil {
			panic(err)
		}

		data.Mask(false)
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeId.String()))
	}
}
