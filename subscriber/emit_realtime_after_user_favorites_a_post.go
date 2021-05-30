package subscriber

import (
	"context"
	"nhaancs/component"
	"nhaancs/pubsub"
	"nhaancs/socket"
)

type HasUserId interface {
	GetUserId() int
}

// emit to author of the post
func EmitRealtimeAfterUserFavoritesAPost(appCtx component.AppContext, rtEngine socket.RealtimeEngine) subscribedJob {
	return subscribedJob{
		Title: "Emit realtime after user favorites a post",
		Handler: func(ctx context.Context, message *pubsub.Message) error {
			data := message.Data().(HasUserId)
			return rtEngine.EmitToUser(data.GetUserId(), string(message.Topic()), data)
		},
	}
}