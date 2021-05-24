package categorybiz

import (
	"context"
	"nhaancs/common"
	"nhaancs/modules/category/model"
	"nhaancs/pubsub"
)

type DeleteStore interface {
	Get(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*categorymodel.Category, error)
	SoftDelete(
		ctx context.Context,
		id int,
	) error
}

type deleteBiz struct {
	store  DeleteStore
	pubsub pubsub.Pubsub
}

func NewDeleteBiz(store DeleteStore, pubsub pubsub.Pubsub) *deleteBiz {
	return &deleteBiz{store: store, pubsub: pubsub}
}

func (biz *deleteBiz) Delete(ctx context.Context, id int) error {
	oldData, err := biz.store.Get(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return common.ErrCannotGetEntity(categorymodel.EntityName, err)
	}
	if oldData.DeletedAt != nil {
		return common.ErrEntityDeleted(categorymodel.EntityName, nil)
	}

	if err := biz.store.SoftDelete(ctx, id); err != nil {
		return common.ErrCannotDeleteEntity(categorymodel.EntityName, err)
	}

	// create a cron job to delete all posts belong to this category
	biz.pubsub.Publish(ctx, common.TopicCategoryDeleted, pubsub.NewMessage(id))

	return nil
}
