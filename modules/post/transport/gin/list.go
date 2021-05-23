package ginpost

import (
	"net/http"
	"nhaancs/common"
	"nhaancs/component"
	postbiz "nhaancs/modules/post/biz"
	postmodel "nhaancs/modules/post/model"
	postrepo "nhaancs/modules/post/repo"
	poststore "nhaancs/modules/post/store"

	"github.com/gin-gonic/gin"
)

func List(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter postmodel.Filter
		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		filter.Fulfill()

		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		paging.Fulfill()

		store := poststore.NewSQLStore(appCtx.GetMainDBConnection())
		// favoriteStore := favoritestore.NewSQLStore(appCtx.GetMainDBConnection())
		repo := postrepo.NewListRepo(store)
		biz := postbiz.NewListBiz(repo)
		result, err := biz.List(c.Request.Context(), &filter, &paging, common.IsRequesterAdmin(c))
		if err != nil {
			panic(err)
		}

		for i := range result {
			result[i].Mask(false)

			if i == len(result)-1 {
				paging.NextCursor = result[i].FakeId.String()
			}
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))
	}
}
