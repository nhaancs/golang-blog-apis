package productcategorybiz

import (
	"context"
	"nhaancs/common"
	"nhaancs/modules/productcategory/productcategorymodel"
)

type UpdateProductCategoryStore interface {
	FindDataByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*productcategorymodel.ProductCategory, error)

	UpdateData(
		ctx context.Context,
		id int,
		data *productcategorymodel.ProductCategoryUpdate,
	) error
}

type updateProductCategoryBiz struct {
	store UpdateProductCategoryStore
}

func NewUpdateProductCategoryBiz(store UpdateProductCategoryStore) *updateProductCategoryBiz {
	return &updateProductCategoryBiz{store: store}
}

func (biz *updateProductCategoryBiz) UpdateProductCategory(ctx context.Context, id int, data *productcategorymodel.ProductCategoryUpdate) error {
	oldData, err := biz.store.FindDataByCondition(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return common.ErrEntityNotFound(productcategorymodel.EntityName, err)
	}
	if oldData.DeletedAt != nil {
		return common.ErrEntityDeleted(productcategorymodel.EntityName, nil)
	}

	if err := biz.store.UpdateData(ctx, id, data); err != nil {
		return common.ErrCannotUpdateEntity(productcategorymodel.EntityName, err)
	}

	return nil
}
