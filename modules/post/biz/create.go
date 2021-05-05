package postbiz

import (
	"context"
	"nhaancs/common"
	"nhaancs/modules/post/model"
)

type CreateStore interface {
	Create(ctx context.Context, data *postmodel.PostCreate) error
	Get(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*postmodel.Post, error)
}

type createBiz struct {
	store CreateStore
}

func NewCreateBiz(store CreateStore) *createBiz {
	return &createBiz{store: store}
}

func (biz *createBiz) Create(ctx context.Context, data *postmodel.PostCreate) error {
	if err := data.Validate(); err != nil {
		return err
	}
	{
		_, err := biz.store.Get(ctx, map[string]interface{}{"slug": data.Slug})
		if err != common.ErrRecordNotFound {
			return common.ErrEntityExisted(postmodel.EntityName, nil)
		}
	}

	// todo: validate category
	// todo: validate author. use common.Requester
	data.UserId = 1
	if err := biz.store.Create(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(postmodel.EntityName, err)
	}

	return nil
}
