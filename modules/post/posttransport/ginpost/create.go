package ginpost

import (
	"net/http"
	"nhaancs/common"
	"nhaancs/component"
	"nhaancs/modules/post/postbiz"
	"nhaancs/modules/post/postmodel"
	"nhaancs/modules/post/poststore"

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

		// data.GenUID(common.DbTypePost)

		// c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeId.String()))
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.Id))
	}
}
