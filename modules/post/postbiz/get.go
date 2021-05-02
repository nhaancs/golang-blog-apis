package postbiz

import (
	"context"
	"nhaancs/common"
	"nhaancs/modules/post/postmodel"
)

type GetStore interface {
	Get(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*postmodel.Post, error)
}

type getBiz struct {
	store GetStore
}

func NewGetBiz(store GetStore) *getBiz {
	return &getBiz{store: store}
}

func (biz *getBiz) Get(ctx context.Context, id int) (*postmodel.Post, error) {
	data, err := biz.store.Get(ctx, map[string]interface{}{"id": id})

	if err != nil {
		if err != common.ErrRecordNotFound {
			return nil, common.ErrCannotGetEntity(postmodel.EntityName, err)
		}

		return nil, common.ErrCannotGetEntity(postmodel.EntityName, err)
	}
	if data.DeletedAt != nil {
		return nil, common.ErrEntityDeleted(postmodel.EntityName, nil)
	}

	return data, err
}
