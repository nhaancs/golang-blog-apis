package postbiz

import (
	"context"
	"nhaancs/common"
	categorymodel "nhaancs/modules/category/model"
	"nhaancs/modules/post/model"
)

type CreateStore interface {
	Create(ctx context.Context, data *postmodel.PostCreate) error
	Get(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*postmodel.Post, error)
}
type CategoryStore interface {
	Get(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*categorymodel.Category, error)
}

type createBiz struct {
	store CreateStore
	categoryStore CategoryStore
}

func NewCreateBiz(store CreateStore, categoryStore CategoryStore) *createBiz {
	return &createBiz{store: store, categoryStore: categoryStore}
}

func (biz *createBiz) Create(ctx context.Context, data *postmodel.PostCreate) error {
	if err := data.Validate(); err != nil {
		return err
	}
	{
		_, err := biz.store.Get(ctx, map[string]interface{}{"slug": data.Slug})
		if err != common.ErrRecordNotFound {
			return common.ErrEntityExisted(postmodel.EntityName, nil)
		}
	}
	{
		cat, err := biz.categoryStore.Get(ctx, map[string]interface{}{"id": data.CategoryId})
		if cat == nil {
			return common.ErrEntityNotFound(categorymodel.EntityName, err)
		}
		if cat.DeletedAt != nil || !cat.IsEnabled {
			return common.ErrEntityDeleted(categorymodel.EntityName, nil)
		}
	}

	if err := biz.store.Create(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(postmodel.EntityName, err)
	}

	return nil
}
