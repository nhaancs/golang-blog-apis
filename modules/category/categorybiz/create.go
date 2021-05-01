package categorybiz

import (
	"context"
	"nhaancs/modules/category/categorymodel"
)

type CreateStore interface {
	Create(ctx context.Context, data *categorymodel.CategoryCreate) error
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

	err := biz.store.Create(ctx, data)

	return err
}
