package postbiz

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

type getBiz struct {
	store         GetStore
	favoriteStore FavoriteStore
}

func NewGetBiz(store GetStore, favoriteStore FavoriteStore) *getBiz {
	return &getBiz{store: store, favoriteStore: favoriteStore}
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

	ids := []int{data.Id}
	postFavoriteMap, _ := biz.favoriteStore.GetFavoriteCountsOfPosts(ctx, ids) // ignore error
	if v := postFavoriteMap; v != nil {
		data.FavoriteCount = postFavoriteMap[data.Id]
	}

	return data, nil
}
