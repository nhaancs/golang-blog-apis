package postbiz

import (
	"context"
	"nhaancs/common"
	"nhaancs/modules/post/model"
)

type DeleteStore interface {
	Get(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*postmodel.Post, error)
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
		return common.ErrCannotGetEntity(postmodel.EntityName, err)
	}
	if oldData.DeletedAt != nil {
		return common.ErrEntityDeleted(postmodel.EntityName, nil)
	}

	if err := biz.store.SoftDelete(ctx, id); err != nil {
		return common.ErrCannotDeleteEntity(postmodel.EntityName, err)
	}

	return nil
}
