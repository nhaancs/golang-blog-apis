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

func (biz *listBiz) ListFavoritedUsers(
	ctx context.Context,
	filter *favoritemodel.Filter,
	paging *common.Paging,
) ([]*common.SimpleUser, error) {
	result, err := biz.store.List(ctx, nil, filter, paging, "User")
	if err != nil {
		return nil, common.ErrCannotListEntity(favoritemodel.EntityName, err)
	}

	users := make([]*common.SimpleUser, len(result))
	for i, item := range result {
		users[i] = item.User
		users[i].CreatedAt = item.CreatedAt
		users[i].UpdatedAt = nil
	}

	return users, nil
}
