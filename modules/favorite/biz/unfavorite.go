package favoritebiz

import (
	"context"
	"nhaancs/common"
	favoritemodel "nhaancs/modules/favorite/model"
)

type UnfavoriteStore interface {
	Get(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*favoritemodel.Favorite, error)
	Delete(
		ctx context.Context,
		userId int,
		postId int,
	) error
}

type unfavoriteBiz struct {
	store UnfavoriteStore
}

func NewUnfavoriteBiz(store UnfavoriteStore) *unfavoriteBiz {
	return &unfavoriteBiz{store: store}
}

func (biz *unfavoriteBiz) Unfavorite(ctx context.Context, userId int, postId int) error {
	if _, err := biz.store.Get(ctx, map[string]interface{}{"post_id": postId, "user_id": userId}); err != nil {
		return nil
	}
	if err := biz.store.Delete(ctx, userId, postId); err != nil {
		return common.ErrCannotDeleteEntity(favoritemodel.EntityName, err)
	}

	return nil
}
