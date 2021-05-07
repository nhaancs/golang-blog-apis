package categorybiz

import (
	"context"
	"nhaancs/common"
	"nhaancs/modules/category/model"
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

func (biz *getBiz) Get(ctx context.Context, conditions map[string]interface{}, isAdmin bool) (*categorymodel.Category, error) {
	if conditions != nil && !isAdmin {
		conditions["is_enabled"] = true
	}
	
	data, err := biz.store.Get(ctx, conditions)
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
