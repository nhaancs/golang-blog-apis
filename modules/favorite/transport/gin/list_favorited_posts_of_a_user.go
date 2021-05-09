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

func ListFavoritedPostsOfAUser(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter favoritemodel.Filter
		requester := c.MustGet(common.CurrentUser).(common.Requester)
		filter.UserId = requester.GetUserId()

		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		paging.Fulfill()

		store := favoritestore.NewSQLStore(appCtx.GetMainDBConnection())
		biz := favoritebiz.NewListFavoritedPostsOfAUserBiz(store)
		result, err := biz.ListFavoritedPostsOfAUser(c.Request.Context(), &filter, &paging)
		if err != nil {
			panic(err)
		}

		for i := range result {
			result[i].Mask(false)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))
	}
}
