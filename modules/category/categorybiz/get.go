package categorybiz

import (
	"context"
	"nhaancs/common"
	"nhaancs/modules/category/categorymodel"
)

type GetStore interface {
	Get(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*categorymodel.Category, error)
}

type getBiz struct {
	store GetStore
}

func NewGetBiz(store GetStore) *getBiz {
	return &getBiz{store: store}
}

func (biz *getBiz) Get(ctx context.Context, id int) (*categorymodel.Category, error) {
	data, err := biz.store.Get(ctx, map[string]interface{}{"id": id})

	if err != nil {
		if err != common.ErrRecordNotFound {
			return nil, common.ErrCannotGetEntity(categorymodel.EntityName, err)
		}

		return nil, common.ErrCannotGetEntity(categorymodel.EntityName, err)
	}
	if data.DeletedAt != nil {
		return nil, common.ErrEntityDeleted(categorymodel.EntityName, nil)
	}

	return data, nil
}
