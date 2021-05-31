package subscriber

import (
	"context"
	"nhaancs/component"
	"nhaancs/pubsub"
	"nhaancs/socketengine"
)

type HasUserId interface {
	GetUserId() int
}

// emit to the user favorites the post
func EmitRealtimeAfterUserFavoritesAPost(appCtx component.AppContext, rtEngine socketengine.RealtimeEngine) subscribedJob {
	return subscribedJob{
		Title: "Emit realtime after user favorites a post",
		Handler: func(ctx context.Context, message *pubsub.Message) error {
			data := message.Data().(HasUserId)
			return rtEngine.EmitToUser(data.GetUserId(), string(message.Topic()), data)
		},
	}
}
