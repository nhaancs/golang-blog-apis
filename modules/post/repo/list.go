package postrepo

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

type listRepo struct {
	store ListStore
	// favoriteStore FavoriteStore
}

func NewListRepo(store ListStore) *listRepo {
	return &listRepo{store: store}
}

func (biz *listRepo) List(
	ctx context.Context,
	conditions map[string]interface{},
	filter *postmodel.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]postmodel.Post, error) {
	result, err := biz.store.List(ctx, conditions, filter, paging, "User", "Category")
	if err != nil {
		return nil, common.ErrCannotListEntity(postmodel.EntityName, err)
	}

	// ids := make([]int, len(result))
	// for i := range result {
	// 	ids[i] = result[i].Id
	// }
	// postFavoriteMap, _ := biz.favoriteStore.GetFavoriteCountsOfPosts(ctx, ids) // ignore error
	// if v := postFavoriteMap; v != nil {
	// 	for i := range result {
	// 		result[i].FavoriteCount = postFavoriteMap[result[i].Id]
	// 	}
	// }

	return result, nil
}
