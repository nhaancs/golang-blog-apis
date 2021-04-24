package productbiz

import (
	"context"
	"nhaancs/common"
	"nhaancs/modules/product/productmodel"
)

type UpdateProductStore interface {
	FindDataByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*productmodel.Product, error)

	UpdateData(
		ctx context.Context,
		id int,
		data *productmodel.ProductUpdate,
	) error
}

type updateProductBiz struct {
	store UpdateProductStore
}

func NewUpdateProductBiz(store UpdateProductStore) *updateProductBiz {
	return &updateProductBiz{store: store}
}

func (biz *updateProductBiz) UpdateProduct(ctx context.Context, id int, data *productmodel.ProductUpdate) error {
	oldData, err := biz.store.FindDataByCondition(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return common.ErrEntityNotFound(productmodel.EntityName, err)
	}
	if oldData.DeletedAt != nil {
		return common.ErrEntityDeleted(productmodel.EntityName, nil)
	}

	if err := biz.store.UpdateData(ctx, id, data); err != nil {
		return common.ErrCannotUpdateEntity(productmodel.EntityName, err)
	}

	return nil
}
