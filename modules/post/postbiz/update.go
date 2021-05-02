package postbiz

import (
	"context"
	"nhaancs/common"
	"nhaancs/modules/post/postmodel"
)

type UpdateStore interface {
	Get(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*postmodel.Post, error)
	UpdateData(
		ctx context.Context,
		id int,
		data *postmodel.PostUpdate,
	) error
}

type updateBiz struct {
	store UpdateStore
}

func NewUpdateBiz(store UpdateStore) *updateBiz {
	return &updateBiz{store: store}
}

func (biz *updateBiz) Update(ctx context.Context, id int, data *postmodel.PostUpdate) error {
	if err := data.Validate(); err != nil {
		return err
	}

	// todo: validate category
	// todo: validate only the author can update

	oldData, err := biz.store.Get(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return common.ErrCannotGetEntity(postmodel.EntityName, err)
	}
	if oldData.DeletedAt != nil {
		return common.ErrEntityDeleted(postmodel.EntityName, nil)
	}

	if data.Slug != "" {
		_, err := biz.store.Get(ctx, map[string]interface{}{"slug": data.Slug})
		if data.Slug != oldData.Slug && err != common.ErrRecordNotFound {
			return common.ErrEntityExisted(postmodel.EntityName, nil)
		}
	}

	if err := biz.store.UpdateData(ctx, id, data); err != nil {
		return common.ErrCannotUpdateEntity(postmodel.EntityName, err)
	}
	return nil
}
