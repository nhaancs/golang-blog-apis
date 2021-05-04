package favoritebiz

import (
	"context"
	"nhaancs/common"
	favoritemodel "nhaancs/modules/favorite/model"
)

type FavoriteStore interface {
	Create(ctx context.Context, data *favoritemodel.FavoriteCreate) error
	Get(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*favoritemodel.Favorite, error)
}

type favoriteBiz struct {
	store FavoriteStore
}

func NewFavoriteBiz(store FavoriteStore) *favoriteBiz {
	return &favoriteBiz{store: store}
}

func (biz *favoriteBiz) Favorite(ctx context.Context, data *favoritemodel.FavoriteCreate) error {
	if err := data.Validate(); err != nil {
		return err
	}

	//todo: get a post to check(exist, not deleted) 
	{
		_, err := biz.store.Get(ctx, map[string]interface{}{"user_id": data.UserId, "post_id": data.PostId})
		if err != common.ErrRecordNotFound {
			return nil
		}
	}

	if err := biz.store.Create(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(favoritemodel.EntityName, err)
	}

	return nil
}
