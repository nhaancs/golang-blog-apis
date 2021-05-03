package categorybiz

import (
	"context"
	"nhaancs/common"
	"nhaancs/modules/category/model"
)

type ListStore interface {
	List(ctx context.Context,
		conditions map[string]interface{},
		filter *categorymodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]categorymodel.Category, error)
}

type listBiz struct {
	store ListStore
}

func NewListBiz(store ListStore) *listBiz {
	return &listBiz{store: store}
}

func (biz *listBiz) List(
	ctx context.Context,
	filter *categorymodel.Filter,
	paging *common.Paging,
) ([]categorymodel.Category, error) {
	result, err := biz.store.List(ctx, nil, filter, paging)
	if err != nil {
		return nil, common.ErrCannotListEntity(categorymodel.EntityName, err)
	}

	return result, nil
}
