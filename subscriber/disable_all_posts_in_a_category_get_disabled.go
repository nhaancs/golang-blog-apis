package subscriber

import (
	"context"
	"nhaancs/component"
	"nhaancs/pubsub"
)

func RunDisableAllPostsInACategoryGetDisabled(appCtx component.AppContext) subscribedJob {
	return subscribedJob{
		Title: "Disable all posts in a category get disabled",
		Handler: func(ctx context.Context, message *pubsub.Message) error {
			categoryId := message.Data().(int)
			db := appCtx.GetMainDBConnection()
			res := db.Exec("UPDATE posts SET is_enabled = ? WHERE category_id = ?", false, categoryId)
			// log error

			return res.Error
		},
	}
}
