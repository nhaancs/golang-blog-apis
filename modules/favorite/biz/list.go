package favoritebiz

import (
	"context"
	"nhaancs/common"
	favoritemodel "nhaancs/modules/favorite/model"
)

type ListStore interface {
	List(ctx context.Context,
		conditions map[string]interface{},
		filter *favoritemodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]favoritemodel.Favorite, error)
}

type listBiz struct {
	store ListStore
}

func NewListBiz(store ListStore) *listBiz {
	return &listBiz{store: store}
}

func (biz *listBiz) List(
	ctx context.Context,
	filter *favoritemodel.Filter,
	paging *common.Paging,
) ([]favoritemodel.Favorite, error) {
	result, err := biz.store.List(ctx, nil, filter, paging, "User", "Post")
	if err != nil {
		return nil, common.ErrCannotListEntity(favoritemodel.EntityName, err)
	}

	return result, nil
}
