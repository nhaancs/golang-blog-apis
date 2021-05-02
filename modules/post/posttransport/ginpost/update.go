package ginpost

import (
	"net/http"
	"nhaancs/common"
	"nhaancs/component"
	"nhaancs/modules/post/postbiz"
	"nhaancs/modules/post/postmodel"
	"nhaancs/modules/post/poststore"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Update(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		// uid, err := common.FromBase58(c.Param("id"))
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var data postmodel.PostUpdate
		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := poststore.NewSQLStore(appCtx.GetMainDBConnection())
		biz := postbiz.NewUpdateBiz(store)
		if err := biz.Update(c.Request.Context(), id, &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
