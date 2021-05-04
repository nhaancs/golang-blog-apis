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

func Update(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var data = new(postmodel.PostUpdate)
		if err := c.ShouldBind(data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		data.Fulfill()

		store := poststore.NewSQLStore(appCtx.GetMainDBConnection())
		biz := postbiz.NewUpdateBiz(store)
		if err := biz.Update(c.Request.Context(), int(uid.GetLocalID()), data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
