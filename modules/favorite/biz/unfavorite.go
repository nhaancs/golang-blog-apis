package favoritebiz

import (
	"context"
	"nhaancs/common"
	favoritemodel "nhaancs/modules/favorite/model"
	"nhaancs/pubsub"
)

type UnfavoriteStore interface {
	Delete(
		ctx context.Context,
		userId int,
		postId int,
	) error
}

type unfavoriteBiz struct {
	store  UnfavoriteStore
	pubsub pubsub.Pubsub
}

func NewUnfavoriteBiz(store UnfavoriteStore, pubsub pubsub.Pubsub) *unfavoriteBiz {
	return &unfavoriteBiz{store: store, pubsub: pubsub}
}

func (biz *unfavoriteBiz) Unfavorite(ctx context.Context, userId int, postId int) error {
	// delete don't return error if entity not found
	if err := biz.store.Delete(ctx, userId, postId); err != nil {
		return common.ErrCannotDeleteEntity(favoritemodel.EntityName, err)
	}

	biz.pubsub.Publish(ctx, common.TopicUserUnfavoritePost, pubsub.NewMessage(postId))

	return nil
}
