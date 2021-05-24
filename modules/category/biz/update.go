package categorybiz

import (
	"context"
	"nhaancs/common"
	categorymodel "nhaancs/modules/category/model"
	"nhaancs/pubsub"
)

type UpdateStore interface {
	Get(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*categorymodel.Category, error)
	UpdateData(
		ctx context.Context,
		id int,
		data *categorymodel.CategoryUpdate,
	) error
}

type updateBiz struct {
	store  UpdateStore
	pubsub pubsub.Pubsub
}

func NewUpdateBiz(store UpdateStore, pubsub pubsub.Pubsub) *updateBiz {
	return &updateBiz{store: store, pubsub: pubsub}
}

func (biz *updateBiz) Update(ctx context.Context, id int, data *categorymodel.CategoryUpdate) error {
	if err := data.Validate(); err != nil {
		return err
	}

	oldData, err := biz.store.Get(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return common.ErrCannotGetEntity(categorymodel.EntityName, err)
	}
	if oldData.DeletedAt != nil {
		return common.ErrEntityDeleted(categorymodel.EntityName, nil)
	}

	if data.Slug != "" {
		_, err := biz.store.Get(ctx, map[string]interface{}{"slug": data.Slug})
		if data.Slug != oldData.Slug && err != common.ErrRecordNotFound {
			return common.ErrEntityExisted(categorymodel.EntityName, nil)
		}
	}

	if err := biz.store.UpdateData(ctx, id, data); err != nil {
		return common.ErrCannotUpdateEntity(categorymodel.EntityName, err)
	}

	if data.IsEnabled != nil && !*data.IsEnabled {
		// create a cron job to disable all posts in this category
		biz.pubsub.Publish(ctx, common.TopicCategoryDisabled, pubsub.NewMessage(id))
	}

	return nil
}
