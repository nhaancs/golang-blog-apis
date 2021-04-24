package productbiz

import (
	"context"
	"nhaancs/common"
	"nhaancs/modules/product/productmodel"
)

type DeleteProductStore interface {
	FindDataByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*productmodel.Product, error)

	SoftDelete(
		ctx context.Context,
		id int,
	) error
}

type deleteProductBiz struct {
	store DeleteProductStore
}

func NewDeleteProductBiz(store DeleteProductStore) *deleteProductBiz {
	return &deleteProductBiz{store: store}
}

func (biz *deleteProductBiz) DeleteProduct(ctx context.Context, id int) error {
	oldData, err := biz.store.FindDataByCondition(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return common.ErrEntityNotFound(productmodel.EntityName, err)
	}
	if oldData.DeletedAt != nil {
		return common.ErrEntityDeleted(productmodel.EntityName, nil)
	}

	if err := biz.store.SoftDelete(ctx, id); err != nil {
		return common.ErrCannotDeleteEntity(productmodel.EntityName, err)
	}

	return nil
}
