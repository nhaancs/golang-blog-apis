package ginfavorite

import (
	"net/http"
	"nhaancs/common"
	"nhaancs/component"
	favoritebiz "nhaancs/modules/favorite/biz"
	favoritestore "nhaancs/modules/favorite/store"

	"github.com/gin-gonic/gin"
)

func Unfavorite(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		postId, err := common.FromBase58(c.Param("postId"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		// todo: get user id 
		userId := 1
		store := favoritestore.NewSQLStore(appCtx.GetMainDBConnection())
		biz := favoritebiz.NewUnfavoriteBiz(store)
		if err := biz.Unfavorite(c.Request.Context(), userId, int(postId.GetLocalID())); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
