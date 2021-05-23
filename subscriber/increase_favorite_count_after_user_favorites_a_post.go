package subscriber

import (
	"context"
	"nhaancs/common"
	"nhaancs/component"
	"nhaancs/modules/post/store"
	"nhaancs/pubsub"
)

type HasPostId interface {
	GetPostId() int
}

func RunIncreaseFavoriteCountAfterUserFavoritesAPost(appCtx component.AppContext) subscribedJob {
	return subscribedJob{
		Title: "Increase favorite count after user favorites a post",
		Handler: func(ctx context.Context, message *pubsub.Message) error {
			defer common.AppRecover()
			store := poststore.NewSQLStore(appCtx.GetMainDBConnection())
			favoriteData := message.Data().(HasPostId)
			return store.IncreaseFavoriteCount(ctx, favoriteData.GetPostId())
		},
	}
}
