package gincategory

import (
	"net/http"
	"nhaancs/common"
	"nhaancs/component"
	categorybiz "nhaancs/modules/category/biz"
	categorystore "nhaancs/modules/category/store"

	"github.com/gin-gonic/gin"
)

// get by slug or id
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

		store := categorystore.NewSQLStore(appCtx.GetMainDBConnection())
		biz := categorybiz.NewGetBiz(store)
		data, err := biz.Get(c.Request.Context(), conditions, common.IsRequesterAdmin(c))
		if err != nil {
			panic(err)
		}

		data.Mask(false)
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
