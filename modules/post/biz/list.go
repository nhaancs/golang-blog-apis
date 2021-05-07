package postbiz

import (
	"context"
	"nhaancs/common"
	"nhaancs/modules/post/model"
)

type ListStore interface {
	List(
		ctx context.Context,
		conditions map[string]interface{},
		filter *postmodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]postmodel.Post, error)
}

type FavoriteStore interface {
	GetFavoriteCountsOfPosts(
		ctx context.Context,
		postIds []int,
	) (map[int]int, error)
}

type listBiz struct {
	store         ListStore
	favoriteStore FavoriteStore
}

func NewListBiz(store ListStore, favoriteStore FavoriteStore) *listBiz {
	return &listBiz{store: store, favoriteStore: favoriteStore}
}

func (biz *listBiz) List(
	ctx context.Context,
	filter *postmodel.Filter,
	paging *common.Paging,
	isAdmin bool,
) ([]postmodel.Post, error) {
	conditions := map[string]interface{}{}
	if !isAdmin {
		conditions["is_enabled"] = true
	}
	result, err := biz.store.List(ctx, conditions, filter, paging)
	if err != nil {
		return nil, common.ErrCannotListEntity(postmodel.EntityName, err)
	}

	ids := make([]int, len(result))
	for i := range result {
		ids[i] = result[i].Id
	}
	postFavoriteMap, _ := biz.favoriteStore.GetFavoriteCountsOfPosts(ctx, ids) // ignore error
	if v := postFavoriteMap; v != nil {
		for i := range result {
			result[i].FavoriteCount = postFavoriteMap[result[i].Id]
		}
	}

	return result, nil
}
