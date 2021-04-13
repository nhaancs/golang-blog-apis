package productcategorybiz

import (
	"context"
	"errors"
	"nhaancs/modules/productcategory/productcategorymodel"
)

type UpdateRestaurantStore interface {
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
	store UpdateRestaurantStore
}

func NewUpdateRestaurantBiz(store UpdateRestaurantStore) *updateProductCategoryBiz {
	return &updateProductCategoryBiz{store: store}
}

func (biz *updateProductCategoryBiz) UpdateProductCategory(ctx context.Context, id int, data *productcategorymodel.ProductCategoryUpdate) error {
	oldData , err := biz.store.FindDataByCondition(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return err
	}	
	if oldData.DeletedAt != nil {
		return errors.New("data deleted")
	}

	if err := biz.store.UpdateData(ctx, id, data); err != nil {
		return err
	}

	return nil
}
