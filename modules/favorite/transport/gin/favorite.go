package ginfavorite

import (
	"net/http"
	"nhaancs/common"
	"nhaancs/component"
	favoritebiz "nhaancs/modules/favorite/biz"
	favoritemodel "nhaancs/modules/favorite/model"
	favoritestore "nhaancs/modules/favorite/store"

	"github.com/gin-gonic/gin"
)

func Favorite(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data = new(favoritemodel.FavoriteCreate)
		if err := c.ShouldBind(data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		data.Fulfill()
		//todo: get user id
		data.UserId = 1

		// requester := c.MustGet(common.CurrentUser).(common.Requester)
		// data.OwnerId = requester.GetUserId()

		store := favoritestore.NewSQLStore(appCtx.GetMainDBConnection())
		biz := favoritebiz.NewFavoriteBiz(store)
		if err := biz.Favorite(c.Request.Context(), data); err != nil {
			panic(err)
		}
		
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
