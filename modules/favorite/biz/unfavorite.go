package favoritebiz

import (
	"context"
	"nhaancs/common"
	"nhaancs/component/asyncjob"
	favoritemodel "nhaancs/modules/favorite/model"
)

type UnfavoriteStore interface {
	// Get(
	// 	ctx context.Context,
	// 	conditions map[string]interface{},
	// 	moreKeys ...string,
	// ) (*favoritemodel.Favorite, error)
	Delete(
		ctx context.Context,
		userId int,
		postId int,
	) error
}

type DecreaseFavoriteCountStore interface {
	DecreaseFavoriteCount(
		ctx context.Context,
		postId int,
	) error
}

type unfavoriteBiz struct {
	store    UnfavoriteStore
	decStore DecreaseFavoriteCountStore
}

func NewUnfavoriteBiz(store UnfavoriteStore, decStore DecreaseFavoriteCountStore) *unfavoriteBiz {
	return &unfavoriteBiz{store: store, decStore: decStore}
}

func (biz *unfavoriteBiz) Unfavorite(ctx context.Context, userId int, postId int) error {
	// if _, err := biz.store.Get(ctx, map[string]interface{}{"post_id": postId, "user_id": userId}); err != nil {
	// 	return nil
	// }

	// delete dont return error if entity not found
	if err := biz.store.Delete(ctx, userId, postId); err != nil {
		return common.ErrCannotDeleteEntity(favoritemodel.EntityName, err)
	}

	// side effect
	go func() {
		defer common.AppRecover()
		job := asyncjob.NewJob(func(ctx context.Context) error {
			return biz.decStore.DecreaseFavoriteCount(ctx, postId)
		})
		_ = asyncjob.NewGroup(true, job).Run(ctx)
	}()

	return nil
}
