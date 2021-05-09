package favoritebiz

import (
	"context"
	"nhaancs/common"
	favoritemodel "nhaancs/modules/favorite/model"
	postmodel "nhaancs/modules/post/model"
)

type listFavoritedPostsOfAUserBiz struct {
	store ListStore
}

func NewListFavoritedPostsOfAUserBiz(store ListStore) *listFavoritedPostsOfAUserBiz {
	return &listFavoritedPostsOfAUserBiz{store: store}
}

func (biz *listFavoritedPostsOfAUserBiz) ListFavoritedPostsOfAUser(
	ctx context.Context,
	filter *favoritemodel.Filter,
	paging *common.Paging,
) ([]*postmodel.Post, error) {
	// result, err := biz.store.List(ctx, nil, filter, paging, "Post.Category", "Post.User")
	result, err := biz.store.List(ctx, nil, filter, paging, "Post")
	if err != nil {
		return nil, common.ErrCannotListEntity(favoritemodel.EntityName, err)
	}

	posts := make([]*postmodel.Post, len(result))
	for i, item := range result {
		posts[i] = item.Post
		posts[i].CreatedAt = item.CreatedAt
		posts[i].UpdatedAt = nil
	}

	return posts, nil
}
