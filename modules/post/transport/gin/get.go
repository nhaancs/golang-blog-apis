package ginpost

import (
	"net/http"
	"nhaancs/common"
	"nhaancs/component"
	favoritestore "nhaancs/modules/favorite/store"
	"nhaancs/modules/post/biz"
	"nhaancs/modules/post/store"

	"github.com/gin-gonic/gin"
)

func Get(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		conditions := map[string]interface{}{}
		id := c.Param("id")
		slug := c.Param("slug")

		if id != "" && slug == "" { // find by id
			uid, err := common.FromBase58(c.Param("id"))
			if err != nil {
				panic(common.ErrInvalidRequest(err))
			}
			conditions["id"] = int(uid.GetLocalID())
		} else if id == "" && slug != "" { // find by slug
			conditions["slug"] = slug
		} else {
			panic(common.ErrInvalidRequest(nil))
		}


		store := poststore.NewSQLStore(appCtx.GetMainDBConnection())
		favoriteStore := favoritestore.NewSQLStore(appCtx.GetMainDBConnection())
		biz := postbiz.NewGetBiz(store, favoriteStore)
		data, err := biz.Get(c.Request.Context(), conditions, common.IsRequesterAdmin(c))
		if err != nil {
			panic(err)
		}

		data.Mask(false)
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
