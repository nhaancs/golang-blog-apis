package subscriber

import (
	"context"
	"nhaancs/component"
	"nhaancs/pubsub"
	"time"
)

func RunDeleteAllPostsInACategoryGetDeleted(appCtx component.AppContext) subscribedJob {
	return subscribedJob{
		Title: "Delete all posts in a category get deleted",
		Handler: func(ctx context.Context, message *pubsub.Message) error {
			categoryId := message.Data().(int)
			db := appCtx.GetMainDBConnection()
			res := db.Exec("UPDATE posts SET deleted_at = ? WHERE category_id = ?", time.Now(), categoryId)
			// log error
			
			return res.Error
		},
	}
}
