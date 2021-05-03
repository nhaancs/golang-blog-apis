package categorybiz

import (
	"context"
	"nhaancs/common"
	"nhaancs/modules/category/model"
)

type CreateStore interface {
	Create(ctx context.Context, data *categorymodel.CategoryCreate) error
	Get(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*categorymodel.Category, error)
}

type createBiz struct {
	store CreateStore
}

func NewCreateBiz(store CreateStore) *createBiz {
	return &createBiz{store: store}
}

func (biz *createBiz) Create(ctx context.Context, data *categorymodel.CategoryCreate) error {
	if err := data.Validate(); err != nil {
		return err
	}
	{
		_, err := biz.store.Get(ctx, map[string]interface{}{"slug": data.Slug})
		if err != common.ErrRecordNotFound {
			return common.ErrEntityExisted(categorymodel.EntityName, nil)
		}
	}

	if err := biz.store.Create(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(categorymodel.EntityName, err)
	}

	return nil
}
