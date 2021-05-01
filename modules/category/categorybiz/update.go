package categorybiz

import (
	"context"
	"nhaancs/common"
	"nhaancs/modules/category/categorymodel"
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
	store UpdateStore
}

func NewUpdateBiz(store UpdateStore) *updateBiz {
	return &updateBiz{store: store}
}

func (biz *updateBiz) Update(ctx context.Context, id int, data *categorymodel.CategoryUpdate) error {
	oldData, err := biz.store.Get(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return common.ErrCannotGetEntity(categorymodel.EntityName, err)
	}

	if oldData.Status == 0 {
		return common.ErrEntityDeleted(categorymodel.EntityName, nil)
	}

	if err := biz.store.UpdateData(ctx, id, data); err != nil {
		return common.ErrCannotUpdateEntity(categorymodel.EntityName, err)
	}

	return nil
}
