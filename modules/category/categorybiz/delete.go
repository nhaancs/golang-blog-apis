package categorybiz

import (
	"context"
	"nhaancs/common"
	"nhaancs/modules/category/categorymodel"
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
	store DeleteStore
}

func NewDeleteBiz(store DeleteStore) *deleteBiz {
	return &deleteBiz{store: store}
}

func (biz *deleteBiz) Delete(ctx context.Context, id int) error {
	oldData, err := biz.store.Get(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return common.ErrCannotGetEntity(categorymodel.EntityName, err)
	}

	if oldData.Status == 0 {
		return common.ErrEntityDeleted(categorymodel.EntityName, nil)
	}

	if err := biz.store.SoftDelete(ctx, id); err != nil {
		return common.ErrCannotDeleteEntity(categorymodel.EntityName, err)
	}

	return nil
}
