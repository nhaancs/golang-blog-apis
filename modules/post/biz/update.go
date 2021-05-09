package postbiz

import (
	"context"
	"nhaancs/common"
	categorymodel "nhaancs/modules/category/model"
	"nhaancs/modules/post/model"
)

type UpdateStore interface {
	Get(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*postmodel.Post, error)
	UpdateData(
		ctx context.Context,
		id int,
		data *postmodel.PostUpdate,
	) error
}

type updateBiz struct {
	store         UpdateStore
	categoryStore CategoryStore
}

func NewUpdateBiz(store UpdateStore, categoryStore CategoryStore) *updateBiz {
	return &updateBiz{store: store, categoryStore: categoryStore}
}

func (biz *updateBiz) Update(ctx context.Context, id int, data *postmodel.PostUpdate) error {
	if err := data.Validate(); err != nil {
		return err
	}

	oldData, err := biz.store.Get(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return common.ErrCannotGetEntity(postmodel.EntityName, err)
	}
	if oldData.DeletedAt != nil {
		return common.ErrEntityDeleted(postmodel.EntityName, nil)
	}

	if data.Slug != "" {
		_, err := biz.store.Get(ctx, map[string]interface{}{"slug": data.Slug})
		if data.Slug != oldData.Slug && err != common.ErrRecordNotFound {
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

	if err := biz.store.UpdateData(ctx, id, data); err != nil {
		return common.ErrCannotUpdateEntity(postmodel.EntityName, err)
	}
	return nil
}
