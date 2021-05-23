package favoritebiz

import (
	"context"
	"nhaancs/common"
	"nhaancs/component/asyncjob"
	favoritemodel "nhaancs/modules/favorite/model"
	postmodel "nhaancs/modules/post/model"
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

	IncreaseFavoriteCount(
		ctx context.Context,
		postId int,
	) error
}

type favoriteBiz struct {
	store     FavoriteStore
	postStore PostStore
}

func NewFavoriteBiz(store FavoriteStore, postStore PostStore) *favoriteBiz {
	return &favoriteBiz{store: store, postStore: postStore}
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

	// side effect
	go func() {
		defer common.AppRecover()
		job := asyncjob.NewJob(func(ctx context.Context) error {
			return biz.postStore.IncreaseFavoriteCount(ctx, data.PostId)
		})
		_ = asyncjob.NewGroup(true, job).Run(ctx)
	}()

	return nil
}
