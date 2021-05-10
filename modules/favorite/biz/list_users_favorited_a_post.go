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

type listUsersFavoritedAPostBiz struct {
	store ListStore
}

func NewListUsersFavoritedAPostBiz(store ListStore) *listUsersFavoritedAPostBiz {
	return &listUsersFavoritedAPostBiz{store: store}
}

func (biz *listUsersFavoritedAPostBiz) ListUsersFavoritedAPost(
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
		// preload in gorm is actually the second query, not a join statememt,
		// so item.User can be nil
		if item.User == nil {
			continue
		}
		users[i] = item.User
		users[i].CreatedAt = item.CreatedAt
		users[i].UpdatedAt = nil
	}

	return users, nil
}
