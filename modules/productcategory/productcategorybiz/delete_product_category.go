package productcategorybiz

import (
	"context"
	"nhaancs/common"
	"nhaancs/modules/productcategory/productcategorymodel"
)

type DeleteProductCategoryStore interface {
	FindDataByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*productcategorymodel.ProductCategory, error)

	SoftDelete(
		ctx context.Context,
		id int,
	) error
}

type deleteProductCategoryBiz struct {
	store DeleteProductCategoryStore
}

func NewDeleteProductCategoryBiz(store DeleteProductCategoryStore) *deleteProductCategoryBiz {
	return &deleteProductCategoryBiz{store: store}
}

func (biz *deleteProductCategoryBiz) DeleteProductCategory(ctx context.Context, id int) error {
	oldData, err := biz.store.FindDataByCondition(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return common.ErrEntityNotFound(productcategorymodel.EntityName, err)
	}
	if oldData.DeletedAt != nil {
		return common.ErrEntityDeleted(productcategorymodel.EntityName, nil)
	}

	if err := biz.store.SoftDelete(ctx, id); err != nil {
		return common.ErrCannotDeleteEntity(productcategorymodel.EntityName, err)
	}

	return nil
}
