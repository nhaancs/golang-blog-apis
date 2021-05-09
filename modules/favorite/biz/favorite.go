package favoritebiz

import (
	"context"
	"nhaancs/common"
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

type GetPostStore interface {
	Get(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*postmodel.Post, error)
}

type favoriteBiz struct {
	store FavoriteStore
	getPostStore GetPostStore
}

func NewFavoriteBiz(store FavoriteStore, getPostStore GetPostStore) *favoriteBiz {
	return &favoriteBiz{store: store, getPostStore: getPostStore}
}

func (biz *favoriteBiz) Favorite(ctx context.Context, data *favoritemodel.FavoriteCreate) error {
	if err := data.Validate(); err != nil {
		return err
	}

	{
		_, err := biz.store.Get(ctx, map[string]interface{}{"user_id": data.UserId, "post_id": data.PostId})
		if err != common.ErrRecordNotFound {
			return nil
		}
	}
	
	{
		post, err := biz.getPostStore.Get(ctx, map[string]interface{}{"id": data.PostId})
		if err != nil || !post.IsEnabled || post.DeletedAt != nil {
			return nil
		}
	}

	if err := biz.store.Create(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(favoritemodel.EntityName, err)
	}

	return nil
}
