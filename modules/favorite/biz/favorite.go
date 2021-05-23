package favoritebiz

import (
	"context"
	"nhaancs/common"
	favoritemodel "nhaancs/modules/favorite/model"
	postmodel "nhaancs/modules/post/model"
	"nhaancs/pubsub"
)

type FavoriteStore interface {
	Create(ctx context.Context, data *favoritemodel.FavoriteCreate) error
	Get(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*favoritemodel.Favorite, error)
}

type PostStore interface {
	Get(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*postmodel.Post, error)
}

type favoriteBiz struct {
	store     FavoriteStore
	postStore PostStore
	pubsub    pubsub.Pubsub
}

func NewFavoriteBiz(store FavoriteStore, postStore PostStore, pubsub pubsub.Pubsub) *favoriteBiz {
	return &favoriteBiz{store: store, postStore: postStore, pubsub: pubsub}
}

func (biz *favoriteBiz) Favorite(ctx context.Context, data *favoritemodel.FavoriteCreate) error {
	if err := data.Validate(); err != nil {
		return err
	}

	{
		_, err := biz.store.Get(ctx, map[string]interface{}{"user_id": data.UserId, "post_id": data.PostId})
		if err != common.ErrRecordNotFound {
			return favoritemodel.ErrFavoriteAPostTwice
		}
	}

	{
		post, err := biz.postStore.Get(ctx, map[string]interface{}{"id": data.PostId})
		if err != nil || !post.IsEnabled || post.DeletedAt != nil {
			return favoritemodel.ErrFavoritePostIsInvalid(err)
		}
	}

	if err := biz.store.Create(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(favoritemodel.EntityName, err)
	}

	biz.pubsub.Publish(ctx, common.TopicUserFavoritePost, pubsub.NewMessage(data))

	return nil
}
