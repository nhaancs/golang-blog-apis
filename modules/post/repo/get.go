package postrepo

import (
	"context"
	"nhaancs/common"
	"nhaancs/modules/post/model"
)

type GetStore interface {
	Get(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*postmodel.Post, error)
}

type FavoriteStore interface {
	GetFavoriteCountsOfPosts(
		ctx context.Context,
		postIds []int,
	) (map[int]int, error)
}

type getRepo struct {
	store         GetStore
	favoriteStore FavoriteStore
}

func NewGetRepo(store GetStore, favoriteStore FavoriteStore) *getRepo {
	return &getRepo{store: store, favoriteStore: favoriteStore}
}

func (biz *getRepo) Get(ctx context.Context, conditions map[string]interface{}, isAdmin bool) (*postmodel.Post, error) {
	data, err := biz.store.Get(ctx, conditions, "User", "Category")
	if err != nil {
		if err != common.ErrRecordNotFound {
			return nil, common.ErrCannotGetEntity(postmodel.EntityName, err)
		}

		return nil, common.ErrCannotGetEntity(postmodel.EntityName, err)
	}
	if data.DeletedAt != nil {
		return nil, common.ErrEntityDeleted(postmodel.EntityName, nil)
	}

	ids := []int{data.Id}
	postFavoriteMap, _ := biz.favoriteStore.GetFavoriteCountsOfPosts(ctx, ids) // ignore error
	if v := postFavoriteMap; v != nil {
		data.FavoriteCount = postFavoriteMap[data.Id]
	}

	return data, nil
}
