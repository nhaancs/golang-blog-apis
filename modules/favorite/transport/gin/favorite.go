package ginfavorite

import (
	"net/http"
	"nhaancs/common"
	"nhaancs/component"
	favoritebiz "nhaancs/modules/favorite/biz"
	favoritemodel "nhaancs/modules/favorite/model"
	favoritestore "nhaancs/modules/favorite/store"
	poststore "nhaancs/modules/post/store"

	"github.com/gin-gonic/gin"
)

func Favorite(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data = new(favoritemodel.FavoriteCreate)
		fakePostId, err := common.FromBase58(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		requester := c.MustGet(common.CurrentUser).(common.Requester)
		data.UserId = requester.GetUserId()
		data.PostId = int(fakePostId.GetLocalID())

		store := favoritestore.NewSQLStore(appCtx.GetMainDBConnection())
		postStore := poststore.NewSQLStore(appCtx.GetMainDBConnection())
		biz := favoritebiz.NewFavoriteBiz(store, postStore)
		if err := biz.Favorite(c.Request.Context(), data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
