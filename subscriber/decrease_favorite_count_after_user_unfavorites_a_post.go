package subscriber

import (
	"context"
	"nhaancs/component"
	poststore "nhaancs/modules/post/store"
	"nhaancs/pubsub"
)

func RunDecreaseUnfavoriteCountAfterUserFavoritesAPost(appCtx component.AppContext) subscribedJob {
	return subscribedJob{
		Title: "Decrease favorite count after user unfavorites a post",
		Handler: func(ctx context.Context, message *pubsub.Message) error {
			store := poststore.NewSQLStore(appCtx.GetMainDBConnection())
			postId := message.Data().(int)
			return store.DecreaseFavoriteCount(ctx, postId)
		},
	}
}
