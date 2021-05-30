package subscriber

import (
	"context"
	"nhaancs/component"
	poststore "nhaancs/modules/post/store"
	"nhaancs/pubsub"
)

type HasPostId interface {
	GetPostId() int
}


func RunIncreaseFavoriteCountAfterUserFavoritesAPost(appCtx component.AppContext) subscribedJob {
	return subscribedJob{
		Title: "Increase favorite count after user favorites a post",
		Handler: func(ctx context.Context, message *pubsub.Message) error {
			store := poststore.NewSQLStore(appCtx.GetMainDBConnection())
			favoriteData := message.Data().(HasPostId)
			return store.IncreaseFavoriteCount(ctx, favoriteData.GetPostId())
		},
	}
}
